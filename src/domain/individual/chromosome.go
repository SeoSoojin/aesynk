package individual

import (
	"math/rand"

	"github.com/seosoojin/aesynk/src/domain/node"
	"golang.org/x/exp/slices"
)

func NewChromosome(initial *node.Node, missing []string) Chromosome {

	size := len(missing)

	missingCopy := slices.Clone(missing)

	chromosome := make(Chromosome, size+1)

	currNode := initial

	for len(missingCopy) > 0 {

		random := rand.Intn(len(missingCopy))

		nextNode := currNode.AdjacentsMap[missingCopy[random]]

		chromosome[(size-len(missingCopy))+1] = nextNode.To

		currNode = nextNode.To

		missingCopy = append(missingCopy[:random], missingCopy[random+1:]...)

	}

	chromosome[0] = initial

	return chromosome

}

func (c Chromosome) Fitness() float64 {

	sum := 0.0

	for i := 0; i < len(c)-1; i++ {

		node := c[i]
		next := c[i+1]

		sum += node.AdjacentsMap[next.Name].Weight

	}

	sum += c[len(c)-1].AdjacentsMap[c[0].Name].Weight

	return sum

}
