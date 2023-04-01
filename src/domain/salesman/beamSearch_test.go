package salesman

import (
	"testing"

	"github.com/seosoojin/aesynk/src/domain/graph"
	"github.com/seosoojin/aesynk/src/domain/node"
)

func BenchmarkBeamSearchSolver_Solve(b *testing.B) {

	completeGraph10 := graph.NewGraph(true, 2).GenerateCompleteGraph(10, true)
	completeGraph30 := graph.NewGraph(true, 2).GenerateCompleteGraph(30, true)
	completeGraph50 := graph.NewGraph(true, 2).GenerateCompleteGraph(50, true)
	completeGraph100 := graph.NewGraph(true, 2).GenerateCompleteGraph(100, true)
	completeGraph150 := graph.NewGraph(true, 2).GenerateCompleteGraph(150, true)
	completeGraph200 := graph.NewGraph(true, 2).GenerateCompleteGraph(200, true)

	tests := []struct {
		name  string
		graph map[string]*node.Node
		width int
	}{
		{"10 nodes width 1", completeGraph10.Nodes(), 1},
		{"30 nodes width 1", completeGraph30.Nodes(), 1},
		{"50 nodes width 1", completeGraph50.Nodes(), 1},
		{"100 nodes width 1", completeGraph100.Nodes(), 1},
		{"150 nodes width 1", completeGraph150.Nodes(), 1},
		{"200 nodes width 1", completeGraph200.Nodes(), 1},
		{"10 nodes width 3", completeGraph10.Nodes(), 3},
		{"30 nodes width 3", completeGraph30.Nodes(), 3},
		{"50 nodes width 3", completeGraph50.Nodes(), 3},
		{"100 nodes width 3", completeGraph100.Nodes(), 3},
		{"150 nodes width 3", completeGraph150.Nodes(), 3},
		{"200 nodes width 3", completeGraph200.Nodes(), 3},
		{"10 nodes width 5", completeGraph10.Nodes(), 5},
		{"30 nodes width 5", completeGraph30.Nodes(), 5},
		{"50 nodes width 5", completeGraph50.Nodes(), 5},
		{"100 nodes width 5", completeGraph100.Nodes(), 5},
		{"150 nodes width 5", completeGraph150.Nodes(), 5},
		{"200 nodes width 5", completeGraph200.Nodes(), 5},
	}

	for _, t := range tests {

		s := NewBeamSearchSolver(t.graph, t.width)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Solve()
			}
		})
	}
}

func BenchmarkBeamSearchSolver_SolveFor1000(b *testing.B) {

	completeGraph10000 := graph.NewGraph(true, 2).GenerateCompleteGraph(1000, true)

	tests := []struct {
		name  string
		width int
	}{
		{"width 1", 1},
		{"width 2", 2},
		{"width 3", 3},
	}

	for _, t := range tests {

		s := NewBeamSearchSolver(completeGraph10000.Nodes(), t.width)

		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Solve()
			}
		})
	}

}
