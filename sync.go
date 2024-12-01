package barrier

import "sync"

// syncBarrier implements a reusable synchronization point using sync.Cond.
// This implementation is more efficient than channels for larger numbers of
// goroutines as it avoids the overhead of message passing.
type syncBarrier struct {
	cond    *sync.Cond  // Coordinates goroutine synchronization
	total   uint32      // Total number of goroutines to synchronize
	current uint32      // Number of goroutines currently waiting
}

// NewBarrier creates a barrier that coordinates n goroutines.
// The barrier can be reused for multiple synchronization rounds.
func NewBarrier(n uint32) *syncBarrier {
	return &syncBarrier{total: n, cond: sync.NewCond(new(sync.Mutex))}
}

// Wait blocks until all goroutines have called Wait. The last goroutine
// to arrive releases all waiting goroutines and resets the barrier.
// This method is safe for concurrent access and can be called repeatedly
// for multiple synchronization rounds.
func (b *syncBarrier) Wait() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.current++
	if b.current == b.total {
		b.current = 0      // Reset for reuse
		b.cond.Broadcast() // Wake up all waiting goroutines
	} else {
		b.cond.Wait() // Wait until all participants arrive
	}
}
