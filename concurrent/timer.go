package concurrent

import (
	"context"
	"sync"
	"time"
)

// Timer represents a timer that triggers an action at regular intervals until stopped.
type Timer struct {
	ctx      context.Context // The context for managing the timer's lifecycle.
	duration time.Duration   // The duration between each action trigger.
	ticker   *time.Ticker    // Ticker for regular time intervals.
	stopChan chan bool       // Channel for stopping the timer.
	wg       sync.WaitGroup  // WaitGroup for coordinating the timer's goroutine.
	action   func()          // The function to be executed on each timer tick.
}

// NewTimer creates a new Timer instance with the given context, duration, and action function.
func NewTimer(ctx context.Context, duration time.Duration, action func()) *Timer {
	return &Timer{
		duration: duration,
		stopChan: make(chan bool),
		action:   action,
		ctx:      ctx,
	}
}

// Start starts the timer's goroutine, triggering the action at regular intervals.
func (t *Timer) Start() {
	t.ticker = time.NewTicker(t.duration)
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-t.ctx.Done():
				t.Stop() // If the context is canceled, stop the timer.
				return
			case <-t.ticker.C:
				t.action() // Execute the action on each timer tick.
			case <-t.stopChan:
				t.ticker.Stop() // Stop the ticker when explicitly requested to stop.
				return
			}
		}
	}()
}

// Stop stops the timer and waits for the timer's goroutine to finish.
func (t *Timer) Stop() {
	close(t.stopChan) // Close the stop channel to signal the timer's goroutine to stop.
	t.wg.Wait()       // Wait for the timer's goroutine to finish before returning.
}
