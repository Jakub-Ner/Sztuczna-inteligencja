package utils

type Player = string

const (
	DEPTH = 4
)

const (
	PLAYER_GREEN  Player = GREEN_PAWN
	PLAYER_YELLOW Player = YELLOW_PAWN
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

const (
	DISPLAY_MARGIN = 5

	BOARD_HEIGHT = 19
	BOARD_WIDTH  = 55
)
