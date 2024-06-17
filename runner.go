package runner

import "sync"

type Runner[I, O any] interface {
	Run(n int, logic func(in chan I, out chan O, err chan error)) (chan I, chan O, chan error)
}

type runner[I, O any] struct {
}

func New[I, O any]() Runner[I, O] {
	return &runner[I, O]{}
}

func (r *runner[I, O]) Run(
	nWorkers int,
	logic func(in chan I, out chan O, err chan error),
) (chan I, chan O, chan error) {
	in, out, err := make(chan I), make(chan O), make(chan error)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(nWorkers)

		for i := 0; i < nWorkers; i++ {
			go func() {
				logic(in, out, err)
				wg.Done()
			}()
		}

		wg.Wait()

		close(out)
		close(err)
	}()

	return in, out, err
}
