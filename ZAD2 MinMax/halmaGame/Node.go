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

var heuristic Heuristic
var visitedNodesCounter int64
var prunedNodesCounter int64

var visitedon1Level int64
var visitedon2Level int64
var visitedon3Level int64

var alpha int64
var beta int64

func (node *Node) selectScore(parentChannel chan struct {
	*Node
	int8
}, kidIdx int8, board *Board, nestingLevel int8, currentPlayer utils.Player) {
	//defer func() {
	//	tab := ""
	//	for i := int8(0); i < nestingLevel; i++ {
	//		tab += "\t"
	//	}
	//	fmt.Println(tab, node.String())
	//}()
	defer func() {
		atomic.AddInt64(&visitedNodesCounter, 1)
		if nestingLevel == 1 {
			atomic.AddInt64(&visitedon1Level, 1)
		} else if nestingLevel == 2 {
			atomic.AddInt64(&visitedon2Level, 1)
		} else if nestingLevel == 3 {
			atomic.AddInt64(&visitedon3Level, 1)
		}
	}()
	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}

	if nestingLevel == utils.DEPTH {
		node.score = heuristic(board, getOpponent(currentPlayer))
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

	for i := range moves {
		currentBoard := new(Board)
		*currentBoard = *board

		currentBoard.Pawns = [utils.PAWNS]*Pawn{}
		for i := int8(0); i < utils.PAWNS; i++ {
			currentBoard.Pawns[i] = new(Pawn)
			*currentBoard.Pawns[i] = *board.Pawns[i]
		}
		var currentPawn *Pawn
		for _, pawn := range currentBoard.Pawns {
			if pawns[i].Coords.X == pawn.Coords.X && pawns[i].Coords.Y == pawn.Coords.Y {
				currentPawn = pawn
				break
			}
		}
		initialPosition := &pawns[i].Coords
		done := currentBoard.MovePawn(currentPawn, moves[i])
		if !done {
			continue
		}
		newNode := NewNode(currentPawn, initialPosition)
		node.Children = append(node.Children, newNode)

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
