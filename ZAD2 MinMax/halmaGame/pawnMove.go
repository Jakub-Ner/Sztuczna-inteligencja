package halmaGame

import "fmt"

type Coords struct {
	X int8
	Y int8
}

func (c *Coords) String() string {
	return fmt.Sprintf("(x=%d, y=%d)", c.X, c.Y)
}

func (c *Coords) Distance(other Coords) int8 {
	return int8((c.X-other.X)*(c.X-other.X) + (c.Y-other.Y)*(c.Y-other.Y))
}

type Move struct {
	NumberOfJumps int8
	Direction     Coords
}

func (m Move) String() string {
	return fmt.Sprintf("{jumps=%d, (x=%d, y=%d)}", m.NumberOfJumps, m.Direction.X, m.Direction.Y)
}
