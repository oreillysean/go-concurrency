package concurrent

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWhenAny(t *testing.T) {
	// Define test tasks that simulate work and return errors.
	task1 := func() error {
		time.Sleep(2 * time.Second)
		return errors.New("Task 1 failed")
	}

	task2 := func() error {
		time.Sleep(1 * time.Second)
		return errors.New("Task 2 failed")
	}

	task3 := func() error {
		time.Sleep(3 * time.Second)
		return nil
	}

	tasks := []func() error{task1, task2, task3}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := WhenAny(ctx, tasks)

	expectedErrorMessage := "Task 2 failed"
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message: %q, but got: %q", expectedErrorMessage, err.Error())
	}
}
