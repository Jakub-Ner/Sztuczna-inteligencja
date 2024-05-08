package halmaGame

import (
	"ZAD2_MinMax/utils"
	"fmt"
)

func printColumnNumbers() {
	print("   ")
	for i := int8(0); i < utils.COLUMNS; i++ {
		print(fmt.Sprintf("%02d ", i))
	}
	println()
}

func (b *Board) selectPawn(currentPlayer utils.Player) *Pawn {
	for {
		x, y := utils.TakeInput("Select a pawn to move")
		pawn, ok := b.Fields[y][x].(*Pawn)
		if ok && pawn.Icon == currentPlayer {
			return pawn
		}
		//utils.PrintMessage("Invalid pawn selected")
	}
}

func displayMoves(pawn Pawn) {
	utils.MoveCursor(3*int(pawn.Coords.X)+4, int(pawn.Coords.Y)+2)
	print("🟣 ")
	for _, move := range pawn.ValidMoves {
		utils.MoveCursor(3*int(move.Direction.X)+4, int(move.Direction.Y)+2)
		print("🟪 ")
	}
}

func (b *Board) Print() {
	utils.ResetDisplay()
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
}
