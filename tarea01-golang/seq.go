package main

import (
	"fmt"
	"math"
	"strings"
	"text/template"
	"time"
)

func format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
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
	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}
	n := 1
	start := time.Now()
	elapsed := time.Since(start).Nanoseconds()
	result := TrapezoidRule(f, 5, 20, n)
	area := format("Area: {{ . }}.", result)
	fmt.Println(area)
	time := format("Time: {{ . }} in ns.", elapsed)
	fmt.Println(time)
	fmt.Println()
}
