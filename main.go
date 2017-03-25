package main

import "fmt"
import "sync"

type direction [2]int

var allDirections []direction

func init() {
	allDirections = []direction{}
	for _, r := range []int{-1, 0, 1} {
		for _, c := range []int{-1, 0, 1} {
			if r != 0 || c != 0 {
				allDirections = append(allDirections, direction{r, c})
			}
		}
	}
}

type boardIndex struct {
	sync.Mutex
	index map[string]struct{}
}
type board struct {
	row, column int
	grid        []byte
	validBoards *boardIndex
	wgChildren  *sync.WaitGroup
}

func main() {
	inputsBoard := readInput()
	var wg sync.WaitGroup
	wg.Add(len(inputsBoard))
	for _, input := range inputsBoard {
		input.wgChildren.Add(1) // for the initial board
		go func(b *board) {
			defer wg.Done()
			b.work()
		}(input)
	}
	wg.Wait()

	for _, input := range inputsBoard {
		input.wgChildren.Wait()
		fmt.Printf("%d\n", len(input.validBoards.index)-1)
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
		b.wgChildren = &sync.WaitGroup{}
		testCases = append(testCases, &b)
	}
	return testCases
}

type queen struct {
	row, column int
	board       *board
}

//walk in the given directions and return true if no other queen was meet
func (q *queen) walk(d []direction) bool {
nextDir:
	for _, dir := range d {
		r, c := q.row, q.column
		for {
			r += dir[0]
			c += dir[1]
			switch q.board.get(r, c) {
			case 'Q':
				return false
			case '#', '0':
				continue nextDir
			case '.':
				q.board.set(r, c, 'x') // later no need to try that cell with a queen
			}
		}
	}
	return true
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
			q := queen{row: i / b.column, column: i % b.column, board: b}
			if !q.walk(allDirections) {
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
			b.wgChildren.Add(1)
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

func (b *board) work() {
	defer b.wgChildren.Done()
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
			go child.work()
		}
	}
}
