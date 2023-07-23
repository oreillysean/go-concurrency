package concurrent

import (
	"context"
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

	ctx := context.Background()
	err := WhenAll(ctx, tasks)
	if err != nil {
		t.Errorf("Expected no errors, but got back %v", err)
	}
}

func TestWhenAllTasks(t *testing.T) {
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

	tasks := []*Task{task1, task2, task3}

	ctx := context.Background()
	err := WhenAllTasks(ctx, tasks)
	if err != nil {
		t.Errorf("Expected no errors, but got back %v", err)
	}
}
