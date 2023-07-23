package concurrent

import (
	"fmt"
	"testing"
)

func TestWhenAll(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("Task 1")
			return nil
		},
		func() error {
			fmt.Println("Task 2")
			return nil
		},
		func() error {
			fmt.Println("Task 3")
			return nil
		},
	}

	err := WhenAll(tasks)
	if err != nil {
		t.Errorf("Expected no errors, but got back %v", err)
	}
}
