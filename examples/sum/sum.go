package main

import (
	"fmt"
	"runtime"

	"github.com/cnnrznn/runner"
)

func main() {
	r := runner.New[int, int]()
	in, out, _ := r.Run(runtime.NumCPU(), func(in, out chan int, err chan error) {
		sum := 0

		for v := range in {
			sum += v
		}

		out <- sum
	})

	for i := 0; i < 10000; i++ {
		in <- i
	}
	close(in)

	total := 0

	for partial := range out {
		total += partial
	}

	fmt.Println(total)
}
