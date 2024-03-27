package halmaGame

import "fmt"

type Coords struct {
	X int8
	Y int8
}

func (c Coords) String() string {
	return fmt.Sprintf("(x=%d, y=%d)", c.X, c.Y)
}

type Move struct {
	NumberOfJumps int8
	Direction     Coords
}

func (m Move) String() string {
	return fmt.Sprintf("{jumps=%d, (x=%d, y=%d)}", m.NumberOfJumps, m.Direction.X, m.Direction.Y)
}
