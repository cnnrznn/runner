package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	r := New[int, int](2)
	in, out := r.In(), r.Out()

	go r.Run(func(in, out chan int, err chan error) {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
	})

	for i := 0; i < 10; i++ {
		in <- i
	}

	r.Done()

	total := 0
	for v := range out {
		total += v
	}

	assert.Equal(t, 45, total)
}

func TestSumLarge(t *testing.T) {
	r := New[int, int](100)
	in, out := r.In(), r.Out()

	const n = 1000000

	go r.Run(func(in, out chan int, err chan error) {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
	})

	go func() {
		defer r.Done()

		for i := 0; i <= n; i++ {
			in <- i
		}
	}()

	total := 0

	for v := range out {
		total += v
	}

	assert.Equal(t, n*(n+1)/2, total)
}
