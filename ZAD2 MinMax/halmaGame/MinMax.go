package halmaGame

import (
	"ZAD2_MinMax/utils"
)

func MoveSelection(board *Board, currentPlayer utils.Player) (*Pawn, *Move) {
	originalMoves, originalPawns := getMovesAndPawns(board, &currentPlayer)

	tree := NewNode(nil, board, 0, currentPlayer, nil, nil)
	channel := make(chan struct {
		*Node
		int8
	}, 1)
	tree.selectScore(channel, 0)

	bestKid := <-channel
	var bestKidNode *Node
	for _, node := range tree.Children {
		if node.score == bestKid.score {
			bestKidNode = node
			break
		}
	}
	//bestKidNode := tree.Children[bestKid.int8]
	for i, originalPawn := range originalPawns {
		if &originalPawn.Coords == bestKidNode.initialCords {
			return originalPawn, &originalMoves[i]
		}
	}

	return nil, nil
}
