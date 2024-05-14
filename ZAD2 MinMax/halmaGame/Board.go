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
	greenPawnCounter := 0
	yellowPawnCounter := utils.PAWNS_PER_PLAYER

	boardPattern := utils.ReadBoardFromFile("boards/initial_board.txt")
	for i := int8(0); i < utils.COLUMNS; i++ {
		for j := int8(0); j < utils.COLUMNS; j++ {
			if boardPattern[i][j] == 1 {
				p := NewPawn(utils.PLAYER_YELLOW, Coords{j, i})
				b.Fields[i][j] = p
				b.Pawns[yellowPawnCounter] = p
				yellowPawnCounter++
			} else if boardPattern[i][j] == 2 {
				p := NewPawn(utils.PLAYER_GREEN, Coords{j, i})
				b.Fields[i][j] = p
				b.Pawns[greenPawnCounter] = p
				greenPawnCounter++
			}
		}
	}
	b.UpdateMoves()
}

func (b *Board) Reset() {
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

func getDirection(from, to int8) int8 {
	if from < to {
		return 1
	} else if from > to {
		return -1
	}
	return 0
}
func canJump(from Coords, to *Coords, b Board) bool {
	to.X, to.Y = to.X+getDirection(from.X, to.X), to.Y+getDirection(from.Y, to.Y)
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
			if b.Fields[move.Direction.Y][move.Direction.X].CanMoveHere() {
				move.NumberOfJumps = 0
				m := move
				pawn.ValidMoves = append(pawn.ValidMoves, m)
			} else {
				b.checkJumpMoves(pawn, pawn.Coords, move, 0)
			}
		}
	}
	return
}

func isUniqueMove(moves []Move, move Move) bool {
	for _, m := range moves {
		if m.Direction.X == move.Direction.X && m.Direction.Y == move.Direction.Y {
			return false
		}
	}
	return true
}

func (b *Board) checkJumpMoves(pawn *Pawn, from Coords, move Move, ttl int8) {
	if ttl == 5 {
		return
	}
	if canJump(from, &move.Direction, *b) {
		move.NumberOfJumps = +1
		m := move
		if isUniqueMove(pawn.ValidMoves, m) {
			pawn.ValidMoves = append(pawn.ValidMoves, m)
		}

		from = move.Direction
		for _, neighbour := range getNeighbour(move.Direction.X, move.Direction.Y, move.NumberOfJumps) {
			if b.Fields[neighbour.Direction.Y][neighbour.Direction.X].CanMoveHere() == false {
				b.checkJumpMoves(pawn, from, neighbour, ttl+1)
			}
		}
	}
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
		return false
	}

	emptyFieldColor := utils.EMPTY_RED
	if pawn.Coords.Y%2 != pawn.Coords.X%2 {
		emptyFieldColor = utils.EMPTY_BLUE
	}

	b.Fields[pawn.Coords.Y][pawn.Coords.X] = EmptyField{emptyFieldColor}
	pawn.Coords.X = move.Direction.X
	pawn.Coords.Y = move.Direction.Y
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
