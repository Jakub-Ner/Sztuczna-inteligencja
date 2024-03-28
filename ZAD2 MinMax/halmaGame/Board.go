package halmaGame

import (
	"ZAD2_MinMax/utils"
)

type BoardObject interface {
	GetIcon() string
	CanMoveHere() bool
}

type Board struct {
	Fields [utils.COLUMNS][utils.COLUMNS]BoardObject
	Pawns  [utils.PAWNS]*Pawn
}

func (b *Board) ResetPawns() {
	var pawn string
	pawn = utils.GREEN_PAWN
	pawnCounter := 0
	for i := int8(0); i <= 6; i++ {
		for j := int8(5); j >= i; j-- {
			if i == 0 && j == 0 {
				b.Fields[0][10] = EmptyField{utils.EMPTY_RED}
			} else if i == 5 && j == 5 {
				b.Fields[5][15] = EmptyField{utils.EMPTY_RED}
			} else {
				p := NewPawn(pawn, Coords{j + 10, i})
				b.Fields[i][j+10] = p
				b.Pawns[pawnCounter] = p
				pawnCounter++
			}

		}
	}

	pawn = utils.YELLOW_PAWN
	for i := int8(15); i >= 10; i-- {
		for j := int8(10); j <= i; j++ {
			if i == 10 && j == 10 {
				b.Fields[10][0] = EmptyField{utils.EMPTY_BLUE}
			} else if i == 15 && j == 15 {
				b.Fields[15][5] = EmptyField{utils.EMPTY_BLUE}
			} else {
				p := NewPawn(pawn, Coords{j - 10, i})
				b.Fields[i][j-10] = p
				b.Pawns[pawnCounter] = p
				pawnCounter++
			}
		}
	}
	b.initMoves()
}

func (b *Board) Reset() {
	// Fill the halmaGame with empty fields
	for i := int8(0); i < utils.COLUMNS; i++ {
		for j := int8(0); j < utils.COLUMNS; j++ {
			if i%2 == j%2 {
				b.Fields[i][j] = EmptyField{utils.EMPTY_RED}
			} else {
				b.Fields[i][j] = EmptyField{utils.EMPTY_BLUE}
			}
		}
	}
	b.ResetPawns()
}

func NewBoard() *Board {
	board := &Board{}
	board.Reset()
	return board
}

func validateMove(x, y int8) bool {
	return x >= 0 && x < utils.COLUMNS && y >= 0 && y < utils.COLUMNS
}

func canJump(from Coords, to Coords, b Board) bool {
	diffX := to.X - from.X
	diffY := to.Y - from.Y
	x, y := to.X+diffX, to.Y+diffY
	return validateMove(x, y) && b.Fields[y][x].CanMoveHere()
}

func (b *Board) initMoves() {
	for _, pawn := range b.Pawns {
		potentialMoves := getNeighbour(pawn.Coords.X, pawn.Coords.Y, 0)
		for _, move := range potentialMoves {
			if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() && move.NumberOfJumps == 0 {
				move.NumberOfJumps = 0
				pawn.ValidMoves = append(pawn.ValidMoves, move)
			} else {
				if canJump(pawn.Coords, move.Direction, *b) {
					move.NumberOfJumps = +1
					pawn.ValidMoves = append(pawn.ValidMoves, move)
					newNeighbours := getNeighbour(move.Direction.X, move.Direction.Y, move.NumberOfJumps)
					potentialMoves = append(potentialMoves, newNeighbours...)
				}
			}
		}
	}
	return
}

func (b *Board) selectMove(pawn *Pawn) *Move {
	x, y := utils.TakeInput("Select Your move")
	for _, move := range pawn.ValidMoves {
		if move.Direction.X == x && move.Direction.Y == y {
			return &move
		}
	}
	utils.PrintMessage("Invalid move")
	return nil
}

func (b *Board) MovePawn(pawn *Pawn, move Move) {
	b.Fields[pawn.Coords.Y][pawn.Coords.X] = EmptyField{utils.EMPTY_BLUE}
	pawn.Coords = move.Direction
	b.Fields[pawn.Coords.Y][pawn.Coords.X] = pawn
	pawn.ValidMoves = nil
}

func getNeighbour(x, y, numberOfJumps int8) []Move {
	var moves []Move
	for xi := x - 1; xi <= x+1; xi++ {
		for yi := y - 1; yi <= y+1; yi++ {
			if validateMove(xi, yi) && (xi != x || yi != y) {
				moves = append(moves, Move{
					Direction:     Coords{X: xi, Y: yi},
					NumberOfJumps: numberOfJumps,
				})
			}
		}
	}
	return moves
}
