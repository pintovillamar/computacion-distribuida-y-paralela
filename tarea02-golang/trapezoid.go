package main

import (
	"fmt"
	"math"
	"time"
)

func worker(jobs chan int, results chan float64) {
	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}

	for n := range jobs {
		results <- TrapezoidRule(f, 5, 20, n)
	}
}

func TrapezoidRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

func main() {

	n := 1000000000

	jobs := make(chan int, n)
	results := make(chan float64, n)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	start := time.Now()
	for i := 0; i < n; i++ {
		jobs <- i
	}
	elapsed := time.Since(start).Nanoseconds()
	fmt.Println(elapsed)

	close(jobs)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(<-results)
	// }
}
