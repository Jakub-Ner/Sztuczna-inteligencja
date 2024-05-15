package halmaGame

import "ZAD2_MinMax/utils"

type EmptyField struct {
	Icon utils.Player
}

func (e EmptyField) GetIcon() string {
	return e.Icon
}

func (e EmptyField) CanMoveHere() bool {
	return true
}
