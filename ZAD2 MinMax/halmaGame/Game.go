package halmaGame

import "ZAD2_MinMax/utils"

type Game struct {
	board         *Board
	currentPlayer utils.Player
}

func validateResult() bool {
	return false
}

func (g *Game) RunGame() {
	g.board = NewBoard()
	g.currentPlayer = utils.PLAYER_YELLOW

	for !validateResult() {
		g.board.Print()
		// color validation
		selectedPawn := g.board.selectPawn(g.currentPlayer)
		displayMoves(*selectedPawn)
		selectedMove := g.board.selectMove(selectedPawn)
		if selectedMove == nil {
			continue
		}
		g.board.MovePawn(selectedPawn, *selectedMove)
		//	updating the moves
	}
}
