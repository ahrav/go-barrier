package barrier

import (
	"sync"
	"testing"
	"time"
)

func TestBarrier_Basic(t *testing.T) {
	const numGoroutines = 3
	b := newBarrier(numGoroutines)

	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id*100) * time.Millisecond) // Simulate different work times
			b.Await()
			if elapsed := time.Since(startTime); elapsed < time.Duration(200)*time.Millisecond {
				t.Errorf("Goroutine %d finished too early: %v", id, elapsed)
			}
		}(i)
	}

	wg.Wait()
}

func TestBarrier_MultipleRounds(t *testing.T) {
	const numGoroutines = 4
	const numRounds = 3
	b := newBarrier(numGoroutines)

	var wg sync.WaitGroup

	for round := 0; round < numRounds; round++ {
		wg.Add(numGoroutines)
		roundStart := time.Now()

		for i := 0; i < numGoroutines; i++ {
			go func(id, round int) {
				defer wg.Done()
				time.Sleep(time.Duration(id*50) * time.Millisecond)
				b.Await()
				elapsed := time.Since(roundStart)
				if elapsed < time.Duration(150)*time.Millisecond {
					t.Errorf("Round %d, Goroutine %d finished too early: %v", round, id, elapsed)
				}
			}(i, round)
		}

		wg.Wait()
	}
}

func TestBarrier_Stress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	const numGoroutines = 100
	const numRounds = 10
	b := newBarrier(numGoroutines)

	var wg sync.WaitGroup
	for round := 0; round < numRounds; round++ {
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				b.Await()
			}()
		}
		wg.Wait()
	}
}

// Benchmarks
func BenchmarkBarrier(b *testing.B) {
	benchmarks := []struct {
		name          string
		numGoroutines int
	}{
		{"Small", 2},
		{"Medium", 10},
		{"Large", 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			barrier := newBarrier(bm.numGoroutines)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var wg sync.WaitGroup
				wg.Add(bm.numGoroutines)

				for j := 0; j < bm.numGoroutines; j++ {
					go func() {
						defer wg.Done()
						barrier.Await()
					}()
				}

				wg.Wait()
			}
		})
	}
}