package halmaGame

import (
	"ZAD2_MinMax/utils"
)

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

type SelectionMove func() (*Pawn, *Move)

func (g *Game) runGame(yellowSelectionMove SelectionMove, greenSelectionMove SelectionMove) {
	g.board = NewBoard()
	g.currentPlayer = utils.PLAYER_YELLOW
	var selectedPawn *Pawn
	var selectedMove *Move

	for !validateResult() {
		g.board.Print()
		if g.currentPlayer == utils.PLAYER_YELLOW {
			selectedPawn, selectedMove = yellowSelectionMove()
		} else {
			selectedPawn, selectedMove = greenSelectionMove()
		}

		g.board.MovePawn(selectedPawn, *selectedMove)
		g.board.UpdateMoves()

		g.currentPlayer = getOpponent(g.currentPlayer)
	}

}

func (g *Game) letPlayerMove() (*Pawn, *Move) {
	var selectedMove *Move
	var selectedPawn *Pawn
	for selectedMove == nil {
		selectedPawn = g.board.selectPawn(g.currentPlayer)
		displayMoves(*selectedPawn)
		selectedMove = g.board.selectMove(selectedPawn)
	}
	return selectedPawn, selectedMove
}
func (g *Game) RunGamePlayerVSPlayer() {
	g.runGame(g.letPlayerMove, g.letPlayerMove)
}

func (g *Game) RunGamePlayerVSComputer() {
	g.runGame(g.letPlayerMove, g.minMaxMoveSelection)
}

func (g *Game) RunGameComputerVSComputer() {
	g.runGame(g.minMaxMoveSelection, g.minMaxMoveSelection)
}

func (g *Game) minMaxMoveSelection() (*Pawn, *Move) {
	return MoveSelection(g.board)
}
