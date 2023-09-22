package astar

import (
	"container/heap"

	"github.com/ssouthcity/astar-samf/maps"
)

// Creates an instance of the A* implementation that uses
// a middleware to update the goal to where the friends would
// be during the calculation
func NewMovingGoal(
	s maps.Point,
	gs []maps.Point,
	sd int,
	m *maps.Map,
) *Astar {
	a := &Astar{
		start:  s,
		goal:   gs[0],
		mapObj: m,

		open:   make(PriorityQueue, 0),
		closed: make(PriorityQueue, 0),

		middleware: func(a *Astar) {
			stepsTargetsTaken := a.currentNode.Len() / sd

			if stepsTargetsTaken >= len(gs) {
				stepsTargetsTaken = len(gs) - 1
			}

			a.goal = gs[stepsTargetsTaken]
		},

		currentNode: nil,
		goalReached: false,
	}

	heap.Push(&a.open, NewNode(a.start, 0, 0, nil))

	return a
}
