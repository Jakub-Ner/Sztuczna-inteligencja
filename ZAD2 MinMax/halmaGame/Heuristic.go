package halmaGame

import (
	"ZAD2_MinMax/utils"
)

func DistanceScore(board *Board, currentPlayer utils.Player) int {
	//fmt.Print("calculating for: ", currentPlayer)
	distancesSum := 0
	target := Coords{15, 0}
	startingPawnIdx := utils.PAWNS_PER_PLAYER
	if currentPlayer == utils.PLAYER_GREEN {
		target = Coords{0, 15}
		startingPawnIdx = 0
	}

	for i := startingPawnIdx; i < startingPawnIdx+utils.PAWNS_PER_PLAYER; i++ {
		distancesSum += int(board.Pawns[i].Coords.Distance(target))
	}
	//board.Print()
	return -distancesSum + 1000
}
