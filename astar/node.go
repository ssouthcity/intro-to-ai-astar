package astar

import "github.com/ssouthcity/astar-samf/maps"

// A node represents a point that the person has reached.
// It is represented as a linked list, so that the path the
// person took to get here can be traced back.
type Node struct {
	Point maps.Point

	F int
	G int
	H int

	Next *Node
}

func NewNode(point maps.Point, g int, h int, parent *Node) *Node {
	return &Node{
		Point: point,
		F:     g + h,
		G:     g,
		H:     h,
		Next:  parent,
	}
}

// Returns the path taken to get to the current node
func (n *Node) Path() []*Node {
	nodes := []*Node{}

	for n != nil {
		nodes = append(nodes, n)
		n = n.Next
	}

	return nodes
}

// Returns the amount of steps taken to reach the current node
func (n *Node) Len() int {
	return len(n.Path())
}
