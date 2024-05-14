package halmaGame

import (
	"ZAD2_MinMax/utils"
	"math"
	"strconv"
)

var turnCounter = 0

func initMinimax() {
	alpha = math.MinInt
	beta = math.MaxInt

	visitedNodesCounter = 0
	prunedNodesCounter = 0

	visitedon1Level = 0
	visitedon2Level = 0
	visitedon3Level = 0
}
func MoveSelection(board *Board, currentPlayer utils.Player) (*Pawn, *Move) {
	heuristic = DistanceScore
	initMinimax()

	tree := NewNode(nil, nil)
	channel := make(chan struct {
		*Node
		int8
	}, 1)
	utils.PrintMessage("Turn: " + strconv.Itoa(turnCounter/2) + " Player: " + currentPlayer)
	tree.selectScore(channel, 0, board, 0, currentPlayer)
	turnCounter++

	bestKid := <-channel
	var bestKidNode *Node

	checkScore := 0
	for _, node := range tree.Children {
		if node.score == bestKid.score {
			checkScore = node.score
			bestKidNode = node
			break
		}
	}
	if bestKid.score != checkScore {
		panic("score is wrong!!!!")
	}

	//for _, node := range tree.Children {
	//	if node.score == bestKid.score {
	//		bestKidNode = node
	//		break
	//	}
	//}
	for _, originalPawn := range board.Pawns {
		if &originalPawn.Coords == bestKidNode.initialCords {
			utils.PrintMessage("Best Score: " + strconv.Itoa(bestKid.score) + "; Visited nodes: " + strconv.Itoa(int(visitedNodesCounter)))
			utils.PrintMessage("Pruned nodes: " + strconv.FormatInt(prunedNodesCounter, 10))
			utils.PrintMessage("On 1st level: " + strconv.Itoa(int(visitedon1Level)) + "; On 2nd level: " + strconv.Itoa(int(visitedon2Level)) + "; On 3rd level: " + strconv.Itoa(int(visitedon3Level)))
			return originalPawn, &Move{0, bestKidNode.pawn.Coords}
		}
	}

	return nil, nil
}
