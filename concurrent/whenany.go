package concurrent

import (
	"context"
)

// WhenAny executes multiple tasks concurrently and returns the error from the first task to complete.
// It respects the cancellation signal from the provided context.
func WhenAny(ctx context.Context, tasks []func() error) error {
	type result struct {
		index int
		err   error
	}

	results := make(chan result, len(tasks))

	// Launch goroutines for each task to be executed concurrently.
	for i, task := range tasks {
		go func(i int, t func() error) {
			err := t()
			if err != nil {
				select {
				case results <- result{i, err}:
				case <-ctx.Done():
					return // Exit the goroutine early if the context is canceled.
				}
			}
		}(i, task)
	}

	select {
	case res := <-results:
		return res.err // Return the error from the first task to complete.
	case <-ctx.Done():
		return ctx.Err() // Return context error if the context is canceled.
	}
}
