package salesman

import (
	"fmt"
	"sort"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"github.com/seosoojin/aesynk/src/domain/state"
	"github.com/seosoojin/aesynk/src/domain/utils"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type BeamSearchSolver struct {
	graph map[string]*node.Node
	width int
}

func NewBeamSearchSolver(graph map[string]*node.Node, width int) Solver {
	return &BeamSearchSolver{
		graph: graph,
		width: width,
	}
}

func nextStates(input state.State, width int, firstNode *node.Node) []*state.State {

	states := make([]*state.State, 0)

	for _, edge := range input.Current.Adjacents {

		_, ok := input.MissingNodes[edge.To.Name]
		if !ok || (edge.To.Name == firstNode.Name && len(input.MissingNodes) > 1) {
			continue
		}

		copyMissing := maps.Clone(input.MissingNodes)

		delete(copyMissing, edge.To.Name)

		nodes := slices.Clone(input.Path.Nodes)

		nodes = append(nodes, edge.To)

		costCopy := input.Path.Cost

		newState := state.State{
			Current:      edge.To,
			MissingNodes: copyMissing,
			Path:         path.Path{Nodes: nodes, Cost: costCopy + edge.Weight},
		}

		states = append(states, &newState)
	}

	return bestStates(states, width)

}

func bestStates(states []*state.State, width int) []*state.State {

	sort.SliceStable(states, func(i, j int) bool {
		return states[i].Path.Cost < states[j].Path.Cost
	})

	if len(states) <= width {
		return states
	}

	return states[:width]

}

func (b *BeamSearchSolver) Solve() (path.Path, error) {

	firstNode := utils.GenerateRandomNode(b.graph)
	initialState := state.State{
		Current:      firstNode,
		MissingNodes: utils.StartMissingNodes(b.graph),
		Path:         path.Path{Nodes: []*node.Node{firstNode}},
	}

	beam := []*state.State{&initialState}

	for len(beam) > 0 {

		nextBeam := []*state.State{}

		for _, state := range beam {

			for _, nextState := range nextStates(*state, b.width, firstNode) {

				if validateSolution(b.graph, nextState.MissingNodes) {
					return nextState.Path, nil
				}

				nextBeam = append(nextBeam, nextState)

			}

		}

		if len(nextBeam) > b.width {
			nextBeam = bestStates(nextBeam, b.width)
		}

		beam = nextBeam

	}

	return path.Path{}, fmt.Errorf("no solution found")

}

func validateSolution(input map[string]*node.Node, missing map[string]struct{}) bool {

	return len(missing) == 0

}
