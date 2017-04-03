package main

import "testing"

func BenchmarkWork(b *testing.B) {
	r, c := 8, 8
	// testing: warning: no tests to run
	// BenchmarkWork-8                5         260846520 ns/op
	// PASS
	// ok      github.com/dbenque/queensOnBoard        2.357s
	for n := 0; n < b.N; n++ {
		b := board{
			row:         r,
			column:      c,
			grid:        make([]byte, c*r),
			validBoards: map[string]struct{}{},
			//allBoards:   map[string]struct{}{},
		}
		for i := range b.grid {
			b.grid[i] = '.'
		}
		b.work()
	}
}
