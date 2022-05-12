package go_math

import (
	"math/rand"
	"testing"
)

func gens[T ~int | ~float32 | ~float64](m, n int) [][]T {
	sss := make([][]T, m)
	for i := range sss {
		sss[i] = make([]T, n)
		for j := range sss[i] {
			sss[i][j] = T(rand.Float64() * 10)
		}
	}
	return sss
}

var xss = gens[int](48000, 10)
var hss = gens[int](9, 1)

func Test_filter2full(t *testing.T) {
	t.Log(xss)
	t.Log(hss)
	t.Log(filter2full[int](hss, xss))
}

func Benchmark_filter2full(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2full[int](hss, xss)
	}
}

func Test_filter2same(t *testing.T) {
	t.Log(xss)
	t.Log(hss)
	t.Log(filter2same[int](hss, xss))
}

func Benchmark_filter2same(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2same[int](hss, xss)
	}
}

func Test_filter2valid(t *testing.T) {
	t.Log(xss)
	t.Log(hss)
	t.Log(filter2valid[int](hss, xss))
}

func Benchmark_filter2valid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filter2valid[int](hss, xss)
	}
}
