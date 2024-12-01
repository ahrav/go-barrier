package barrier

// Package barrier implements a synchronization primitive that enables multiple goroutines
// to wait for each other at a specific point of execution before proceeding further.
// This is particularly useful in scenarios where you need to ensure that a group of
// goroutines have all reached a certain point before any of them can continue.

// chanBarrier coordinates the synchronization using two channels - one for goroutines
// to signal their arrival and another to release them once all have arrived.
type chanBarrier struct {
	inCh  chan bool // Channel for goroutines to signal their arrival.
	outCh chan bool // Channel to release waiting goroutines.
}

// newBarrier creates a barrier that coordinates 'count' number of goroutines.
// It launches a background goroutine to manage the synchronization lifecycle.
func newBarrier(count uint32) *chanBarrier {
	result := &chanBarrier{inCh: make(chan bool), outCh: make(chan bool)}
	go result.loop(count)
	return result
}

// Wait blocks the calling goroutine until all other goroutines (as specified by count)
// have also called Wait. This creates a synchronization point in the program.
func (b *chanBarrier) Wait() {
	b.inCh <- true
	<-b.outCh
}

// loop manages the barrier's lifecycle by collecting signals from arriving goroutines
// and releasing them once all have arrived. This process repeats indefinitely,
// allowing the barrier to be reused.
func (b *chanBarrier) loop(count uint32) {
	for {
		// Wait for all goroutines to arrive.
		for range count {
			<-b.inCh
		}
		// Release all waiting goroutines simultaneously.
		for range count {
			b.outCh <- true
		}
	}
}
