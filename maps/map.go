package maps

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Represents the map of Samfundet
type Map struct {
	matrix [][]Cell
}

func New() *Map {
	return &Map{matrix: make([][]Cell, 0)}
}

// Returns the height of the map
func (m *Map) Height() int {
	return len(m.matrix)
}

// Returns the width of the map
func (m *Map) Width() int {
	return len(m.matrix[0])
}

// Returns the map cell for a given point
func (m *Map) Point(p Point) Cell {
	x, y := p[0], p[1]
	return m.matrix[y][x]
}

// Returns true if the cell can be navigated to
func (m *Map) Walkable(p Point) bool {
	return m.Point(p) != Void
}

// Returns the cardinal neighbors of a tile, if they can be moved to
func (m *Map) Neighbors(p Point) []Point {
	neighbors := []Point{}

	cardinals := []Point{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	for _, c := range cardinals {
		n := Point{p[0] + c[0], p[1] + c[1]}

		if m.Walkable(n) {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

// Implementation of the MapLoadable interface, to load inbound tokens
func (m *Map) Load(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ",")

		row := make([]Cell, len(tokens))

		for i, token := range tokens {
			val, err := strconv.Atoi(token)
			if err != nil {
				return err
			}

			row[i] = Cell(val)
		}

		m.matrix = append(m.matrix, row)
	}

	return nil
}

// Prints a given point on the map to the writer
func (m *Map) PrintPoint(w io.Writer, p Point) {
	v := m.Point(p)

	switch v {
	case Void:
		fmt.Fprint(w, "███")
	case FlatGround:
		fmt.Fprint(w, "   ")
	case Stairs:
		fmt.Fprint(w, "---")
	case CrowdedStairs:
		fmt.Fprint(w, "☰☰☰")
	case CrowdedRoom:
		fmt.Fprint(w, ":::")
	}
}

// Prints the entire map to the writer
func (m *Map) Print(w io.Writer) {
	for y, row := range m.matrix {
		for x := range row {
			m.PrintPoint(w, Point{x, y})
		}
		fmt.Fprint(w, "\n")
	}
}
