package salesman

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/seosoojin/aesynk/src/domain/individual"
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type GeneticSolver struct {
	graph                  map[string]*node.Node
	generations            int
	populationSize         int
	eliteGenerationSize    int
	remainigGenerationSize int
	mutationProb           float64
}

func NewGeneticSolver(graph map[string]*node.Node, generations int, populationSize int, elitismPercent float64, mutationProb float64) Solver {

	if elitismPercent > 1 {
		elitismPercent = 1
	}

	eliteGenerationSize := int(float64(populationSize/2) * elitismPercent)

	return &GeneticSolver{
		graph:                  graph,
		generations:            generations,
		populationSize:         populationSize,
		eliteGenerationSize:    eliteGenerationSize,
		remainigGenerationSize: (populationSize / 2) - eliteGenerationSize,
		mutationProb:           mutationProb,
	}

}

func (g *GeneticSolver) Solve() (path.Path, error) {

	initialPopulation := randomPopulation(g.graph, g.populationSize)

	finalpop := g.solve(initialPopulation)

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

func (g *GeneticSolver) solve(population []individual.Individual) []individual.Individual {

	currPopulation := population

	for i := 0; i < g.generations; i++ {

		currPopulation = g.reproduce(currPopulation)

	}

	return currPopulation

}

func (g *GeneticSolver) reproduce(base []individual.Individual) []individual.Individual {

	parents := g.selectParents(base)

	children := slices.Clone(parents)

	for i := 0; i < g.remainigGenerationSize; i += 2 {

		parentIndex := rand.Intn(len(parents))
		parent := parents[parentIndex]

		partnerIndex := rand.Intn(len(parents))

		for partnerIndex == parentIndex {
			partnerIndex = rand.Intn(len(parents))
		}

		partner := parents[partnerIndex]

		child1, child2 := parent.Breed(partner, g.mutationProb)

		children = append(children, child1, child2)

	}

	return children

}

func randomPopulation(graph map[string]*node.Node, size int) []individual.Individual {

	firstNode := generateRandomNode(graph)

	graphCopy := maps.Clone(graph)

	delete(graphCopy, firstNode.Name)

	missing := maps.Keys(graphCopy)

	population := make([]individual.Individual, size)

	currNode := firstNode

	for i := 0; i < size; i++ {

		chromosome := individual.NewChromosome(currNode, missing)
		population[i] = individual.Individual{Chromosome: chromosome, Fitness: chromosome.Fitness()}

	}

	return population

}

func (g GeneticSolver) selectParents(population []individual.Individual) []individual.Individual {

	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})

	elitePopulation := population[:g.eliteGenerationSize]

	if g.remainigGenerationSize == 0 {
		return elitePopulation
	}

	remainingPopulation := population[g.eliteGenerationSize:]

	remainingParents := make([]individual.Individual, g.remainigGenerationSize)

	for i := 0; i < g.remainigGenerationSize; i++ {

		remainingParents[i] = individual.RoundSelect(remainingPopulation[rand.Intn(len(remainingPopulation))], remainingPopulation[rand.Intn(len(remainingPopulation))])

	}

	return append(elitePopulation, remainingParents...)

}
