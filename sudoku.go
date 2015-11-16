package sudoku

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	Size     int
	Board    [][]int
	maxDigit int
}

func getDigit(value int) int {
	ret := 0

	if value == 0 {
		return 1
	}
	for value > 0 {
		value /= 10
		ret++
	}
	return ret
}

func Parse(in io.Reader) (*Board, error) {
	board := Board{}
	boardInt := make([]int, 0, 512)

	content, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	tab := strings.Split(string(content), "\n")
	tab = tab[:len(tab)-1]
	if len(tab) < 2 {
		return nil, fmt.Errorf("col must be greater than 2")
	}
	lastLength := 0
	for y, line := range tab {
		tabx := strings.Fields(line)
		if y == 0 {
			lastLength = len(tabx)
			if lastLength < 2 {
				return nil, fmt.Errorf("row must be greater than 2")
			}
			if lastLength != len(tab) {
				return nil, fmt.Errorf("columns and row must be equal")
			}
		}
		if len(tabx) != lastLength {
			return nil, fmt.Errorf("the columns must be equal")
		}
		for _, number := range tabx {
			value, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			boardInt = append(boardInt, value)
		}
	}
	board.Size = lastLength
	board.Board = make([][]int, lastLength)
	board.maxDigit = getDigit(boardInt[0])
	for y := range board.Board {
		board.Board[y] = make([]int, lastLength)
		for x := range board.Board[y] {
			board.Board[y][x] = boardInt[y*lastLength+x]
			digit := getDigit(board.Board[y][x])
			if digit > board.maxDigit {
				board.maxDigit = digit
			}
		}
	}
	return &board, nil
}

func (b *Board) Solve() error {
	return nil
}

func stringWithSpace(value, max int) string {
	ret := strconv.Itoa(value)
	digit := getDigit(value)
	for digit < max {
		ret += " "
		digit++
	}
	return ret
}

func (b *Board) Print(out io.Writer) {
	for _, y := range b.Board {
		for n, x := range y {
			fmt.Fprintf(out, "%s", stringWithSpace(x, b.maxDigit))
			if n != b.Size {
				fmt.Fprintf(out, " ")
			}
		}
		fmt.Println()
	}
}
