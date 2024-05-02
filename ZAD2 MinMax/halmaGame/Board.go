package halmaGame

import (
	"ZAD2_MinMax/utils"
	"fmt"
)

type BoardObject interface {
	GetIcon() string
	CanMoveHere() bool
}

type Board struct {
	Fields        [utils.COLUMNS][utils.COLUMNS]BoardObject
	Pawns         [utils.PAWNS]*Pawn
	CurrentPlayer utils.Player
}

func (b *Board) ResetPawns() {
	var pawn string
	pawn = utils.PLAYER_GREEN
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

	pawn = utils.PLAYER_YELLOW
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

func (b *Board) UpdateMoves() error {
	//fmt.Println("Try on board:", &b, "update moves:", utils.GetGoroutine())
	fmt.Printf("board=%p %s", b, utils.GetGoroutine())
	for _, pawn := range b.Pawns {
		//fmt.Println("Try pawn:", pawn.String(), utils.GetGoroutine())
		pawn.ValidMoves = nil
		potentialMoves := getNeighbour(pawn.Coords.X, pawn.Coords.Y, 0)
		//fmt.Println("Potential moves: ", potentialMoves, utils.GetGoroutine())
		for _, move := range potentialMoves {
			//println()
			if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() && move.NumberOfJumps == 0 {
				//fmt.Printf("1)pawn=%p %s", &pawn, utils.GetGoroutine())
				move.NumberOfJumps = 0
				m := move
				fmt.Println("Move:", m, "will be added to", pawn.ValidMoves, utils.GetGoroutine())
				pawn.ValidMoves = append(pawn.ValidMoves, m)
				//fmt.Print(2, utils.GetGoroutine())
			} else {
				if canJump(pawn.Coords, &move.Direction, *b) {
					//fmt.Printf("3)pawn=%p %s", &pawn, utils.GetGoroutine())
					move.NumberOfJumps = +1
					m := move
					fmt.Println("Move: ", m, "will be added to", pawn.ValidMoves, utils.GetGoroutine())
					pawn.ValidMoves = append(pawn.ValidMoves, m)
					//fmt.Print(4, utils.GetGoroutine())
					newNeighbours := getNeighbour(move.Direction.X, move.Direction.Y, move.NumberOfJumps)
					//fmt.Print(5, utils.GetGoroutine())
					potentialMoves = append(potentialMoves, newNeighbours...)
					//fmt.Print(6, utils.GetGoroutine())
				}
			}
		}
	}
	//fmt.Println("Done on board:", &b, "update moves:", utils.GetGoroutine())
	return nil
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
	if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() == false {
		defer func() {
			pawn.ValidMoves = nil // clear the valid moves
		}()
		fmt.Println("Invalid move", pawn.String(), move.Direction.X, move.Direction.Y, "\n")
		return
	}

	b.Fields[pawn.Coords.Y][pawn.Coords.X] = EmptyField{utils.EMPTY_BLUE}
	pawn.Coords = move.Direction
	b.Fields[pawn.Coords.Y][pawn.Coords.X] = pawn
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
