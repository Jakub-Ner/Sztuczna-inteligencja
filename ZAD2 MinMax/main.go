package main

import (
	"ZAD2_MinMax/halmaGame"
	"ZAD2_MinMax/utils"
	"math"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func handleExit(sig chan os.Signal) {
	<-sig
	utils.MoveCursor(0, utils.BOARD_HEIGHT+1)
	os.Exit(0)
}

func testBoardCopy() {
	board := halmaGame.NewBoard()
	board.Print()
	boardCopy := new(halmaGame.Board)
	*boardCopy = *board

	board.MovePawn(board.Pawns[19], halmaGame.Move{0, halmaGame.Coords{2, 2}})
	board.Print()
	boardCopy.Print()

	boardCopy.MovePawn(board.Pawns[19], halmaGame.Move{0, halmaGame.Coords{5, 2}})
	board.Print()
	boardCopy.Print()
}

func main() {
	//testBoardCopy()

	debug.SetMaxStack(math.MaxInt64)
	//debug.SetGCPercent(10000)
	debug.SetMemoryLimit(math.MaxInt64)
	debug.SetMaxThreads(math.MaxInt64)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go handleExit(sig)

	game := halmaGame.Game{}
	game.RunGameComputerVSComputer()
}
