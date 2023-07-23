package concurrent

import (
	"context"
	"testing"
)

func TestForEachWithCancel(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel() // cancel when we are finished consuming integers

	data := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	ForEach(ctx, data, func(ctx context.Context, item int) int {
		return item * 2
	})

	if len(data) != len(expected) {
		t.Errorf("Expected length of %d, but got %d", len(expected), len(data))
	}

	for i := 0; i < len(expected); i++ {
		if data[i] != expected[i] {
			t.Errorf("Expected %d, but got %d at index %d", expected[i], data[i], i)
		}
	}
}

func TestForEach(t *testing.T) {
	ctx := context.Background()
	data := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	ForEach(ctx, data, func(ctx context.Context, item int) int {
		return item * 2
	})

	if len(data) != len(expected) {
		t.Errorf("Expected length of %d, but got %d", len(expected), len(data))
	}

	for i := 0; i < len(expected); i++ {
		if data[i] != expected[i] {
			t.Errorf("Expected %d, but got %d at index %d", expected[i], data[i], i)
		}
	}
}
