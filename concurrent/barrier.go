package concurrent

import (
	"sync"
)

// Barrier is a concurrent synchronization primitive.
// It allows multiple goroutines to wait until a specified number of participants arrive at the barrier.
type Barrier struct {
	participants int              // The total number of participants required to reach the barrier.
	count        int              // The number of participants remaining to reach the barrier.
	mutex        sync.Mutex       // Mutex for protecting the count variable.
	cond         *sync.Cond       // Condition variable used to block and wake up goroutines.
	action       func(b *Barrier) // The function to be executed when all participants reach the barrier.
}

// NewBarrier creates a new Barrier with the specified number of participants and action function.
// The action function will be executed whenever all participants arrive at the barrier.
func NewBarrier(participants int, action func(b *Barrier)) *Barrier {
	b := &Barrier{
		participants: participants,
		count:        participants,
		action:       action,
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

// SignalAndWait signals that a participant has arrived at the barrier and waits until all participants arrive.
// When all participants have reached the barrier, the action function is executed, and all waiting goroutines are released.
func (b *Barrier) SignalAndWait() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.count--

	if b.count > 0 {
		b.cond.Wait()
	} else {
		b.count = b.participants
		b.action(b)

		// Broadcast to all waiting goroutines that the barrier has been reached.
		b.cond.Broadcast()
	}
}
