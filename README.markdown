    # Concurrent Go Functions

This package provides a set of concurrent Go (Golang) functions that allow you to efficiently perform concurrent tasks and manage synchronization in your applications.

## Functions

### ForEachReturn

```go
func ForEachReturn[T any, K any](data []T, action func(T) K) []K
```

This function takes a slice of data and applies the specified action function to each element concurrently. It creates a new slice containing the results of the action function applied to each element and returns it. The function ensures that all elements are processed concurrently, potentially improving the overall performance.

Usage Example:

```go
data := []int{1, 2, 3, 4, 5}
action := func(item int) int {
    return item * 2
}
result := ForEachReturn(data, action)
// 'result' will be a new slice containing [2, 4, 6, 8, 10]
```

### ForEachInPlace

```go
func ForEachInPlace[T any](data []T, action func(T) T)
```

This function takes a slice of data and applies the specified action function to each element concurrently. However, it updates the original 'data' slice in place with the results of the action function. This function can be used when you want to modify the existing data slice directly.

Usage Example:
```go
data := []string{"apple", "banana", "orange"}
action := func(item string) string {
    return strings.ToUpper(item)
}
ForEachInPlace(data, action)
// The 'data' slice will be modified to contain ["APPLE", "BANANA", "ORANGE"]
```

### ForEachInPlaceWithContext

```go
func ForEachInPlaceWithContext[T any](ctx context.Context, data []T, action func(context.Context, T) T)
```

This function is similar to `ForEachInPlace`, but it respects the cancellation signal from the provided context. If the context is canceled while processing elements, the function stops the concurrent processing and returns immediately. This ensures that you can handle scenarios where you may want to cancel the processing prematurely.

Usage Example:
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

data := []int{1, 2, 3, 4, 5}
action := func(ctx context.Context, item int) int {
    // Simulate some long-running task
    time.Sleep(time.Second)
    return item * 2
}

ForEachInPlaceWithContext(ctx, data, action)
// If the context is canceled before the processing completes, the function returns early.
```

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
```go
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