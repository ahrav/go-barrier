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
- **Custom reusable barrier** designed for multiple phases of synchronization.
- More to come as experimentation progresses!
