package halmaGame

import (
	"ZAD2_MinMax/utils"
	"fmt"
	"math"
)

type Node struct {
	Children []Node // to be calculated
	Parent   *Node  // passed to children
	score    int    // to be calculated

	currentPlayer utils.Player
	board         Board                       // passed to children
	pawn          *Pawn                       // to be calculated
	initialCords  *Coords                     // to be calculated
	nestingLevel  int8                        // passed to children
	comparator    func(val int, max int) bool // to be calculated
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

func NewNode(parent *Node, board Board, nestingLevel int8, score int, nextPlayer utils.Player, pawn *Pawn, initialCords *Coords) *Node {
	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}
	//par := parent
	bor := board
	return &Node{
		Children:      make([]Node, 0),
		Parent:        parent,
		score:         score,
		currentPlayer: nextPlayer,
		board:         bor,
		pawn:          pawn,
		initialCords:  initialCords,
		nestingLevel:  nestingLevel,
		comparator:    comparator,
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("Node{score: %d, nestingLevel: %d, pawnAddr: %p, parrentAddr: %p}", node.score, node.nestingLevel, node, node.Parent)
}

func (node *Node) selectScore(parentChannel chan *Node) {
	defer func() {
		fmt.Printf("Node: %s\n", node.String())
	}()
	if node.nestingLevel > utils.DEPTH {
		parentChannel <- node
		return
	}

	moves, pawns := getMovesAndPawns(&node.board, &node.currentPlayer)
	nextPlayer := getOpponent(node.currentPlayer)
	for i := range moves {
		currentPawn := *pawns[i]
		initialPosition := &pawns[i].Coords
		done := node.board.MovePawn(&currentPawn, moves[i])
		if !done {
			continue
		}
		newNode := NewNode(node, node.board, node.nestingLevel+1, int(moves[i].NumberOfJumps)+node.score, nextPlayer, &currentPawn, initialPosition) // use heuristic
		node.Children = append(node.Children, *newNode)
	}

	channelForKids := make(chan *Node, len(node.Children))
	if !isOpponent(node.nestingLevel) {
		for _, child := range node.Children {
			go child.selectScore(channelForKids)
		}
	} else {
		for _, child := range node.Children {
			child.selectScore(parentChannel)
		}
	}
	node.score = 0
	if isOpponent(node.nestingLevel) {
		node.score = math.MaxInt
	}
	for range node.Children {
		kidNode := <-channelForKids
		if node.comparator(kidNode.score, node.score) {
			node.score = kidNode.score
		}
	}
	//fmt.Printf("Best score for a parent %p, %s=%s", node, node.pawn.String(), node.move.String())

	parentChannel <- node
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
