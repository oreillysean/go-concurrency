package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBarrier(t *testing.T) {
	const numParticipants = 3

	barrier := NewBarrier(numParticipants, func(b *Barrier) {
		fmt.Printf("Post-Phase action: count=%d, phase=%d\n", b.participants, b.count)
	})

	for i := 0; i < numParticipants; i++ {
		go func(id int) {
			fmt.Printf("Participant %d is waiting\n", id)
			barrier.SignalAndWait()
			fmt.Printf("Participant %d has passed the barrier\n", id)
		}(i)
	}

	// Wait for goroutines to complete
	time.Sleep(time.Second * 10)
}
