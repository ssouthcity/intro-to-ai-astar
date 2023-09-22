package astar

import (
	"container/heap"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/ssouthcity/astar-samf/maps"
)

// Struct holding the state of the A* algorithm
// middleware is used to modify state externally during runtime
type Astar struct {
	start  maps.Point
	goal   maps.Point
	mapObj *maps.Map

	open   PriorityQueue
	closed PriorityQueue

	middleware func(*Astar)

	currentNode *Node
	goalReached bool
}

func New(s maps.Point, g maps.Point, m *maps.Map) *Astar {
	a := &Astar{
		start:  s,
		goal:   g,
		mapObj: m,

		open:   make(PriorityQueue, 0),
		closed: make(PriorityQueue, 0),

		middleware: func(a *Astar) {},

		currentNode: nil,
		goalReached: false,
	}

	heap.Push(&a.open, NewNode(a.start, 0, 0, nil))

	return a
}

// Returns true when the algorithm is completed
func (a *Astar) Done() bool {
	return a.goalReached || a.open.Empty()
}

// Single iteration of the algorithm
func (a *Astar) Step() {
	q := heap.Pop(&a.open).(*Node)
	a.currentNode = q

	successors := a.mapObj.Neighbors(q.Point)

	for _, successor := range successors {
		if successor == a.goal {
			a.currentNode = NewNode(successor, 0, 0, q)
			a.goalReached = true
			return
		}

		// Middleware allows an external entity to modify the algorithm during runtime
		// Used to update the goal position when the target moves
		a.middleware(a)

		g := q.G + int(a.mapObj.Point(successor))
		h := successor.Distance(a.goal)

		f := g + h

		// if node is already found with a better route
		if n := a.open.Find(successor); n != nil && n.F < f {
			continue
		}

		// if node has already been visited with a better route
		if n := a.closed.Find(successor); n != nil && n.F < f {
			continue
		}

		heap.Push(&a.open, NewNode(successor, g, h, q))
	}

	heap.Push(&a.closed, q)
}

// Run the algorithm with visualization printed to the writer
func (a *Astar) Visualize(w io.Writer) error {
	for !a.Done() {
		a.Step()

		fmt.Fprint(w, "\033[H")
		a.Print(w)

		time.Sleep(5 * time.Millisecond)
	}

	if !a.goalReached {
		return errors.New("unable to reach goal")
	}

	return nil
}

// Complete the algorithm for solution without visualization
func (a *Astar) Solve() ([]maps.Point, error) {
	for !a.Done() {
		a.Step()
	}

	if !a.goalReached {
		return nil, errors.New("unable to reach goal")
	}

	victoryPath := a.currentNode.Path()

	pathPoints := make([]maps.Point, len(victoryPath))
	for i, n := range victoryPath {
		pathPoints[i] = n.Point
	}

	return pathPoints, nil
}

// Prints single tile to the writer, uses chain of responsibility pattern,
// i.e. if the astar algorithm wants to print the tile, it does. Otherwise
// it passes the call to the map
func (a *Astar) PrintPoint(w io.Writer, point maps.Point) {
	if point == a.start {
		fmt.Fprint(w, " S ")
		return
	}

	if point == a.goal {
		fmt.Fprint(w, " G ")
		return
	}

	for _, n := range a.currentNode.Path() {
		if point == n.Point {
			fmt.Fprint(w, " # ")
			return
		}
	}

	a.mapObj.PrintPoint(w, point)
}

// Prints every tile in the map to the writer
func (a *Astar) Print(w io.Writer) {
	for y := 0; y < a.mapObj.Height(); y++ {
		for x := 0; x < a.mapObj.Width(); x++ {
			p := maps.Point{x, y}
			a.PrintPoint(w, p)
		}
		fmt.Fprint(w, "\n")
	}
}
