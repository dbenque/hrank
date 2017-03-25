package main

import "testing"
import "sync"

func BenchmarkWork(b *testing.B) {
	r, c := 8, 8
	// 8x8
	// testing: warning: no tests to run
	// BenchmarkWork-8                2         827101630 ns/op
	// PASS
	// ok      github.com/dbenque/queensOnBoard        2.496s
	for n := 0; n < b.N; n++ {
		b := board{
			row:         r,
			column:      c,
			grid:        make([]byte, c*r),
			validBoards: &boardIndex{index: map[string]struct{}{}},
			wgChildren:  &sync.WaitGroup{},
		}
		for i := range b.grid {
			b.grid[i] = '.'
		}
		b.wgChildren.Add(1)
		b.work()
		b.wgChildren.Wait()
	}
}
