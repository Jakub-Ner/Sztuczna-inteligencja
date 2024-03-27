package main

import "ZAD2_MinMax/halmaGame"

func main() {
	board := halmaGame.NewBoard()
	board.Print()

	halmaGame.MoveCursor(0, halmaGame.BOARD_HEIGHT+1)
}
