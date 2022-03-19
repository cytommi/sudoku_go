package sudoku_go

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Board [][]int

type Puzzle struct {
	board Board
}

func NewPuzzle(filePath string) (*Puzzle, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Error when opening the puzzle file: " + err.Error())
	}
	defer file.Close()

	var board [][]int

	scanner := bufio.NewScanner(file)
	rows := 0
	for scanner.Scan() {
		rows += 1
		numsInLine := strings.Split(scanner.Text(), "")

		if len(numsInLine) != 9 {
			return nil, errors.New("Invalid puzzle")
		}

		var row []int
		for _, num := range numsInLine {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}

		board = append(board, row)
	}

	if rows != 9 {
		return nil, errors.New("Invalid puzzle")
	}

	puzzle := new(Puzzle)
	puzzle.board = board
	return puzzle, nil
}

func (p *Puzzle) getNum(row, col int) (int, error) {
	if row > 8 || col > 8 || row < 0 || col < 0 {
		return -1, errors.New("Invalid col, row")
	}

	return p.board[row][col], nil
}
