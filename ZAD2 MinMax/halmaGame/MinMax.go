package halmaGame

import (
	"ZAD2_MinMax/utils"
)

func MoveSelection(board *Board, currentPlayer utils.Player) (*Pawn, *Move) {
	//currentPlayer =
	originalMoves, originalPawns := getMovesAndPawns(board, &currentPlayer)

	tree := NewNode(nil, *board, 0, 0, getOpponent(currentPlayer), nil, nil)
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
	//println("Best score has: ", bestKid.pawn.String(), bestKid.move.String())

	//return bestKid.pawn, bestKid.
	return nil, nil
}
