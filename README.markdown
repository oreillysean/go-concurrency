    # Concurrent Go Functions

This package provides a set of concurrent Go (Golang) functions that allow you to efficiently perform concurrent tasks and manage synchronization in your applications.

## Functions

The `concurrent` package provides a useful function called `ForEach` that allows you to process a slice of data concurrently using goroutines. The function respects the cancellation signal from the provided context, enabling safe and controlled concurrent processing.

## `ForEach` Function Signature

```golang
func ForEach[T any](ctx context.Context, data []T, action func(context.Context, T) T)`
```

## Usage

The `ForEach` function takes three arguments:

1.  `ctx context.Context`: The context used to control the execution of the goroutines. It allows for graceful cancellation if needed.
2.  `data []T`: The input slice of data that you want to process concurrently. The `ForEach` function will process each element in this slice.
3.  `action func(context.Context, T) T`: A function that takes the context and an element from the data slice as input, and returns the processed result.

## How It Works

1.  The `ForEach` function initializes a channel `c` to hold intermediate results. It creates a goroutine for each element in the data slice, which processes the item and sends the result with its index over the channel.

2.  The `ForEach` function loops through the data slice, collecting the processed results from the channel and updating the data slice in place.

3.  While processing each element concurrently, the function checks the cancellation status of the provided context before processing an item. If the context is canceled, the goroutine returns early without performing any processing on that specific item.

4.  The function respects the cancellation signal and ensures that all running goroutines are allowed to finish before returning when the context is canceled.


## Example

Here's an example demonstrating how to use the `ForEach` function:

```golang
package main

import (
    "context"
    "fmt"
    "github.com/oreillysean/go-concurrency"
)

func main() {
data := []int{1, 2, 3, 4, 5}

	// A sample action that doubles the value of each element.
	action := func(ctx context.Context, val int) int {
		return val * 2
	}

	ctx := context.Background()
	concurrent.ForEach(ctx, data, action)

	fmt.Println(data) // Output: [2 4 6 8 10]
}
```

In this example, the `ForEach` function is used to double each element in the `data` slice concurrently. The resulting `data` slice will contain the doubled values after the `ForEach` function completes the processing.

### Timer

```go
type Timer struct {
    ctx context.Context
    duration time.Duration
    ticker *time.Ticker
    stopChan chan bool
    wg sync.WaitGroup
    action func()
}
```

This struct represents a timer that triggers a specified action function at regular intervals until stopped. The Timer uses goroutines and channels to provide concurrent execution.

Usage Example:
```golang
func main() {
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

action := func() {
    fmt.Println("Timer tick!")
}

duration := 2 * time.Second
timer := NewTimer(ctx, duration, action)
timer.Start()

time.Sleep(10 * time.Second) // Wait for the timer ticks.
timer.Stop() // Stop the timer after 10 seconds.
}
```

### WhenAll

```go
func WhenAll(tasks []func() error) error
```

This function takes a slice of functions, where each function represents a task that returns an error, and executes them concurrently. It waits for all tasks to finish and collects any encountered errors. If any of the tasks return an error, it returns an error indicating the number of errors encountered during execution.

Usage Example:
```go
func main() {
    task1 := func() error {
    time.Sleep(2 * time.Second)
    fmt.Println("Task 1 complete!")
    return nil
}

task2 := func() error {
    time.Sleep(3 * time.Second)
    fmt.Println("Task 2 complete!")
    return fmt.Errorf("Task 2 encountered an error")
}

err := WhenAll([]func() error{task1, task2})
if err != nil {
    fmt.Println(err) // Output: "encountered 1 errors"
}
}
```

These concurrent functions offer powerful tools to improve the performance and efficiency of your Go applications. They handle concurrent execution, synchronization, and cancellation gracefully, allowing you to build responsive and scalable applications.