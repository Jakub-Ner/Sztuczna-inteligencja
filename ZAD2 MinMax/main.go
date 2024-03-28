package main

import (
	"ZAD2_MinMax/halmaGame"
	"ZAD2_MinMax/utils"
	"os"
	"os/signal"
	"syscall"
)

func handleExit(sig chan os.Signal) {
	<-sig
	utils.MoveCursor(0, utils.BOARD_HEIGHT+1)
	os.Exit(0)
}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go handleExit(sig)

	game := halmaGame.Game{}
	game.RunGame()
}
