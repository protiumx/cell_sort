package main

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

const size = 1000

func data() []int {
	l := make([]int, size)
	for i := range size {
		l[i] = rand.Intn(100)
	}
	return l
}

func BenchmarkCells(b *testing.B) {
	cells, start := build_list(data())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort_cells(start.right_neighbor, cells)
	}
}

func BenchmarkStd(b *testing.B) {
	l := data()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Ints(l)
	}
}

func BenchmarkSlices(b *testing.B) {
	l := data()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slices.Sort(l)
	}
}
