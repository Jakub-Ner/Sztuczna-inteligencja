package halmaGame

import (
	"ZAD2_MinMax/utils"
)

func MoveSelection(board *Board, currentPlayer utils.Player) (*Pawn, *Move) {
	originalMoves, originalPawns := getMovesAndPawns(board, &currentPlayer)

	tree := NewNode(nil, *board, 0, currentPlayer, nil, nil)
	channel := make(chan *Node, 1)
	tree.selectScore(channel)

	bestKid := <-channel
	for _, kid := range tree.Children {
		if kid.score == bestKid.score {
			for i, originalPawn := range originalPawns {
				if &originalPawn.Coords == kid.initialCords {
					return originalPawn, &originalMoves[i]
				}
			}
		}
	}
	return nil, nil
}
