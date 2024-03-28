package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var LineCounter = 0

func PrintMessage(message string) {
	MoveCursor(BOARD_WIDTH+DISPLAY_MARGIN, DISPLAY_MARGIN+LineCounter)
	fmt.Println(message)
	LineCounter++
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
	x, _ := strconv.ParseInt(coords[0], 10, 8)
	y, _ := strconv.ParseInt(coords[1], 10, 8)

	return int8(x), int8(y)
}
