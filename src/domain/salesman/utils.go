package salesman

import (
	"math/rand"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/maps"
)

func generateRandomNode(graph map[string]*node.Node) *node.Node {

	values := maps.Values(graph)

	index := rand.Intn(len(values))

	return values[index]

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

func validateSolutionGenetic(input map[string]*node.Node, output path.Path) bool {

	if len(output.Nodes) != len(input)+1 {
		return false
	}

	nodesMap := map[string]struct{}{}

	for _, node := range output.Nodes {
		nodesMap[node.Name] = struct{}{}
	}

	return len(nodesMap) == len(input) && output.Nodes[0].Name == output.Nodes[len(output.Nodes)-1].Name

}
