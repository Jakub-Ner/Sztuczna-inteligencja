package main

import (
	"ZAD2_MinMax/halmaGame"
	"ZAD2_MinMax/utils"
)

func main() {
	game := halmaGame.Game{}
	game.RunGame()

	utils.MoveCursor(0, utils.BOARD_HEIGHT+1)
}
