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

func (c Chromosome) MappedCrossoverRandomIndex(other Chromosome) (Chromosome, Chromosome) {

	cGenesSize := 1 + rand.Intn(len(c)/2)
	cIndexes := make(map[int]struct{}, cGenesSize)
	cGenesMap := make(map[string]struct{})

	childCh := make(Chromosome, len(c))
	schildCh := make(Chromosome, len(c))

	for i := 0; i < cGenesSize; i++ {
		random := 1 + rand.Intn(len(c)-1)
		cIndexes[random] = struct{}{}
	}

	for i := 0; i < len(c); i++ {
		_, ok := cIndexes[i]
		if ok {
			childCh[i] = c[i]
			cGenesMap[c[i].Name] = struct{}{}
		}
		if !ok {
			schildCh[i] = c[i]
		}
	}

	j := 0

	for i := 0; i < len(childCh); i++ {

		if childCh[i] != nil {
			continue
		}

		_, ok := cGenesMap[other[j].Name]
		for ok {
			j++
			_, ok = cGenesMap[other[j].Name]
		}

		childCh[i] = other[j]
		j++

	}

	k := 0
	for i := 0; i < len(schildCh); i++ {

		if schildCh[i] != nil {
			continue
		}

		_, ok := cGenesMap[other[k].Name]
		for !ok {
			k++
			_, ok = cGenesMap[other[k].Name]
		}

		schildCh[i] = other[k]
		k++

	}

	return childCh, schildCh

}

func (c Chromosome) Mapped2PointCrossover(other Chromosome) (Chromosome, Chromosome) {

	cGenesSize := 1 + rand.Intn(len(c)/2)
	geneIndex := rand.Intn(len(c) - cGenesSize)

	cGenesMap := make(map[string]struct{})

	childCh := make(Chromosome, len(c))
	schildCh := make(Chromosome, len(c))

	for i := 0; i < len(c); i++ {
		if i >= geneIndex && i < geneIndex+cGenesSize {
			childCh[i] = c[i]
			cGenesMap[c[i].Name] = struct{}{}
		} else {
			schildCh[i] = c[i]
		}
	}

	j := 0

	for i := 0; i < len(childCh); i++ {

		if childCh[i] != nil {
			continue
		}

		_, ok := cGenesMap[other[j].Name]
		for ok {
			j++
			_, ok = cGenesMap[other[j].Name]
		}

		childCh[i] = other[j]
		j++

	}

	k := 0
	for i := 0; i < len(schildCh); i++ {

		if schildCh[i] != nil {
			continue
		}

		_, ok := cGenesMap[other[k].Name]
		for !ok {
			k++
			_, ok = cGenesMap[other[k].Name]
		}

		schildCh[i] = other[k]
		k++

	}

	return childCh, schildCh

}
