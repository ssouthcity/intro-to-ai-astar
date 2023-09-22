package astar

import (
	"github.com/ssouthcity/astar-samf/maps"
)

// Simple golang priority queue implementation
type PriorityQueue []*Node

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].F < p[j].F
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*p = append(*p, node)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	node := old[n-1]
	*p = old[0 : n-1]
	return node
}

// Syntax sugar to check if queue is empty
func (p PriorityQueue) Empty() bool {
	return p.Len() == 0
}

// Returns a node from the queue, if the node's value is the given point
func (p PriorityQueue) Find(point maps.Point) *Node {
	for _, n := range p {
		if n.Point == point {
			return n
		}
	}
	return nil
}
