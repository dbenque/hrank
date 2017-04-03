package main

import "fmt"

type direction [2]int

var allDirections []direction

func init() {
	allDirections = []direction{
		{-1, 0}, {-1, -1}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {-1, 1}, {1, -1}, // optimized order to first find collision before marking path
	}
}

type board struct {
	row, column int
	blocks      map[uint64]struct{}
}

type queen struct {
	row, column int
}

//walk in the given directions and return true if no other queen was meet
func (q *queen) walk(d []direction, b *board) int {
	ret := 0
nextDir:
	for _, dir := range d {
		r, c := q.row, q.column
		for {
			r += dir[0]
			c += dir[1]
			switch b.get(r, c) {
			case -1, 1:
				continue nextDir
			case 0:
				ret++
			}
		}
	}
	return ret
}

// 0 contain nothing, -1 out of board, 1 contain blocker
func (b *board) get(row, column int) int {
	if !b.in(row, column) {
		return -1
	}
	if _, ok := b.blocks[uint64(row)+(uint64(column)<<uint64(32))]; ok {
		return 1
	}
	return 0
}

func (b *board) in(row, column int) bool {
	return row >= 0 && column >= 0 && row < b.row && column < b.column
}
func main() {
	b, q := readInput()
	fmt.Printf("%d", q.walk(allDirections, b))
}

func readInput() (*board, *queen) {
	b := board{}
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	b.row = n
	b.column = n
	q := queen{}
	fmt.Scanf("%d %d", &q.row, &q.column)
	q.row--
	q.column--
	b.blocks = map[uint64]struct{}{}
	var r, c int
	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d", &r, &c)
		b.blocks[uint64(r-1)+(uint64(c-1)<<uint64(32))] = struct{}{}
	}
	return &b, &q
}
