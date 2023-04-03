package node

import (
	"math"
	"sort"
)

const (
	DIMENSION = 2
)

type Edge struct {
	From   *Node
	To     *Node
	Weight float64
}

type Node struct {
	Name         string
	Coordinates  []float64
	Adjacents    []*Edge
	AdjacentsMap map[string]*Edge
}

func NewNode(name string, coordinates []float64) *Node {

	if len(coordinates) != DIMENSION {
		panic("Invalid number of coordinates")
	}

	return &Node{
		Name:         name,
		Coordinates:  coordinates,
		Adjacents:    make([]*Edge, 0),
		AdjacentsMap: make(map[string]*Edge),
	}

}

func (n *Node) AddAdjacent(dest *Node) *Node {

	if n == dest {
		return n
	}

	if n == nil || dest == nil {
		return n
	}

	sum := 0.0

	_, ok := n.AdjacentsMap[dest.Name]
	if ok {
		return n
	}

	for index, coordinate := range n.Coordinates {
		sum += (coordinate - dest.Coordinates[index]) * (coordinate - dest.Coordinates[index])
	}

	edge := &Edge{
		From:   n,
		To:     dest,
		Weight: math.Sqrt(sum),
	}

	tempAdjacents := append(n.Adjacents, edge)

	sort.SliceStable(tempAdjacents, func(i, j int) bool {
		return tempAdjacents[i].Weight < tempAdjacents[j].Weight
	})

	n.Adjacents = tempAdjacents

	n.AdjacentsMap[dest.Name] = edge

	return n

}

func (n *Node) AddAdjacents(dests ...*Node) *Node {

	for _, dest := range dests {
		n.AddAdjacent(dest)
	}

	return n

}
