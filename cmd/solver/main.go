package main

import (
	"os"

	"github.com/QuentinPerez/NSudoku"
	"github.com/Sirupsen/logrus"
)

func main() {
	var err error

	in := os.Stdin
	if len(os.Args) == 2 {
		in, err = os.Open(os.Args[1])
		if err != nil {
			logrus.Fatal(err)
		}
	}
	board, err := sudoku.Parse(in)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := board.Solve(); err != nil {
		logrus.Fatal(err)
	}
	board.Print(os.Stdout)
}
