package main

import (
	"ZAD2_MinMax/halmaGame"
	"ZAD2_MinMax/utils"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func handleExit(sig chan os.Signal) {
	<-sig
	utils.MoveCursor(0, utils.BOARD_HEIGHT+1)
	os.Exit(0)
}

func sendSignals(parentChannel chan int) {
	for i := 0; i < 5; i++ {
		parentChannel <- i
		time.Sleep(1 * time.Second)
	}
	print("All signals sent")
	close(parentChannel)
}

func findMin(parentChannel chan int) {
	min := 1000
	for i := range parentChannel {
		if i < min {
			min = i
		}
		print("Received: ", i, "\n")
	}
	print("Min: ", min)
}

func main() {
	//set max stack size
	debug.SetMaxStack(10_000_000_000_000)
	debug.SetGCPercent(10000)
	debug.SetMaxThreads(1000000000)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go handleExit(sig)

	game := halmaGame.Game{}
	game.RunGameComputerVSComputer()
}
