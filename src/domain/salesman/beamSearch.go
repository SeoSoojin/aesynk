package salesman

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
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

func nextStates(state State, width int, firstNode *node.Node) []*State {

	states := make([]*State, 0)

	for _, edge := range state.Current.Adjacents {

		_, ok := state.MissingNodes[edge.To.Name]
		if !ok || (edge.To.Name == firstNode.Name && len(state.MissingNodes) > 1) {
			continue
		}

		copyMissing := maps.Clone(state.MissingNodes)

		delete(copyMissing, edge.To.Name)

		nodes := slices.Clone(state.Path.Nodes)

		nodes = append(nodes, edge.To)

		costCopy := state.Path.Cost

		newState := State{
			Current:      edge.To,
			MissingNodes: copyMissing,
			Path:         path.Path{Nodes: nodes, Cost: costCopy + edge.Weight},
		}

		states = append(states, &newState)
	}

	return bestStates(states, width)

}

func bestStates(states []*State, width int) []*State {

	sort.SliceStable(states, func(i, j int) bool {
		return states[i].Path.Cost < states[j].Path.Cost
	})

	if len(states) <= width {
		return states
	}

	return states[:width]

}

func (b *BeamSearchSolver) Solve() (path.Path, error) {

	names := maps.Keys(b.graph)

	key := rand.Intn(len(names))

	firstNode := b.graph[names[key]]

	initialState := State{
		Current:      firstNode,
		MissingNodes: startMissingNodes(b.graph),
		Path:         path.Path{Nodes: []*node.Node{firstNode}},
	}

	beam := []*State{&initialState}

	for len(beam) > 0 {

		nextBeam := []*State{}

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

func startMissingNodes(input map[string]*node.Node) map[string]struct{} {

	missingNodes := map[string]struct{}{}

	for key := range input {
		missingNodes[key] = struct{}{}
	}

	return missingNodes

}

func validateSolution(input map[string]*node.Node, missing map[string]struct{}) bool {

	return len(missing) == 0

}

func PrintState(state State) {

	fmt.Printf("Current: %s\n", state.Current.Name)

	fmt.Printf("Missing: ")
	for key := range state.MissingNodes {
		fmt.Printf("%s ", key)
	}

	fmt.Printf("Path: ")
	for i := 0; i < len(state.Path.Nodes)-1; i++ {
		fmt.Printf("%s -> ", state.Path.Nodes[i].Name)
	}

	fmt.Printf("%s\n", state.Path.Nodes[len(state.Path.Nodes)-1].Name)

	fmt.Println("Cost: ", state.Path.Cost)

	fmt.Println()

}
