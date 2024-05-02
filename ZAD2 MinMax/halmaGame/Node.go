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

	board        Board                       // passed to children
	pawn         Pawn                        // to be calculated
	move         Move                        // to be calculated
	nestingLevel int                         // passed to children
	comparator   func(val int, max int) bool // to be calculated
}

func getMovesAndPawns(board *Board) ([]Move, []Pawn) {
	// TODO: return only moves for the current player
	//var startingPawnIdx int8
	//if board.CurrentPlayer == utils.PLAYER_GREEN {
	//	startingPawnIdx = 0
	//} else {
	//	startingPawnIdx = utils.PAWNS_PER_PLAYER
	//}
	fmt.Println("Try:  update moves:", utils.GetGoroutine())
	err := board.UpdateMoves()
	if err != nil {
		fmt.Println("Error: update moves:", err)
	}
	fmt.Println("Done: update moves:", utils.GetGoroutine())
	moves := make([]Move, 0, 80)
	pawns := make([]Pawn, 0, 80)
	for _, pawn := range board.Pawns {
		for _, move := range pawn.ValidMoves {
			moves = append(moves, move)
			pawns = append(pawns, *pawn)
		}
	}
	return moves, pawns
}

func isOpponent(nestingLevel int) bool {
	return nestingLevel%2 == 1
}

func NewNode(parent *Node, board Board, nestingLevel int, score int) *Node {
	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}
	par := parent
	bor := board
	return &Node{
		Children:     make([]Node, 0),
		Parent:       par,
		score:        score,
		board:        bor,
		pawn:         Pawn{},
		move:         Move{},
		nestingLevel: nestingLevel,
		comparator:   comparator,
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("Node{score: %d, nestingLevel: %d, pawnNum: %d, moveNum: %d, boardAddr: %p}", node.score, node.nestingLevel, len(node.Children), len(node.Children), &node.board)
}

func (node *Node) selectScore(parentChannel chan int) {
	if node.nestingLevel >= utils.DEPTH {
		parentChannel <- node.score
		return
	}

	//fmt.Println("getting moves and pawns", node.String(), utils.GetGoroutine())
	moves, pawns := getMovesAndPawns(&node.board)
	for i := range moves {
		newBoard := node.board
		currentPawn := pawns[i]
		//fmt.Println("trying to move pawn", currentPawn.String(), "to", moves[i].String(), utils.GetGoroutine())
		newBoard.MovePawn(&currentPawn, moves[i])
		//fmt.Println("moved pawn", currentPawn.String(), "to", moves[i].String(), utils.GetGoroutine())

		newNode := NewNode(node, newBoard, node.nestingLevel+1, int(moves[i].NumberOfJumps)+node.score) // use heuristic
		node.Children = append(node.Children, *newNode)
	}

	channelForKids := make(chan int, len(node.Children))
	for _, child := range node.Children {
		go child.selectScore(channelForKids)
	}

	node.score = 0
	if isOpponent(node.nestingLevel) {
		node.score = math.MaxInt
	}
	for range node.Children {
		val := <-channelForKids
		if node.comparator(val, node.score) {
			node.score = val
			//node.pawn
			//node.move
		}
	}
	fmt.Println("sending score to parent", node.String())
	parentChannel <- node.score
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
