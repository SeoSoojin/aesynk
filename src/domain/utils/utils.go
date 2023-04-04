package utils

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/maps"
)

func GenerateRandomNode(graph map[string]*node.Node) *node.Node {

	values := maps.Values(graph)

	if len(values) == 0 {
		return nil
	}

	index := rand.Intn(len(values))

	return values[index]

}

func StartMissingNodes(input map[string]*node.Node) map[string]struct{} {

	missingNodes := map[string]struct{}{}

	for key := range input {
		missingNodes[key] = struct{}{}
	}

	return missingNodes

}

func PrintSolution(path path.Path) {

	fmt.Println("Solution: ")
	if len(path.Nodes) == 0 {
		return
	}

	fmt.Printf("Path: ")
	for i := 0; i < len(path.Nodes)-1; i++ {
		fmt.Printf("%s -> ", path.Nodes[i].Name)
	}

	fmt.Printf("%s\n", path.Nodes[len(path.Nodes)-1].Name)

	fmt.Println("Cost: ", path.Cost)

}

func GenerateCostMatrix(graph map[string]*node.Node) [][]float64 {

	costMatrix := make([][]float64, len(graph))
	for i := range costMatrix {
		costMatrix[i] = make([]float64, len(graph))
	}

	for _, node := range graph {
		for _, edge := range node.Adjacents {
			name, err := strconv.Atoi(node.Name)
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(edge.To.Name)
			if err != nil {
				panic(err)
			}
			costMatrix[name-1][to-1] = edge.Weight
		}
	}

	return costMatrix

}
