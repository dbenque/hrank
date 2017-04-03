package main

import "fmt"
import "sync"

type direction [2]int

var allDirections []direction

func init() {
	allDirections = []direction{
		{-1, 0}, {-1, -1}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {-1, 1}, {1, -1}, // optimized order to first find collision before marking path
	}
}

type board struct {
	row, column int
	grid        []byte
	validBoards map[string]struct{}
}

func main() {
	inputsBoard := readInput()
	var wg sync.WaitGroup
	wg.Add(len(inputsBoard))
	for _, input := range inputsBoard {
		go func(b *board) {
			defer wg.Done()
			b.work()
		}(input)
	}
	wg.Wait()

	for _, input := range inputsBoard {
		fmt.Printf("%d\n", len(input.validBoards))
	}
}

func readInput() []*board {
	testCases := []*board{}
	var count int
	fmt.Scanf("%d", &count)
	for t := 0; t < count; t++ {
		b := board{}
		fmt.Scanf("%d %d", &b.row, &b.column)
		b.grid = make([]byte, b.row*b.column)
		for l := 0; l < b.row; l++ {
			var line string
			fmt.Scanln(&line)
			for c := 0; c < b.column; c++ {
				b.grid[l*b.column+c] = line[c]
			}
		}
		b.validBoards = map[string]struct{}{}
		//		b.allBoards = map[string]struct{}{}
		testCases = append(testCases, &b)
	}
	return testCases
}

type queen struct {
	row, column int
}

//walk in the given directions and return true if no other queen was meet
func (q *queen) walk(d []direction, b *board) bool {
	ret := true
nextDir:
	for _, dir := range d {
		r, c := q.row, q.column
		for {
			r += dir[0]
			c += dir[1]
			switch b.get(r, c) {
			case 'Q':
				ret = false
			case '#', '0':
				continue nextDir
			case '.':
				b.set(r, c, 'x') // later no need to try that cell with a queen
			}
		}
	}
	return ret
}

//return the content of the cell or 0 if coordinate are out of the board
func (b *board) get(row, column int) byte {
	if !b.in(row, column) {
		return '0'
	}
	return b.grid[column+row*b.column]
}

func (b *board) set(row, column int, v byte) {
	b.grid[column+row*b.column] = v
}

func (b *board) in(row, column int) bool {
	return row >= 0 && column >= 0 && row < b.row && column < b.column
}

//return true if the grid is ok (no queen can fight)
func (b *board) validate() bool {
	for i, c := range b.grid {
		if c == 'Q' {
			q := queen{row: i / b.column, column: i % b.column}
			if !q.walk(allDirections, b) {
				return false
			}
		}
	}
	return true
}

func (b *board) clone() *board {
	newBoard := *b
	newBoard.grid = make([]byte, len(b.grid))
	copy(newBoard.grid, b.grid)
	return &newBoard
}

//Generate child board push then in chan and return the count generated
func (b *board) generateNext(children chan<- *board) {
	for i, c := range b.grid {
		if c == '.' {
			newboard := b.clone()
			newboard.grid[i] = 'Q'
			children <- newboard
		}
	}
	close(children)
}

func (b *board) addToValidBoard() bool {
	if _, ok := b.validBoards[string(b.grid)]; ok {
		return false
	}
	b.validBoards[string(b.grid)] = struct{}{}
	return true
}

func (b *board) work() {
	for i, c := range b.grid {
		if c == '.' {
			newboard := b.clone()
			newboard.grid[i] = 'Q'
			q := queen{row: i / b.column, column: i % b.column}
			if ok := q.walk(allDirections, newboard); ok {
				if newboard.addToValidBoard() {
					newboard.work()
				}
			}

		}
	}
}
