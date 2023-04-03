package salesman

import (
	"fmt"
	"math/rand"

	"github.com/seosoojin/aesynk/src/domain/individual"
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/slices"
)

type PGeneticSolver struct {
	GeneticSolver
	batch int
}

func NewPGeneticSolver(graph map[string]*node.Node, generations int, populationSize int, elitismPercent float64, mutationProb float64, batch int) Solver {

	if elitismPercent > 1 {
		elitismPercent = 1
	}

	eliteGenerationSize := int(float64(populationSize/2) * elitismPercent)

	switch {
	case batch > 200:
		batch = 200
	case batch > populationSize:
		batch = populationSize / 2
	case batch < 1:
		batch = 0
	}

	if batch%2 != 0 {
		batch++
	}

	return &PGeneticSolver{
		GeneticSolver: GeneticSolver{
			graph:                  graph,
			generations:            generations,
			populationSize:         populationSize,
			eliteGenerationSize:    eliteGenerationSize,
			remainigGenerationSize: (populationSize / 2) - eliteGenerationSize,
			mutationProb:           mutationProb,
		},
		batch: batch,
	}

}

func (g *PGeneticSolver) Solve() (path.Path, error) {

	initialPopulation := randomPopulation(g.graph, g.populationSize)

	finalpop := g.parallelSolve(initialPopulation)

	nodes := finalpop[0].Chromosome

	nodes = append(nodes, nodes[0])

	out := path.Path{
		Nodes: nodes,
		Cost:  finalpop[0].Fitness,
	}

	ok := validateSolutionGenetic(g.graph, out)

	if !ok {
		return path.Path{}, fmt.Errorf("invalid solution")
	}

	return out, nil

}

func (g *PGeneticSolver) parallelSolve(population []individual.Individual) []individual.Individual {

	currPopulation := population

	for i := 0; i < g.generations; i++ {

		currPopulation = g.parallelReproduce(currPopulation, g.mutationProb)

	}

	return currPopulation

}

func (g *PGeneticSolver) parallelReproduce(base []individual.Individual, mutationProb float64) []individual.Individual {

	parents := g.selectParents(base)

	children := slices.Clone(parents)

	for i := 0; i < g.remainigGenerationSize; i += g.batch {

		childCh := make(chan individual.Individual)

		batchChildren := make([]individual.Individual, 0)

		for j := 0; j < g.batch/2; j++ {

			go func() {
				parentIndex := rand.Intn(len(parents))
				parent := parents[parentIndex]

				partnerIndex := rand.Intn(len(parents))

				for partnerIndex == parentIndex {
					partnerIndex = rand.Intn(len(parents))
				}

				partner := parents[partnerIndex]

				child1, child2 := parent.Breed(partner, g.mutationProb)

				childCh <- child1
				childCh <- child2

			}()

		}

		for child := range childCh {

			batchChildren = append(batchChildren, child)

			if len(batchChildren) == g.batch {
				children = append(children, batchChildren...)
				close(childCh)
			}

		}

	}

	return append(children[:g.remainigGenerationSize], parents...)

}
