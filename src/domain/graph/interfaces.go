package graph

import (
	"github.com/seosoojin/aesynk/src/domain/node"
)

type graph interface {
	FromCSV(path string) (graph, error)
	Walk(start string) error
	Nodes() map[string]*node.Node
}
