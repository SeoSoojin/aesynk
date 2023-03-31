package salesman

import (
	"fmt"
	"sort"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type nodeSP struct {
	node            *node.Node
	averageDistance float64
}

type BeamSSolutionService struct {
	graph    map[string]*nodeSP
	b        int
	solved   bool
	solution path.Path
}

func NewBeamSSolutionService(b int, graph map[string]*node.Node) Solver {
	return &BeamSSolutionService{
		b:        b,
		graph:    findAverageDistances(graph),
		solved:   false,
		solution: path.Path{},
	}
}

func (s *BeamSSolutionService) Solve() (path.Path, error) {

	if s.solved {
		return s.solution, nil
	}

	nBestStartNodes := s.findBestStartNodes(s.graph)
	solution := path.Path{}
	betterCost := 0.0
	flag := false

	for _, startNode := range nBestStartNodes {
		visited := make(map[string]struct{})
		p, found := s.step(startNode, s.graph, visited, path.Path{}, 0.0)
		if found && (p.Cost > betterCost || !flag) {
			betterCost = p.Cost
			solution = p
			flag = true
		}
	}

	if flag {
		s.solved = true
		s.solution = solution
		return solution, nil
	}

	return path.Path{}, fmt.Errorf("no solution found")

}

func (s *BeamSSolutionService) step(curr *nodeSP, input map[string]*nodeSP, visited map[string]struct{}, steps path.Path, stepCost float64) (path.Path, bool) {

	if curr == nil {
		return steps, false
	}

	stepsCopy := slices.Clone(steps.Nodes)
	stepsCopy = append(stepsCopy, curr.node)
	steps.Cost += stepCost

	visitedCopy := maps.Clone(visited)

	if _, ok := visitedCopy[curr.node.Name]; ok {
		return steps, false
	}

	visitedCopy[curr.node.Name] = struct{}{}

	if validateSolution(input, visitedCopy) {
		steps.Nodes = stepsCopy
		return steps, true
	}

	qnt := s.b

	adjacentsSample := make([]*node.Edge, 0)

	for i := 0; i < len(curr.node.Adjacents); i++ {

		if len(adjacentsSample) == qnt {
			break
		}

		if _, ok := visitedCopy[curr.node.Adjacents[i].To.Name]; ok {
			continue
		}

		adjacentsSample = append(adjacentsSample, curr.node.Adjacents[i])

	}

	flag := false
	betterCost := 0.0

	for _, adjacent := range adjacentsSample {

		p, found := s.step(input[adjacent.To.Name], input, visitedCopy, path.Path{
			Nodes: stepsCopy,
			Cost:  steps.Cost,
		}, adjacent.Weight)

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

func (s *BeamSSolutionService) findBestStartNodes(input map[string]*nodeSP) []*nodeSP {

	bestNodes := make([]*nodeSP, 0)

	for _, n := range input {
		bestNodes = append(bestNodes, n)
	}

	sort.SliceStable(bestNodes, func(i, j int) bool {
		return bestNodes[i].averageDistance < bestNodes[j].averageDistance
	})

	if len(bestNodes) < s.b {
		return bestNodes
	}

	return bestNodes[:s.b]

}

func findAverageDistances(graph map[string]*node.Node) map[string]*nodeSP {

	average := 0.0

	graphWithDistance := make(map[string]*nodeSP)

	for key, n := range graph {

		average = calcAverageDistance(n)
		graphWithDistance[key] = &nodeSP{
			node:            n,
			averageDistance: average,
		}

	}

	return graphWithDistance

}

func calcAverageDistance(n *node.Node) float64 {

	sum := 0.0

	for _, edge := range n.Adjacents {
		sum += edge.Weight
	}

	return sum / float64(len(n.Adjacents))

}

func validateSolution(input map[string]*nodeSP, visited map[string]struct{}) bool {

	if len(input) != len(visited) {
		return false
	}

	for key := range input {
		if _, ok := visited[key]; !ok {
			return false
		}
	}

	return true

}
