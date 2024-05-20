package halmaGame

import (
	"ZAD2_MinMax/utils"
)

type Heuristic func(board *Board, currentPlayer utils.Player) int

func DistanceScore(board *Board, currentPlayer utils.Player) int {
	distancesSum := 0
	target := Coords{15, 0}
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		target = Coords{0, 15}
		startingPawnIdx = 0
	}

	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		newDist := board.Pawns[i].Coords.DistanceSquare(target)
		distancesSum += newDist
	}
	return 10_000 - distancesSum + finishScore(board, currentPlayer)
}

func DistanceManhattanScore(board *Board, currentPlayer utils.Player) int {
	distancesSum := 0
	target := Coords{15, 0}
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		target = Coords{0, 15}
		startingPawnIdx = 0
	}

	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		distancesSum += board.Pawns[i].Coords.DistanceManhattan(target)
	}
	return 10_000 - distancesSum*10 + finishScore(board, currentPlayer)/10
}

func ManhattanAdaptiveMoveNumScore(board *Board, currentPlayer utils.Player) int {
	distScore := 0
	if _turnCounter > 5 {
		distScore += DistanceManhattanScore(board, currentPlayer)
	}
	moveScore := 0
	if _turnCounter < 30 {
		moveScore = moveNumScoreHelper(board, currentPlayer)
	}
	return int(float64(moveScore)*1.5) + distScore
}

func MoveNumScore(board *Board, currentPlayer utils.Player) int {
	moveScore := moveNumScoreHelper(board, currentPlayer)
	distScore := DistanceScore(board, currentPlayer)
	return moveScore + distScore
}

func MoveNumManhattanScore(board *Board, currentPlayer utils.Player) int {
	moveScore := moveNumScoreHelper(board, currentPlayer)
	distScore := DistanceManhattanScore(board, currentPlayer)
	return int(float64(moveScore)*1.5) + distScore
}

func ManhattanAdaptiveNeighbourScore(board *Board, currentPlayer utils.Player) int {
	distScore := DistanceManhattanScore(board, currentPlayer)
	neighScore := 0
	if _turnCounter > _activationMoment1 && _turnCounter < _activationMoment2 {
		neighScore = neighbourScoreHelper(board, currentPlayer)
	}
	return neighScore + distScore
}

func NeighbourScore(board *Board, currentPlayer utils.Player) int {
	neighScore := neighbourScoreHelper(board, currentPlayer)
	distScore := DistanceScore(board, currentPlayer)
	return neighScore + distScore
}
func NeighbourScore2(board *Board, currentPlayer utils.Player) int {
	neighScore := neighbourScoreHelper(board, currentPlayer)
	distScore := DistanceScore(board, currentPlayer)
	return int(float64(neighScore)*_dominance2) + distScore
}

func NeighbourManhattanScore(board *Board, currentPlayer utils.Player) int {
	neighScore := neighbourScoreHelper(board, currentPlayer)
	distScore := DistanceManhattanScore(board, currentPlayer)
	return neighScore + distScore
}

func NeighbourManhattanScore2(board *Board, currentPlayer utils.Player) int {
	neighScore := neighbourScoreHelper(board, currentPlayer)
	distScore := DistanceManhattanScore(board, currentPlayer)
	return int(float64(neighScore)*_dominance2) + distScore
}

func moveNumScoreHelper(board *Board, currentPlayer utils.Player) int {
	score := 0
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	}
	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		score += len(board.Pawns[i].ValidMoves)
	}
	return score
}

func neighbourScoreHelper(board *Board, currentPlayer utils.Player) int {
	neighboursSum := 0
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	}

	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		pawn := board.Pawns[i]
		for j := int8(0); j < utils.PAWNS; j++ {
			neighbour := board.Pawns[j]
			if pawn.Coords.DistanceManhattan(neighbour.Coords) <= 2 {
				neighboursSum++
			}
		}

	}
	return neighboursSum
}

func finishScore(board *Board, currentPlayer utils.Player) int {
	const BONUS = 500
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	}
	finishedCounter := 0
	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		if DoesPawnFinished(board.Pawns[i], currentPlayer) {
			finishedCounter++
		}
	}
	if finishedCounter == int(utils.PAWNS_PER_PLAYER) {
		return BONUS * 100
	}
	return finishedCounter * BONUS
}
