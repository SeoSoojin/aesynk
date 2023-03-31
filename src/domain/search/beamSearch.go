package search

import (
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type BeamSearchService struct {
	b int
}

func NewBeamSearcher(b int) Searcher {
	return &BeamSearchService{
		b: b,
	}
}

func (s *BeamSearchService) FindShortestWay(nodeMap map[string]*node.Node, start string, obj string) (path.Path, bool) {

	startNode, ok := nodeMap[start]
	if !ok {
		return path.Path{}, false
	}

	visited := make(map[string]struct{})

	return s.step(obj, startNode, 0, visited, path.Path{})

}

func (s *BeamSearchService) step(obj string, curr *node.Node, stepCost float64, visited map[string]struct{}, steps path.Path) (path.Path, bool) {

	if curr == nil {
		return steps, false
	}

	stepsCopy := slices.Clone(steps.Nodes)
	stepsCopy = append(stepsCopy, curr)
	steps.Cost += stepCost

	if curr.Name == obj {
		steps.Nodes = stepsCopy
		return steps, true
	}

	visitedCopy := maps.Clone(visited)

	if _, ok := visitedCopy[curr.Name]; ok {
		return steps, false
	}

	visitedCopy[curr.Name] = struct{}{}

	qnt := s.b

	adjacentsSample := make([]*node.Edge, 0)

	for i := 0; i < len(curr.Adjacents); i++ {

		if len(adjacentsSample) == qnt {
			break
		}

		if _, ok := visitedCopy[curr.Adjacents[i].To.Name]; ok {
			continue
		}

		adjacentsSample = append(adjacentsSample, curr.Adjacents[i])

	}

	flag := false
	betterCost := 0.0

	for _, adjacent := range adjacentsSample {

		p, found := s.step(obj, adjacent.To, adjacent.Weight, visitedCopy, path.Path{Nodes: stepsCopy, Cost: steps.Cost})
		if found && (p.Cost < betterCost || !flag) {

			betterCost = p.Cost
			steps = p
			flag = true

		}

	}

	if flag {
		return steps, true
	}

	return steps, false

}
