package graph

import (
	"github.com/seosoojin/aesynk/src/domain/node"
)

type graph interface {
	FromCSV(path string) (graph, error)
	GenerateCompleteGraph(size int, randomize bool) graph
	ValidateCompleteGraph() bool
	Walk(start string) error
	Nodes() map[string]*node.Node
}
