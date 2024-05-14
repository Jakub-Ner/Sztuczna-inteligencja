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
	//board.Print()
	score := 15_000 - distancesSum
	return score
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
	return -distancesSum + 1000
}

func NeighbourScore(board *Board, currentPlayer utils.Player) int {
	neighboursSum := 0
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		startingPawnIdx = 0
	}

	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		pawn := board.Pawns[i]
		for j := startingPawnIdx; j < startingPawnIdx+utils.PAWNS_PER_PLAYER; j++ {
			neighbour := board.Pawns[j]
			if pawn.Coords.DistanceManhattan(neighbour.Coords) <= 2 {
				neighboursSum++
			}
		}

	}
	return DistanceScoreManhattan(board, currentPlayer) - neighboursSum
}
