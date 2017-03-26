package main

import "fmt"
import "sync"

type direction [2]int

var allDirections []direction

func init() {
	allDirections = []direction{
		{-1, 0}, {-1, -1}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {-1, 1}, {1, -1}, // optimized order to first find collision before marking path
	}

	// for _, r := range []int{-1, 0, 1} {
	// 	for _, c := range []int{-1, 0, 1} {
	// 		if r != 0 || c != 0 {
	// 			allDirections = append(allDirections, direction{r, c})
	// 		}
	// 	}
	// }
}

type boardIndex struct {
	sync.Mutex
	index map[string]struct{}
}
type board struct {
	row, column int
	grid        []byte
	validBoards *boardIndex
	allBoards   *boardIndex
}

func main() {
	inputsBoard := readInput()
	var wg sync.WaitGroup
	wg.Add(len(inputsBoard))
	for _, input := range inputsBoard {
		go func(b *board) {
			defer wg.Done()
			b.work2()
		}(input)
	}
	wg.Wait()

	for _, input := range inputsBoard {
		fmt.Printf("%d\n", len(input.validBoards.index))
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
		b.validBoards = &boardIndex{index: map[string]struct{}{}}
		b.allBoards = &boardIndex{index: map[string]struct{}{}}
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
	b.validBoards.Lock()
	if _, ok := b.validBoards.index[string(b.grid)]; ok {
		b.validBoards.Unlock()
		return false
	}
	b.validBoards.index[string(b.grid)] = struct{}{}
	b.validBoards.Unlock()
	return true
}

//return true if that board has never beeb seen before
func (b *board) addToAllBoard() bool {
	b.allBoards.Lock()
	if _, ok := b.allBoards.index[string(b.grid)]; ok {
		b.allBoards.Unlock()
		return false
	}
	b.allBoards.index[string(b.grid)] = struct{}{}
	b.allBoards.Unlock()
	return true
}

func (b *board) work() {
	if !b.addToAllBoard() {
		return
	}
	if !b.validate() || !b.addToValidBoard() {
		return
	}

	children := make(chan *board, b.row*b.column/4) // dynamic sizing of buffered chan
	go b.generateNext(children)
childLoop:
	for {
		select {
		case child, ok := <-children:
			if !ok {
				break childLoop
			}
			child.work()
		}
	}
}

//-----------------------
func (b *board) work2() {
	for i, c := range b.grid {
		if c == '.' {
			b.grid[i] = 'Q'
			if !b.addToAllBoard() { // check if that board was already seen
				b.grid[i] = '.'
				continue
			}
			newboard := b.clone() // the one with clean signature (without the x)
			b.grid[i] = '.'

			q := queen{row: i / b.column, column: i % b.column}
			if !q.walk(allDirections, newboard) {
				continue
			}
			newboard.addToValidBoard()
			newboard.work2()
		}
	}
}
