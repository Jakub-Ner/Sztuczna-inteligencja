package utils

type Player = int8

const (
	PLAYER_GREEN  Player = 1
	PLAYER_YELLOW Player = 2
)

const (
	RESET_COLOR = "\033[0m"
	LIGHT_RED   = "\033[38;2;248;102;89m"
)

const (
	COLUMNS          = int8(16)
	PAWNS_PER_PLAYER = int8(19)
	PAWNS            = 2 * PAWNS_PER_PLAYER

	GREEN_PAWN  = "🟢"
	YELLOW_PAWN = "🟡"

	EMPTY_RED  = "🟥"
	EMPTY_BLUE = "🟦"
)
