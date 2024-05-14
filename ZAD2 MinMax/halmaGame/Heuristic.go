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
	return 10_000 - distancesSum + FinishScore(board, currentPlayer)
}

func DistanceScoreManhattan(board *Board, currentPlayer utils.Player) int {
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
	return 10_000 - distancesSum*10 + FinishScore(board, currentPlayer)/10
}

func FinishScore(board *Board, currentPlayer utils.Player) int {
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

func MoveNumScore(board *Board, currentPlayer utils.Player) int {
	score := 0
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	}
	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		score += len(board.Pawns[i].ValidMoves)
	}
	return score + DistanceScoreManhattan(board, currentPlayer)
}

func NeighbourScore(board *Board, currentPlayer utils.Player) int {
	const BONUS = 5
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
	return DistanceScore(board, currentPlayer) + neighboursSum*BONUS
}
