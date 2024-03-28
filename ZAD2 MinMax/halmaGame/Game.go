package halmaGame

import "ZAD2_MinMax/utils"

type Game struct {
	board         *Board
	currentPlayer utils.Player
}

// TODO: implement the function
func validateResult() bool {
	return false
}

func getOpponent(player utils.Player) utils.Player {
	if player == utils.PLAYER_YELLOW {
		return utils.PLAYER_GREEN
	}
	return utils.PLAYER_YELLOW
}

func (g *Game) RunGame() {
	g.board = NewBoard()
	g.currentPlayer = utils.PLAYER_YELLOW

	for !validateResult() {
		g.board.Print()
		selectedPawn := g.board.selectPawn(g.currentPlayer)
		displayMoves(*selectedPawn)
		selectedMove := g.board.selectMove(selectedPawn)
		if selectedMove == nil {
			continue
		}
		g.board.MovePawn(selectedPawn, *selectedMove)
		g.board.updateMoves()

		g.currentPlayer = getOpponent(g.currentPlayer)
	}
}
