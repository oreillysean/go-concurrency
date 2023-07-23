package concurrent

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	ctx := context.Background()
	timer := NewTimer(ctx, 1*time.Second, func() {
		fmt.Println("Timer Elapsed")
	})

	timer.Start()

	time.Sleep(5 * time.Second)

	timer.Stop()
}
