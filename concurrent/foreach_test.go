package concurrent

import (
	"context"
	"testing"
)

func TestForEachReturn(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	result := ForEachReturn(data, func(item int) int {
		return item * 2
	})

	if len(result) != len(expected) {
		t.Errorf("Expected length of %d, but got %d", len(expected), len(result))
	}

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %d, but got %d at index %d", expected[i], result[i], i)
		}
	}
}

func FuzzForEachInPlace(f *testing.F) {
	var byteSlice []byte

	for i := 0; i < 1000000; i++ {
		byteSlice = append(byteSlice, byte(i))
	}

	f.Add(byteSlice)
	f.Fuzz(func(t *testing.T, byteSlice []byte) {
		intSlice := make([]int, len(byteSlice))
		for i, b := range byteSlice {
			intSlice[i] = int(b)
		}

		expected := intSlice[:]

		ForEachInPlace(intSlice, func(i int) int {
			return i
		})

		for i := 0; i < len(expected); i++ {
			if intSlice[i] != expected[i] {
				t.Errorf("Expected %d, but got %d at index %d", expected[i], intSlice[i], i)
			}
		}
	})
}
func TestForEachInPlace(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	ForEachInPlace(data, func(item int) int {
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

func TestForEachInPlaceWithCancelContext(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel() // cancel when we are finished consuming integers

	data := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	ForEachInPlaceWithContext(ctx, data, func(ctx context.Context, item int) int {
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

func TestForEachInPlaceWithContext(t *testing.T) {
	ctx := context.Background()
	data := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	ForEachInPlaceWithContext(ctx, data, func(ctx context.Context, item int) int {
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
