package maps

import (
	"fmt"
)

// A point represents the grid location of a tile on the map
type Point [2]int

// Calculates the manhattan distance between two points
func (p Point) Distance(o Point) int {
	a := p[0] - o[0]
	if a < 0 {
		a = a * -1
	}

	b := p[1] - o[1]
	if b < 0 {
		b = b * -1
	}

	return a + b
}

// This method is called when converting a point to a string (debugging)
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p[0], p[1])
}
