package halmaGame

import (
	"ZAD2_MinMax/utils"
	"math"
	"strconv"
	"time"
)

var turnCounter = 0

func initMinimax() {

	visitedNodesCounter = 0
	prunedNodesCounter = 0

	visitedon1Level = 0
	visitedon2Level = 0
	visitedon3Level = 0
}

func runMinimax(tree *Node, board *Board, currentPlayer utils.Player) int {
	channel := make(chan struct {
		*Node
		int8
	}, 1)

	tree.selectScore(channel, 0, board, 0, currentPlayer)
	bestKid := <-channel
	return bestKid.score
}

func runMiniMaxAlphaBeta(tree *Node, board *Board, currentPlayer utils.Player) int {
	score := tree.selectScoreAlphaBeta(board, 0, currentPlayer, math.MinInt, math.MaxInt)
	return score
}
func MoveSelection(board *Board, currentPlayer utils.Player) (*Pawn, *Move) {
	heuristic = DistanceScore
	initMinimax()

	tree := NewNode(nil, nil)

	utils.PrintMessage("Turn: " + strconv.Itoa(turnCounter/2) + " Player: " + currentPlayer)
	before := time.Now()
	//score := int(utils.CountAndShowTime(runMinimax, tree, board, currentPlayer)[0].Int())
	score := int(utils.CountAndShowTime(runMiniMaxAlphaBeta, tree, board, currentPlayer)[0].Int())
	after := time.Now()
	utils.PrintMessage("Finished minmax in: " + strconv.FormatInt(after.Sub(before).Milliseconds(), 10) + "ms")
	turnCounter++

	var bestKidNode *Node

	checkScore := 0
	for _, node := range tree.Children {
		if node.score == score {
			checkScore = node.score
			bestKidNode = node
			break
		}
	}
	if score != checkScore {
		panic("score is wrong!!!!")
	}

	for _, originalPawn := range board.Pawns {
		if &originalPawn.Coords == bestKidNode.initialCords {
			utils.PrintMessage("Best Score: " + strconv.Itoa(score) + "; Visited nodes: " + strconv.Itoa(int(visitedNodesCounter)))
			utils.PrintMessage("Pruned nodes: " + strconv.FormatInt(prunedNodesCounter, 10))
			utils.PrintMessage("On 1st level: " + strconv.Itoa(int(visitedon1Level)) + "; On 2nd level: " + strconv.Itoa(int(visitedon2Level)) + "; On 3rd level: " + strconv.Itoa(int(visitedon3Level)))
			return originalPawn, &Move{0, bestKidNode.pawn.Coords}
		}
	}

	return nil, nil
}
