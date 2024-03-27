package minMax

import "ZAD2_MinMax/halmaGame"

type Node struct {
	halmaGame.Pawn
	Children []Node
}
