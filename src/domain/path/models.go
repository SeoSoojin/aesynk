package path

import "github.com/seosoojin/aesynk/src/domain/node"

type Path struct {
	Nodes []*node.Node
	Cost  float64
}
