package halmaGame

import (
	"ZAD2_MinMax/utils"
	"strconv"
)

type Game struct {
	board         *Board
	currentPlayer utils.Player
}

func DoesPawnFinished(pawn *Pawn, player utils.Player) bool {
	x, y := pawn.Coords.X, pawn.Coords.Y
	if WIN_BOARD[y][x] == player {
		return true
	}
	return false
}

func validateResult(board *Board) string {
	var i int8
	player := utils.PLAYER_YELLOW
	for i = utils.PAWNS_PER_PLAYER; i < utils.PAWNS; i++ {
		if DoesPawnFinished(board.Pawns[i], player) == false {
			break
		}
	}
	if i == utils.PAWNS {
		return player
	}

	player = utils.GREEN_PAWN
	for i = 0; i < utils.PAWNS_PER_PLAYER; i++ {
		if DoesPawnFinished(board.Pawns[i], player) == false {
			break
		}
	}
	if i == utils.PAWNS_PER_PLAYER {
		return player
	}
	return ""
}

func getOpponent(player utils.Player) utils.Player {
	if player == utils.PLAYER_YELLOW {
		return utils.PLAYER_GREEN
	}
	return utils.PLAYER_YELLOW
}

type SelectionMove func() (*Pawn, *Move)

func (g *Game) runGame(yellowSelectionMove SelectionMove, greenSelectionMove SelectionMove) {
	initGlobals()
	g.board = NewBoard()
	g.currentPlayer = utils.PLAYER_YELLOW
	var selectedPawn *Pawn
	var selectedMove *Move

	var win string
	for win = ""; win == ""; win = validateResult(g.board) {
		if _turnCounter/2 > 200 {
			break
		}
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
	g.board.Print()
	if _turnCounter/2 >= 200 {
		utils.PrintMessage("Game ended in a draw")
		return
	}
	if win == utils.YELLOW_PAWN {
		utils.PrintMessage("Yellow player won")
	} else {
		utils.PrintMessage("Green player won")
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
	pc := func() (*Pawn, *Move) {
		return g.minMaxMoveSelection(DistanceScore)
	}
	g.runGame(g.letPlayerMove, pc)
}

func (g *Game) RunGameComputerVSComputer() {
	heuristics := []Heuristic{
		DistanceScore,
		DistanceManhattanScore,
		MoveNumScore,
		MoveNumManhattanScore,
		NeighbourScore,
		NeighbourManhattanScore,
		ManhattanAdaptiveMoveNumScore,
		ManhattanAdaptiveNeighbourScore,
	}

	for _, heuristic1 := range heuristics {
		for _, heuristic2 := range heuristics {
			pc1 := func() (*Pawn, *Move) {
				return g.minMaxMoveSelection(heuristic1)
			}
			pc2 := func() (*Pawn, *Move) {
				return g.minMaxMoveSelection(heuristic2)
			}
			utils.CountAndShowTime(g.runGame, pc1, pc2)

			utils.PrintMessage("------------------------------------------")
			utils.PrintMessage("Yellow uses: " + utils.GetFunctionName(heuristic1))
			utils.PrintMessage("Green uses: " + utils.GetFunctionName(heuristic2))
			utils.PrintMessage(
				"" +
					"Visited nodes: " + strconv.FormatInt(_allVisitedNodesCounter, 10) +
					"; Pruned nodes: " + strconv.FormatInt(_allPrunedNodesCounter, 10) +
					"; Rounds: " + strconv.Itoa(_turnCounter/2))
		}
	}

}

func (g *Game) minMaxMoveSelection(heuristic Heuristic) (*Pawn, *Move) {
	return MoveSelection(g.board, g.currentPlayer, heuristic)
}
