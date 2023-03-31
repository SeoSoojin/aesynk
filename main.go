package main

import (
	"fmt"

	"github.com/seosoojin/aesynk/src/domain/graph"
	"github.com/seosoojin/aesynk/src/domain/path"
	"github.com/seosoojin/aesynk/src/domain/salesman"
)

func main() {

	graph, err := graph.NewGraph(true).FromCSV("graph.csv")
	if err != nil {
		panic(err)
	}

	path, err := salesman.NewBeamSSolutionService(7, graph.Nodes()).Solve()
	if err != nil {
		panic(err)
	}

	printSolution(path)

}

func printSolution(path path.Path) {

	fmt.Printf("Path: ")
	for i := 0; i < len(path.Nodes)-1; i++ {
		fmt.Printf("%s -> ", path.Nodes[i].Name)
	}

	fmt.Printf("%s\n", path.Nodes[len(path.Nodes)-1].Name)

	fmt.Println("Cost: ", path.Cost)

}
