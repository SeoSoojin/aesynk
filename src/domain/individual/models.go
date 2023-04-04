package individual

import "github.com/seosoojin/aesynk/src/domain/node"

type Chromosome []*node.Node

type Individual struct {
	Chromosome Chromosome
	Fitness    float64
	Age        int
}
