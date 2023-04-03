package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"github.com/seosoojin/aesynk/src/domain/state"
)

func main() {

	// err := graph.NewGraph(true, 2).GenerateCompleteGraph(50000, true).ToCSV("graph50k.csv")
	// if err != nil {
	// 	panic(err)
	// }

	node1 := node.Node{
		Name: "1",
	}

	auxState := state.State{
		InitialNode:  &node1,
		Current:      &node1,
		MissingNodes: map[string]struct{}{"2": {}},
		Path:         path.Path{Nodes: []*node.Node{&node1}, Cost: 1},
	}

	auxState.Write(os.Stdout, true)

	// fmt.Println("Graph generated.")

	// startTime := time.Now()

	// _, err := salesman.NewGeneticSolver(graph.Nodes(), 100, 200, 0.1, 0.1).Solve()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Time: ", time.Since(startTime))

	// startTime = time.Now()

	// _, err = salesman.NewPGeneticSolver(graph.Nodes(), 100, 200, 0.1, 0.1, 100).Solve()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Time: ", time.Since(startTime))

}

func printSolution(path path.Path) {

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

func generateCostMatrix(graph map[string]*node.Node) [][]float64 {

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

func printMatrix(matrix [][]float64) {

	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%f ", value)
		}
		fmt.Println()
	}

}
