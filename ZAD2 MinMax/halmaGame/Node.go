package halmaGame

import (
	"ZAD2_MinMax/utils"
	"fmt"
	"math"
)

type Node struct {
	Children []*Node // to be calculated
	Parent   *Node   // passed to children
	score    int     // to be calculated

	currentPlayer utils.Player
	board         *Board                      // passed to children
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

func NewNode(parent *Node, board *Board, nestingLevel int8, nextPlayer utils.Player, pawn *Pawn, initialCords *Coords) *Node {
	comparator := findMax
	if isOpponent(nestingLevel) {
		comparator = findMin
	}

	return &Node{
		Children:      make([]*Node, 0),
		Parent:        parent,
		score:         0,
		currentPlayer: nextPlayer,
		board:         board,
		pawn:          pawn,
		initialCords:  initialCords,
		nestingLevel:  nestingLevel,
		comparator:    comparator,
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("%p : score: %d, kids: %d", node.Parent, node.score, len(node.Children))
}

func (node *Node) selectScore(parentChannel chan struct {
	*Node
	int8
}, kidIdx int8) {
	//defer func() {
	//	tab := ""
	//	for i := int8(0); i < node.nestingLevel; i++ {
	//		tab += "\t"
	//	}
	//	fmt.Println(tab, node.String())
	//}()
	if node.nestingLevel == utils.DEPTH {
		node.score = DistanceScore(node.board, getOpponent(node.currentPlayer))
		parentChannel <- struct {
			*Node
			int8
		}{node, kidIdx}
		return
	}

	moves, pawns := getMovesAndPawns(node.board, &node.currentPlayer)
	nextPlayer := getOpponent(node.currentPlayer)

	for i := range moves {
		currentBoard := new(Board)
		*currentBoard = *node.board

		currentBoard.Pawns = [utils.PAWNS]*Pawn{}
		for i := int8(0); i < utils.PAWNS; i++ {
			currentBoard.Pawns[i] = new(Pawn)
			*currentBoard.Pawns[i] = *node.board.Pawns[i]
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
		newNode := NewNode(node, currentBoard, node.nestingLevel+1, nextPlayer, currentPawn, initialPosition)
		node.Children = append(node.Children, newNode)
	}
	//fmt.Println("Children: ", len(node.Children))
	channelForKids := make(chan struct {
		*Node
		int8
	}, len(node.Children))
	if !isOpponent(node.nestingLevel) { // !isOpponent(node.nestingLevel)
		for i, child := range node.Children {
			go child.selectScore(channelForKids, int8(i))
		}
	} else {
		for i, child := range node.Children {
			child.selectScore(channelForKids, int8(i))
		}
	}
	node.score = 0
	if isOpponent(node.nestingLevel) {
		node.score = math.MaxInt
	}
	bestIdx := int8(-1)
	for range node.Children {
		kidNode := <-channelForKids
		if node.comparator(kidNode.score, node.score) {
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
