package halmaGame

import (
	"ZAD2_MinMax/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DISPLAY_MARGIN = 5

func printColumnNumbers() {
	print("   ")
	for i := int8(0); i < utils.COLUMNS; i++ {
		print(fmt.Sprintf("%02d ", i))
	}
	println()
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

var lineCounter = 0

func PrintMessage(message string) {
	MoveCursor(BOARD_WIDTH+DISPLAY_MARGIN, DISPLAY_MARGIN+lineCounter)
	fmt.Println(message)
	lineCounter++
}
func selectPawn(b Board) *Pawn {
	PrintMessage("Select a Pawn to move")

	MoveCursor(BOARD_WIDTH+DISPLAY_MARGIN, DISPLAY_MARGIN+lineCounter)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	coords := strings.Split(input, " ")
	x, _ := strconv.ParseInt(coords[0], 10, 8)
	y, _ := strconv.ParseInt(coords[1], 10, 8)

	pawn := b.Fields[y][x]
	return pawn.(*Pawn)
}

func resetDisplay() {
	MoveCursor(0, 0)
	lineCounter = 0
	print("\033[H\033[2J")
}

func displayMoves(pawn Pawn) {
	for _, move := range pawn.ValidMoves {
		MoveCursor(3*int(move.Direction.X)+4, int(move.Direction.Y)+2)
		print("🟪 ")
	}
}

func (b *Board) Print() {
	resetDisplay()
	printColumnNumbers()
	for i := int8(0); i < utils.COLUMNS; i++ {
		// print the row number, formatted to 2 digits
		print(fmt.Sprintf("%02d ", i))
		for j := int8(0); j < utils.COLUMNS; j++ {
			print(b.Fields[i][j].GetIcon() + " ")
		}
		println(fmt.Sprintf("%02d ", i))
	}
	printColumnNumbers()

	selectedPawn := selectPawn(*b)
	displayMoves(*selectedPawn)
}
