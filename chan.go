package barrier

// Package barrier implements a synchronization primitive that enables multiple goroutines
// to wait for each other at a specific point of execution before proceeding further.
// This is particularly useful in scenarios where you need to ensure that a group of
// goroutines have all reached a certain point before any of them can continue.

// barrier coordinates the synchronization using two channels - one for goroutines
// to signal their arrival and another to release them once all have arrived.
type barrier struct {
	inCh  chan bool // Channel for goroutines to signal their arrival.
	outCh chan bool // Channel to release waiting goroutines.
}

// newBarrier creates a barrier that coordinates 'count' number of goroutines.
// It launches a background goroutine to manage the synchronization lifecycle.
func newBarrier(count int) *barrier {
	result := &barrier{inCh: make(chan bool), outCh: make(chan bool)}
	go result.loop(count)
	return result
}

// Await blocks the calling goroutine until all other goroutines (as specified by count)
// have also called Await. This creates a synchronization point in the program.
func (b *barrier) Await() {
	b.inCh <- true
	<-b.outCh
}

// loop manages the barrier's lifecycle by collecting signals from arriving goroutines
// and releasing them once all have arrived. This process repeats indefinitely,
// allowing the barrier to be reused.
func (b *barrier) loop(count int) {
	for {
		// Wait for all goroutines to arrive.
		for i := 0; i < count; i++ {
			<-b.inCh
		}
		// Release all waiting goroutines simultaneously.
		for i := 0; i < count; i++ {
			b.outCh <- true
		}
	}
}
