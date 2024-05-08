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
	pawnCounter := 0

	boardPattern := utils.ReadBoardFromFile("boards/initial_board.txt")
	for i := int8(0); i < utils.COLUMNS; i++ {
		for j := int8(0); j < utils.COLUMNS; j++ {
			if boardPattern[i][j] == 1 {
				p := NewPawn(utils.PLAYER_YELLOW, Coords{j, i})
				b.Fields[i][j] = p
				b.Pawns[pawnCounter] = p
				pawnCounter++
			} else if boardPattern[i][j] == 2 {
				p := NewPawn(utils.PLAYER_GREEN, Coords{j, i})
				b.Fields[i][j] = p
				b.Pawns[pawnCounter] = p
				pawnCounter++
			}
		}
	}
	b.UpdateMoves()
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

func canJump(from Coords, to *Coords, b Board) bool {
	diffX := to.X - from.X
	diffY := to.Y - from.Y
	to.X, to.Y = to.X+diffX, to.Y+diffY
	return validateMove(to.X, to.Y) && b.Fields[to.Y][to.X].CanMoveHere()
}

func (b *Board) UpdateMoves() {
	for _, pawn := range b.Pawns {
		if pawn == nil {
			continue
		}
		pawn.ValidMoves = nil
		potentialMoves := getNeighbour(pawn.Coords.X, pawn.Coords.Y, 0)
		for _, move := range potentialMoves {
			if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() && move.NumberOfJumps == 0 {
				move.NumberOfJumps = 0
				m := move
				pawn.ValidMoves = append(pawn.ValidMoves, m)
			} else {
				if canJump(pawn.Coords, &move.Direction, *b) {
					move.NumberOfJumps = +1
					m := move
					pawn.ValidMoves = append(pawn.ValidMoves, m)
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
	//utils.PrintMessage("Invalid move")
	return nil
}

func (b *Board) MovePawn(pawn *Pawn, move Move) bool {
	if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() == false {
		defer func() {
			pawn.ValidMoves = nil // clear the valid moves
		}()
		//fmt.Println("Invalid move", pawn.String(), move.Direction.X, move.Direction.Y, "\n")
		return false
	}

	b.Fields[pawn.Coords.Y][pawn.Coords.X] = EmptyField{utils.EMPTY_BLUE}
	pawn.Coords = move.Direction
	b.Fields[pawn.Coords.Y][pawn.Coords.X] = pawn
	return true
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
