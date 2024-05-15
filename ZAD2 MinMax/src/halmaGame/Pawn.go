package halmaGame

import (
	"ZAD2_MinMax/utils"
)

type Pawn struct {
	Icon       utils.Player
	ValidMoves []Move
	Coords     Coords
}

func NewPawn(icon string, coords Coords) *Pawn {
	return &Pawn{
		Icon:       icon,
		ValidMoves: []Move{},
		Coords:     coords,
	}
}

func (p Pawn) CanMoveHere() bool {
	return false
}

func (p Pawn) GetIcon() string {
	return p.Icon
}

func (p Pawn) String() string {
	return p.Icon
}
