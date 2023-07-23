package concurrent

import (
	"sync"
)

type Task struct {
	action func() error
	wg     sync.WaitGroup
	err    error
	mu     sync.Mutex
}

func NewTask(action func() error) *Task {
	return &Task{
		action: action,
	}
}

func (t *Task) Start() {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		err := t.action()
		if err != nil {
			t.mu.Lock()
			t.err = err
			t.mu.Unlock()
		}
	}()
}

func (t *Task) Wait() error {
	t.wg.Wait()
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.err
}
