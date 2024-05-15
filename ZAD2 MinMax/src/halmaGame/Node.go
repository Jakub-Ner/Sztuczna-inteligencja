package halmaGame

import (
	"ZAD2_MinMax/utils"
	"fmt"
	"math"
	"sync/atomic"
)

type Node struct {
	Children     []*Node
	score        int
	pawn         *Pawn
	initialCords *Coords
}

func getMovesAndPawns(board *Board, currentPlayer *utils.Player) ([]Move, []*Pawn) {
	var startingPawnIdx int8
	if *currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	} else {
		startingPawnIdx = utils.PAWNS_PER_PLAYER
	}
	board.UpdateMoves()

	moves := make([]Move, 0, 40)
	pawns := make([]*Pawn, 0, 40)
	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		pawn := board.Pawns[i]
		for _, move := range pawn.ValidMoves {
			moves = append(moves, move)
			pawns = append(pawns, pawn)
		}
	}
	return moves, pawns
}

func isOpponent(nestingLevel int8) bool {
	return nestingLevel%2 == 1
}

func NewNode(pawn *Pawn, initialCords *Coords) *Node {
	return &Node{
		Children:     make([]*Node, 0),
		score:        0,
		pawn:         pawn,
		initialCords: initialCords,
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("score: %d, kids: %d", node.score, len(node.Children))
}

func (node *Node) selectScore(parentChannel chan struct {
	*Node
	int8
}, kidIdx int8, board *Board, nestingLevel int8, currentPlayer utils.Player) {
	defer func() {
		atomic.AddInt64(&_visitedNodesCounter, 1)
		if nestingLevel == 1 {
			atomic.AddInt64(&_visitedon1Level, 1)
		} else if nestingLevel == 2 {
			atomic.AddInt64(&_visitedon2Level, 1)
		} else if nestingLevel == 3 {
			atomic.AddInt64(&_visitedon3Level, 1)
		}
	}()
	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}

	if nestingLevel == utils.DEPTH {
		opponent := currentPlayer
		if isOpponent(nestingLevel) {
			opponent = getOpponent(currentPlayer)
		}
		node.score = _heuristic(board, opponent)
		parentChannel <- struct {
			*Node
			int8
		}{node, kidIdx}
		return
	}
	moves, pawns := getMovesAndPawns(board, &currentPlayer)

	channelForKids := make(chan struct {
		*Node
		int8
	}, len(moves))

	currentPawnId := 0
	for i := range moves {
		currentBoard := new(Board)
		*currentBoard = *board

		currentBoard.Pawns = [utils.PAWNS]*Pawn{}
		for i := int8(0); i < utils.PAWNS; i++ {
			currentBoard.Pawns[i] = new(Pawn)
			*currentBoard.Pawns[i] = *board.Pawns[i]
		}
		var currentPawn *Pawn
		for id := currentPawnId; id < len(currentBoard.Pawns); id++ {
			if pawns[i].Coords.X == currentBoard.Pawns[id].Coords.X && pawns[i].Coords.Y == currentBoard.Pawns[id].Coords.Y {
				currentPawn = currentBoard.Pawns[id]
				currentPawnId = id
				break
			}
		}

		initialPosition := &pawns[i].Coords
		//if !done {
		//	continue
		//}
		newNode := NewNode(currentPawn, initialPosition)
		node.Children = append(node.Children, newNode)
		currentBoard.MovePawn(currentPawn, moves[i])
		if validateResult(board) != "" {
			node.score = 100_000
			parentChannel <- struct {
				*Node
				int8
			}{node, kidIdx}
			return
		}
		if true { // !isOpponent(nestingLevel)
			go newNode.selectScore(channelForKids, int8(i), currentBoard, nestingLevel+1, getOpponent(currentPlayer))
		} else {
			newNode.selectScore(channelForKids, int8(i), currentBoard, nestingLevel+1, getOpponent(currentPlayer))
		}
	}

	node.score = 0
	if isOpponent(nestingLevel) {
		node.score = math.MaxInt
	}
	bestIdx := int8(-1)
	for range node.Children {
		kidNode := <-channelForKids
		if comparator(kidNode.score, node.score) {
			node.score = kidNode.score
			bestIdx = kidNode.int8
		}
	}
	parentChannel <- struct {
		*Node
		int8
	}{node, bestIdx}
}

func (node *Node) selectScoreAlphaBeta(board *Board, nestingLevel int8, currentPlayer utils.Player, alpha, beta int) int {
	defer func() {
		atomic.AddInt64(&_visitedNodesCounter, 1)
		if nestingLevel == 1 {
			atomic.AddInt64(&_visitedon1Level, 1)
		} else if nestingLevel == 2 {
			atomic.AddInt64(&_visitedon2Level, 1)
		} else if nestingLevel == 3 {
			atomic.AddInt64(&_visitedon3Level, 1)
		}
	}()
	if nestingLevel == utils.DEPTH {
		opponent := currentPlayer
		if isOpponent(nestingLevel) {
			opponent = getOpponent(currentPlayer)
		}
		node.score = _heuristic(board, opponent)
		return node.score
	}

	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}
	node.score = 0
	if isOpponent(nestingLevel) {
		node.score = math.MaxInt
	}

	moves, pawns := getMovesAndPawns(board, &currentPlayer)

	currentPawnId := 0
	for i := range moves {
		currentBoard := new(Board)
		*currentBoard = *board

		currentBoard.Pawns = [utils.PAWNS]*Pawn{}
		for i := int8(0); i < utils.PAWNS; i++ {
			currentBoard.Pawns[i] = new(Pawn)
			*currentBoard.Pawns[i] = *board.Pawns[i]
		}
		var currentPawn *Pawn
		for id := currentPawnId; id < len(currentBoard.Pawns); id++ {
			if pawns[i].Coords.X == currentBoard.Pawns[id].Coords.X && pawns[i].Coords.Y == currentBoard.Pawns[id].Coords.Y {
				currentPawn = currentBoard.Pawns[id]
				currentPawnId = id
				break
			}
		}

		initialPosition := &pawns[i].Coords
		currentBoard.MovePawn(currentPawn, moves[i])
		//if !done {
		//	continue
		//}
		newNode := NewNode(currentPawn, initialPosition)
		node.Children = append(node.Children, newNode)
		if validateResult(board) != "" {
			node.score = 100_000
			return node.score
		}
		newNodeScore := newNode.selectScoreAlphaBeta(currentBoard, nestingLevel+1, getOpponent(currentPlayer), alpha, beta)

		if comparator(newNodeScore, node.score) {
			node.score = newNodeScore
		}

		if isOpponent(nestingLevel) {
			alpha = max(alpha, newNodeScore)
			if beta <= alpha {
				atomic.AddInt64(&_prunedNodesCounter, 1)
				return node.score
			}
		} else {
			beta = min(beta, newNodeScore)
			if beta <= alpha {
				atomic.AddInt64(&_prunedNodesCounter, 1)
				return node.score
			}
		}
	}
	return node.score
}

func findMax(val int, max int) bool {
	if val > max {
		return true
	}
	return false
}

func findMin(val int, min int) bool {
	if val < min {
		return true
	}
	return false
}
