package barrier

import "sync"

type Barrier struct {
	cond    *sync.Cond
	total   uint32
	current uint32
}

func NewBarrier(n uint32) *Barrier {
	return &Barrier{total: n, cond: sync.NewCond(new(sync.Mutex))}
}

func (b *Barrier) Wait() {
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
