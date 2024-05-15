package halmaGame

import "fmt"

type Coords struct {
	X int8
	Y int8
}

func (c *Coords) String() string {
	return fmt.Sprintf("(x=%d, y=%d)", c.X, c.Y)
}

func (c *Coords) DistanceSquare(other Coords) int {
	xDist := int(c.X - other.X)
	yDist := int(c.Y - other.Y)
	return (xDist * xDist) + (yDist * yDist)
}

func (c *Coords) DistanceManhattan(other Coords) int {
	return int(abs(c.X-other.X) + abs(c.Y-other.Y))
}

func abs(number int8) int8 {
	if number < 0 {
		return -number
	}
	return number
}

type Move struct {
	NumberOfJumps int8
	Direction     Coords
}

func (m Move) String() string {
	return fmt.Sprintf("{jumps=%d, (x=%d, y=%d)}", m.NumberOfJumps, m.Direction.X, m.Direction.Y)
}
