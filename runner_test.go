package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	r := New[int, int]()

	in, out, _ := r.Run(2, func(in, out chan int, err chan error) {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
	})

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	total := 0
	for v := range out {
		total += v
	}

	assert.Equal(t, 45, total)
}

func TestSumLarge(t *testing.T) {
	const n = 1000000

	r := New[int, int]()

	in, out, _ := r.Run(100, func(in, out chan int, err chan error) {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
	})

	for i := 0; i <= n; i++ {
		in <- i
	}
	close(in)

	total := 0

	for v := range out {
		total += v
	}

	assert.Equal(t, n*(n+1)/2, total)
}
