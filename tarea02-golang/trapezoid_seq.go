package main

import (
	"fmt"
	"math"
	"time"
)

func TrapezoidRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

func main() {

	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}

	n := 1000000000

	for i := 0; i < 10; i++ {
		start := time.Now()
		TrapezoidRule(f, 5, 20, n)
		elapsed := time.Since(start).Nanoseconds()
		fmt.Println(elapsed)
	}
}
