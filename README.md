# runner

This package implements a thread(goroutine) worker pool.

This project was inspired by having to re-implement a job pool across projects.
I noticed the same paradigms repeating themselves:

1. Instantiate input and output channels
1. Launch worker goroutines
1. Write to and close input channel
1. Wait for goroutines to finish processing
1. Close output channel

This produces a lot of synchronization boilerplate that is repeated across projects.
This library aims to eliminate that boilerplate by offering a simple interface for
defining and using a Job pool.

Users can define their own output types using Golang generics.
The runner package is type agnostic.

## Usage

See `examples/` for a collection of working code.
