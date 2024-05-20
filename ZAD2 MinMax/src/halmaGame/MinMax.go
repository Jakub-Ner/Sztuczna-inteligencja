package halmaGame

import (
	"ZAD2_MinMax/utils"
	"math"
	"strconv"
)

func initMinimax() {
	_visitedNodesCounter = 0
	_prunedNodesCounter = 0

	_visitedon1Level = 0
	_visitedon2Level = 0
	_visitedon3Level = 0
}

func calculateStats() {
	_allVisitedNodesCounter += _visitedNodesCounter
	_allPrunedNodesCounter += _prunedNodesCounter
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
func MoveSelection(board *Board, currentPlayer utils.Player, heuristic Heuristic) (*Pawn, *Move) {
	defer func() {
		calculateStats()
	}()

	_heuristic = heuristic
	initMinimax()

	tree := NewNode(nil, nil)

	utils.PrintMessage("Turn: " + strconv.Itoa(_turnCounter/2) + " Player: " + currentPlayer)
	//score := int(utils.CountAndShowTime(runMinimax, tree, board, currentPlayer)[0].Int())
	score := int(utils.CountAndShowTime(runMiniMaxAlphaBeta, tree, board, currentPlayer)[0].Int())

	//score := runMiniMaxAlphaBeta(tree, board, currentPlayer)
	//score := runMinimax(tree, board, currentPlayer)
	_turnCounter++

	var bestKidNode *Node

	for _, node := range tree.Children {
		if node.score == score {
			bestKidNode = node
			break
		}
	}

	for _, originalPawn := range board.Pawns {
		if &originalPawn.Coords == bestKidNode.initialCords {
			utils.PrintMessage("Best Score: " + strconv.Itoa(score) + "; Visited nodes: " + strconv.Itoa(int(_visitedNodesCounter)))
			utils.PrintMessage("Pruned nodes: " + strconv.FormatInt(_prunedNodesCounter, 10))
			//utils.PrintMessage("On 1st level: " + strconv.Itoa(int(_visitedon1Level)) + "; On 2nd level: " + strconv.Itoa(int(_visitedon2Level)) + "; On 3rd level: " + strconv.Itoa(int(_visitedon3Level)))
			return originalPawn, &Move{0, bestKidNode.pawn.Coords}
		}
	}

	return nil, nil
}
