package salesman

import (
	"fmt"
	"math/rand"
	"sort"

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
			graph:               graph,
			generations:         generations,
			populationSize:      populationSize,
			eliteGenerationSize: eliteGenerationSize,
			mutationProb:        mutationProb,
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

	sort.SliceStable(base, func(i, j int) bool {
		return base[i].Fitness < base[j].Fitness
	})

	survivingElites := int(float64(g.eliteGenerationSize) * 0.8)

	children := slices.Clone(base[:survivingElites])

	eliteBase := base[:g.eliteGenerationSize]

	nonEliteBase := base[g.eliteGenerationSize:]

	for len(children) < g.populationSize {

		childCh := make(chan individual.Individual)

		batchChildren := make([]individual.Individual, 0)

		for j := 0; j < g.batch/2; j++ {

			go func() {
				var parents []individual.Individual

				if rand.Float64() < 0.6 {

					parents = eliteBase

				} else {

					parents = nonEliteBase

				}

				parent := tournamentSelect(parents)

				if rand.Float64() < 0.6 {

					parents = eliteBase

				} else {

					parents = nonEliteBase

				}

				partner := tournamentSelect(parents)
				for partner.Fitness == parent.Fitness {
					partner = tournamentSelect(parents)
				}

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

	return children

}
