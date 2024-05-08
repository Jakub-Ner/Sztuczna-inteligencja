package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

var LineCounter = 0

func PrintMessage(message string) {
	MoveCursor(BOARD_WIDTH+DISPLAY_MARGIN, DISPLAY_MARGIN+LineCounter)
	fmt.Println(message)
	LineCounter++
}

func GetGoroutine() string {
	return " Go=" + string(bytes.Fields(debug.Stack())[1]) + " "
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func ResetDisplay() {
	MoveCursor(0, 0)
	LineCounter = 0
	print("\033[H\033[2J")
}

func TakeInput(message string) (int8, int8) {
	PrintMessage(message)

	MoveCursor(BOARD_WIDTH+DISPLAY_MARGIN, DISPLAY_MARGIN+LineCounter)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	coords := strings.Split(input, " ")
	if len(coords) != 2 {
		PrintMessage("Invalid input")
		return TakeInput(message)
	}

	x, err1 := strconv.ParseInt(coords[0], 10, 8)
	y, err2 := strconv.ParseInt(coords[1], 10, 8)

	if err1 != nil || err2 != nil {
		PrintMessage("Invalid input")
		return TakeInput(message)
	}
	return int8(x), int8(y)
}

func ReadBoardFromFile(filename string) [COLUMNS][COLUMNS]int8 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boardPattern := [COLUMNS][COLUMNS]int8{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		for i, char := range line {
			boardPattern[LineCounter][int8(i)] = int8(char - '0')
		}
		LineCounter++
	}
	return boardPattern
}
