package go_math

import (
	"testing"
)

var xss = [][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
}
var hss = [][]int{
	{1, 1, 1},
	{1, 1, 1},
	{1, 1, 1},
}

func Test_filter2full(t *testing.T) {
	t.Log(filter2full[int](hss, xss))
}

func Benchmark_filter2full(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2full[int](hss, xss)
	}
}

func Test_filter2same(t *testing.T) {
	t.Log(filter2same[int](hss, xss))
}

func Benchmark_filter2same(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2same[int](hss, xss)
	}
}

func Test_filter2valid(t *testing.T) {
	t.Log(filter2valid[int](hss, xss))
}

func Benchmark_filter2valid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2valid[int](hss, xss)
	}
}
