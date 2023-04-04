package salesman

import (
	"testing"

	"github.com/seosoojin/aesynk/src/domain/graph"
	"github.com/seosoojin/aesynk/src/domain/node"
)

func BenchmarkGeneticSolver_Solve(b *testing.B) {

	completeGraph10, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph10.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph30, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph30.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph50, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph50.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph100, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph100.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph150, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph150.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph200, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph200.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph500, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph500.csv")
	if err != nil {
		b.Fatal(err)
	}

	completeGraph1000, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph1000.csv")
	if err != nil {
		b.Fatal(err)
	}

	tests := []struct {
		name           string
		graph          map[string]*node.Node
		populationSize int
		generationSize int
	}{
		{
			name:           "10 nodes",
			graph:          completeGraph10.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "30 nodes",
			graph:          completeGraph30.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "50 nodes",
			graph:          completeGraph50.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "100 nodes",
			graph:          completeGraph100.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "150 nodes",
			graph:          completeGraph150.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "200 nodes",
			graph:          completeGraph200.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "500 nodes",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "1000 nodes",
			graph:          completeGraph1000.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
	}

	for _, t := range tests {

		s := NewGeneticSolver(t.graph, t.generationSize, t.populationSize, 0.2, 0.1)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Solve()
			}
		})
	}
}

func BenchmarkGeneticSolver_SolveVPopulation(b *testing.B) {

	completeGraph500, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph500.csv")
	if err != nil {
		b.Fatal(err)
	}

	tests := []struct {
		name           string
		graph          map[string]*node.Node
		populationSize int
		generationSize int
	}{
		{
			name:           "population 100",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "population 250",
			graph:          completeGraph500.Nodes(),
			populationSize: 250,
			generationSize: 100,
		},
		{
			name:           "population 500",
			graph:          completeGraph500.Nodes(),
			populationSize: 500,
			generationSize: 100,
		},
		{
			name:           "population 1000",
			graph:          completeGraph500.Nodes(),
			populationSize: 1000,
			generationSize: 100,
		},
		{
			name:           "population 1500",
			graph:          completeGraph500.Nodes(),
			populationSize: 1500,
			generationSize: 100,
		},
	}

	for _, t := range tests {

		s := NewGeneticSolver(t.graph, t.generationSize, t.populationSize, 0.2, 0.1)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Solve()
			}
		})
	}
}

func BenchmarkGeneticSolver_SolveVGeneration(b *testing.B) {

	completeGraph500, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph500.csv")
	if err != nil {
		b.Fatal(err)
	}

	tests := []struct {
		name           string
		graph          map[string]*node.Node
		populationSize int
		generationSize int
	}{
		{
			name:           "generation 100",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 100,
		},
		{
			name:           "generation 250",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 250,
		},
		{
			name:           "generation 500",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 500,
		},
		{
			name:           "generation 1000",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 1000,
		},
		{
			name:           "generation 1500",
			graph:          completeGraph500.Nodes(),
			populationSize: 100,
			generationSize: 1500,
		},
	}

	for _, t := range tests {

		s := NewGeneticSolver(t.graph, t.generationSize, t.populationSize, 0.2, 0.1)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := s.Solve()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

}

func BenchmarkGenetic_SolveFor2000(b *testing.B) {

	completeGraph2000, err := graph.NewGraph(true, 2).FromCSV("../../../testdata/graph2000.csv")
	if err != nil {
		b.Fatal(err)
	}

	tests := []struct {
		name           string
		graph          map[string]*node.Node
		populationSize int
	}{
		{
			name:           "population 100",
			graph:          completeGraph2000.Nodes(),
			populationSize: 100,
		},
		{
			name:           "population 250",
			graph:          completeGraph2000.Nodes(),
			populationSize: 250,
		},
		{
			name:           "population 500",
			graph:          completeGraph2000.Nodes(),
			populationSize: 500,
		},
		{
			name:           "population 1000",
			graph:          completeGraph2000.Nodes(),
			populationSize: 1000,
		},
		{
			name:           "population 1500",
			graph:          completeGraph2000.Nodes(),
			populationSize: 1500,
		},
	}

	for _, t := range tests {

		s := NewGeneticSolver(t.graph, 100, t.populationSize, 0.2, 0.1)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Solve()
			}
		})
	}

}
