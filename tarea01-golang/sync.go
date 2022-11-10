package main

import (
	"fmt"
	"sync"
	"time"
)

var time1 time.Duration
var wg sync.WaitGroup

type Trapezoid struct {
	f func(float64) float64
	a float64
	b float64
	n int
}

func (t *Trapezoid) TrapezoidRule() float64 {
	h := (t.b - t.a) / float64(t.n)
	sum := 0.5 * (t.f(t.a) + t.f(t.b))
	for i := 1; i < t.n; i++ {
		wg.Add(1)
		sum += t.f(t.a + float64(i)*h)
		wg.Done()
	}
	return sum * h
}

func main() {
	f := func(x float64) float64 {
		return ((x*x + 1) / 2)
	}
	t := Trapezoid{f, 5, 20, 10}
	start := time.Now()
	result := t.TrapezoidRule()
	elapsed := time.Since(start).Nanoseconds()
	defer fmt.Println(elapsed)
	area := fmt.Sprintf("Area: %v", result)
	fmt.Println(area)
	wg.Wait()
}
