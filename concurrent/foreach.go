package concurrent

import (
	"context"
	"sync"
)

// chSlice is a generic type used as a channel slice to record the index of the result and the result itself.
type chSlice[T any] struct {
	i   int // Stores the index of the result.
	res T   // Stores the result of the action.
}

// ForEachReturn takes a slice of data and passes each entry to action. This creates a new slice
// and returns it.
func ForEachReturn[T any, K any](data []T, action func(T) K) []K {
	var wg sync.WaitGroup
	c := make(chan chSlice[K], len(data))
	defer close(c)

	// Launch goroutines to process each item in the data slice concurrently.
	for i, item := range data {
		wg.Add(1)

		chs := chSlice[K]{
			i: i,
		}

		go func(item T, c chan chSlice[K], chs chSlice[K]) {
			defer wg.Done()
			chs.res = action(item)
			c <- chs
		}(item, c, chs)
	}

	// Collect the results from the channel and create a new slice with the processed results.
	result := make([]K, len(data))
	for i := 0; i < len(data); i++ {
		chs := <-c
		result[chs.i] = chs.res
	}

	wg.Wait()

	return result
}

// ForEachInPlace takes a slice of data and passes each entry to action. This does not create
// a new slice. It updates data in place.
func ForEachInPlace[T any](data []T, action func(T) T) {
	var wg sync.WaitGroup
	c := make(chan chSlice[T], len(data))
	defer close(c)

	// Launch goroutines to process each item in the data slice concurrently.
	for i, item := range data {
		wg.Add(1)

		chs := chSlice[T]{
			i: i,
		}

		go func(item T, c chan chSlice[T], chs chSlice[T]) {
			defer wg.Done()
			chs.res = action(item)
			c <- chs
		}(item, c, chs)
	}

	// Collect the results from the channel and update the data slice in place with the processed results.
	for i := 0; i < len(data); i++ {
		chs := <-c
		data[chs.i] = chs.res
	}

	wg.Wait()
}

// ForEachInPlaceWithContext takes a slice of data, passes each entry to action, and updates data in place.
// It respects the cancellation signal from the provided context.
func ForEachInPlaceWithContext[T any](ctx context.Context, data []T, action func(context.Context, T) T) {
	var wg sync.WaitGroup
	c := make(chan chSlice[T], len(data))
	defer close(c)

	// Launch goroutines to process each item in the data slice concurrently, with respect to the context.
	for i, item := range data {
		wg.Add(1)

		chs := chSlice[T]{
			i: i,
		}

		go func(ctx context.Context, item T, c chan chSlice[T], chs chSlice[T]) {
			defer wg.Done()

			// Check if the context is canceled before processing the item.
			select {
			case <-ctx.Done():
				return // If the context is canceled, return early.
			default:
				chs.res = action(ctx, item)
				c <- chs
			}
		}(ctx, item, c, chs)
	}

	// Collect the results from the channel and update the data slice in place with the processed results.
	for i := 0; i < len(data); i++ {
		select {
		case <-ctx.Done():
			wg.Wait() // Wait for all goroutines to finish before returning if the context is canceled.
			return    // Return if the context is canceled.
		case chs := <-c:
			data[chs.i] = chs.res
		}
	}

	wg.Wait()
}
