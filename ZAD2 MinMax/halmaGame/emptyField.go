package halmaGame

type EmptyField struct {
	Icon string
}

func (e EmptyField) GetIcon() string {
	return e.Icon
}

func (e EmptyField) CanMoveHere() bool {
	return true
}
