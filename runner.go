package runner

import "sync"

type Runner[I, O any] interface {
	Run(logic func(in chan I, out chan O, err chan error))

	In() chan<- I
	Out() <-chan O
	Err() <-chan error

	Done()
}

type runner[I, O any] struct {
	nWorkers int
	in       chan I
	out      chan O
	err      chan error
}

func New[I, O any](n int) Runner[I, O] {
	return &runner[I, O]{
		nWorkers: n,
		in:       make(chan I),
		out:      make(chan O),
		err:      make(chan error),
	}
}

func (r *runner[I, O]) Run(logic func(in chan I, out chan O, err chan error)) {
	wg := sync.WaitGroup{}
	wg.Add(r.nWorkers)

	for i := 0; i < r.nWorkers; i++ {
		go func() {
			logic(r.in, r.out, r.err)
			wg.Done()
		}()
	}

	wg.Wait()

	close(r.out)
	close(r.err)
}

func (r *runner[I, O]) Done() {
	close(r.in)
}

func (r *runner[I, O]) In() chan<- I {
	return r.in
}

func (r *runner[I, O]) Out() <-chan O {
	return r.out
}

func (r *runner[I, O]) Err() <-chan error {
	return r.err
}
