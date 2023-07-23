package concurrent

import (
	"fmt"
	"sync"
)

// WhenAll executes multiple tasks concurrently and waits for all of them to finish.
// It returns an error if any of the tasks encounter an error. The returned error
// contains information about the number of errors encountered during execution.
func WhenAll(tasks []func() error) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(tasks))

	// Launch goroutines for each task to be executed concurrently.
	for _, task := range tasks {
		wg.Add(1)
		go func(t func() error) {
			defer wg.Done()
			err := t()
			if err != nil {
				errChan <- err // Send the error to the errChan if encountered.
			}
		}(task)
	}

	wg.Wait()
	close(errChan) // Close the errChan to signal that all goroutines have finished.

	var errors []error
	for err := range errChan {
		errors = append(errors, err) // Collect the errors from the errChan.
	}

	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors", len(errors)) // Return an error with error count information.
	}

	return nil
}
