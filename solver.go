package sudoku_go

import (
	"fmt"
)

type Solver struct {
	puzzle     Puzzle
	rowsHash   map[int][]bool
	colsHash   map[int][]bool
	regionHash [][][]bool // 3x3x9, row, col, num
}

func NewSolver(puzzle Puzzle) Solver {
	rowsHash := make(map[int][]bool)
	colsHash := make(map[int][]bool)
	regionHash := make([][][]bool, 3)

	for i := 0; i < 9; i++ {
		rowsHash[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		colsHash[i] = make([]bool, 9)
	}

	for i := range regionHash {
		regionHash[i] = make([][]bool, 3)
		for j := range regionHash[i] {
			regionHash[i][j] = make([]bool, 9)
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if val := puzzle.board[row][col]; val != 0 {
				rowsHash[row][val-1] = true
				colsHash[col][val-1] = true
				regionHash[row/3][col/3][val-1] = true
			}
		}
	}
	return Solver{
		puzzle:     puzzle,
		rowsHash:   rowsHash,
		colsHash:   colsHash,
		regionHash: regionHash,
	}

}

func (s *Solver) Solve() {
	fmt.Println("Solving: ")
	s.Print()
	fmt.Println()
	s.solveHelper(0, 0)
	fmt.Println("Solution: ")
	s.Print()
}

func (s *Solver) solveHelper(row, col int) bool {
	if col >= 9 {
		col = 0
		row++
	}
	if row >= 9 {
		return true
	}

	num, err := s.puzzle.getNum(row, col)
	if err != nil {
		panic(err.Error())
	}

	// if populated
	if num != 0 {
		return s.solveHelper(row, col+1)
	}

	// try from 1 to 9
	for i := 1; i <= 9; i++ {
		if s.isAvailable(row, col, i) {
			s.setNum(row, col, i)
			if s.solveHelper(row, col+1) {
				return true
			}
			s.unsetNum(row, col)
		}
	}
	return false
}

func (s *Solver) isAvailable(row, col, val int) bool {
	valIndex := val - 1
	if s.rowsHash[row][valIndex] {
		return false
	}

	if s.colsHash[col][valIndex] {
		return false
	}

	if s.regionHash[row/3][col/3][valIndex] {
		return false
	}

	return true
}

func (s *Solver) unsetNum(row int, col int) {
	originalVal, err := s.puzzle.getNum(row, col)
	if err != nil {
		panic(err.Error())
	}
	if originalVal != 0 {
		s.rowsHash[row][originalVal-1] = false
		s.colsHash[col][originalVal-1] = false
		s.regionHash[row/3][col/3][originalVal-1] = false
		s.setNum(row, col, 0)
	}
}

func (s *Solver) setNum(row int, col int, val int) {
	s.puzzle.board[row][col] = val
	if val != 0 {
		s.rowsHash[row][val-1] = true
		s.colsHash[col][val-1] = true
		s.regionHash[row/3][col/3][val-1] = true
	}
}

func (s *Solver) Print() {
	for row := 0; row < 9; row++ {
		fmt.Println(s.puzzle.board[row])
	}
}
