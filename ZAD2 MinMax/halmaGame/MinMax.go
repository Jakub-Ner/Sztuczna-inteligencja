package halmaGame

func MoveSelection(board *Board) (*Pawn, *Move) {

	tree := NewNode(nil, *board, 0, 0)
	channel := make(chan int)
	tree.selectScore(channel)

	bestScore := <-channel
	println("Best score: ", bestScore)

	for _, child := range tree.Children {
		if child.score == bestScore {
			return &child.pawn, &child.move
		}
	}

	return nil, nil
}
