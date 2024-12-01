# go-barrier

**go-barrier** is an experimental Go package exploring different implementations of barriers,
a synchronization primitive used to coordinate multiple goroutines at specific points during execution.

## What is a Barrier?

A barrier is a synchronization mechanism that ensures multiple goroutines (or threads) reach a defined point in execution before any of them proceed. This is particularly useful in parallel computing, where tasks are divided into phases, and all participants must complete one phase before starting the next.

## Why this Repo?

This repository serves as a learning sandbox to experiment with implementing barriers in Go using various techniques. It's not intended for production use but rather to:
- Explore the concept of barriers in Go.
- Compare different approaches and their trade-offs.
- Provide examples and insights into implementing synchronization primitives in Go.

## Implementations

The repository includes:
- **Basic barrier implementation** using `sync.Cond`.
- **Channel-based barrier** for lightweight synchronization.
- More to come as experimentation progresses!

## Benchmarks

The following benchmarks compare different barrier implementations on an Apple M1 Pro (arm64):

```text
goos: darwin
goarch: arm64
pkg: github.com/ahrav/go-barrier
cpu: Apple M1 Pro
=== RUN   BenchmarkBarrierComparison
BenchmarkBarrierComparison
=== RUN   BenchmarkBarrierComparison/Sync-Small-2
BenchmarkBarrierComparison/Sync-Small-2
BenchmarkBarrierComparison/Sync-Small-2-8                1433503               790.7 ns/op            64 B/op          3 allocs/op
=== RUN   BenchmarkBarrierComparison/Chan-Small-2
BenchmarkBarrierComparison/Chan-Small-2
BenchmarkBarrierComparison/Chan-Small-2-8                 961046              1235 ns/op              64 B/op          3 allocs/op
=== RUN   BenchmarkBarrierComparison/Sync-Medium-10
BenchmarkBarrierComparison/Sync-Medium-10
BenchmarkBarrierComparison/Sync-Medium-10-8               260760              4546 ns/op             256 B/op         11 allocs/op
=== RUN   BenchmarkBarrierComparison/Chan-Medium-10
BenchmarkBarrierComparison/Chan-Medium-10
BenchmarkBarrierComparison/Chan-Medium-10-8               184558              6580 ns/op             256 B/op         11 allocs/op
=== RUN   BenchmarkBarrierComparison/Sync-Large-100
BenchmarkBarrierComparison/Sync-Large-100
BenchmarkBarrierComparison/Sync-Large-100-8                21662             63892 ns/op            2429 B/op        101 allocs/op
=== RUN   BenchmarkBarrierComparison/Chan-Large-100
BenchmarkBarrierComparison/Chan-Large-100
BenchmarkBarrierComparison/Chan-Large-100-8                19168             58819 ns/op            2428 B/op        101 allocs/op
```

