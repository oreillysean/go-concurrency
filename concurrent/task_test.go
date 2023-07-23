package concurrent

import (
	"fmt"
	"testing"
)

func TestNewTask(t *testing.T) {
	task1 := NewTask(func() error {
		fmt.Println("Task 1")
		return nil
	})

	task2 := NewTask(func() error {
		fmt.Println("Task 2")
		return nil
	})

	task3 := NewTask(func() error {
		fmt.Println("Task 3")
		return nil
	})

	// Start all tasks concurrently.
	task1.Start()
	task2.Start()
	task3.Start()

	// Wait for all tasks to finish.
	err1 := task1.Wait()
	err2 := task2.Wait()
	err3 := task3.Wait()

	// Check for errors (if any).
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("one or more of the tasks has errors, this should not have happened")
	}
}
