package main

import (
	"fmt"

	"github.com/seosoojin/aesynk/src/domain/graph"
	"github.com/seosoojin/aesynk/src/domain/salesman"
)

func main() {
	grafo, err := graph.NewGraph(true, 2).FromCSV("testdata/graph500.csv")
	if err != nil {
		panic(err)
	}

	// pathBeam, err := salesman.NewBeamSearchSolver(grafo.Nodes(), 3).Solve()
	// if err != nil {
	// 	panic(err)
	// }

	pathGenetic, err := salesman.NewGeneticSolver(grafo.Nodes(), 700, 2000, 0.45, 1).Solve()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Beam Search solution cost: ", pathBeam.Cost)
	fmt.Println("Genetic solution cost: ", pathGenetic.Cost)

}
