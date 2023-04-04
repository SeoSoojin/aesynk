package salesman

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/seosoojin/aesynk/src/domain/individual"
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
	"github.com/seosoojin/aesynk/src/domain/utils"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type GeneticSolver struct {
	graph               map[string]*node.Node
	generations         int
	populationSize      int
	eliteGenerationSize int
	mutationProb        float64
}

func NewGeneticSolver(graph map[string]*node.Node, generations int, populationSize int, elitismPercent float64, mutationProb float64) Solver {

	if elitismPercent > 1 {
		elitismPercent = 1
	}

	eliteGenerationSize := int(float64(populationSize) * elitismPercent)

	return &GeneticSolver{
		graph:               graph,
		generations:         generations,
		populationSize:      populationSize,
		eliteGenerationSize: eliteGenerationSize,
		mutationProb:        mutationProb,
	}

}

func (g *GeneticSolver) Solve() (path.Path, error) {

	initialPopulation := randomPopulation(g.graph, g.populationSize)
	if len(initialPopulation) == 0 {
		return path.Path{}, fmt.Errorf("error generating initial population")
	}

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

	sort.SliceStable(base, func(i, j int) bool {
		return base[i].Fitness < base[j].Fitness
	})

	survivingElites := int(float64(g.eliteGenerationSize) * 0.8)

	children := slices.Clone(base[:survivingElites])

	eliteBase := base[:g.eliteGenerationSize]

	nonEliteBase := base[g.eliteGenerationSize:]

	for len(children) < g.populationSize {

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

		child1, _ := parent.Breed(partner, g.mutationProb)

		children = append(children, child1)

	}

	return children[:g.populationSize]

}

func tournamentSelect(population []individual.Individual) individual.Individual {

	k := 2 + rand.Intn(10)

	minIndex := -1

	for j := 0; j < k; j++ {

		aux := rand.Intn(len(population))
		if minIndex == -1 || aux < minIndex {
			minIndex = aux
		}

	}

	return population[minIndex]

}

func randomPopulation(graph map[string]*node.Node, size int) []individual.Individual {

	firstNode := utils.GenerateRandomNode(graph)

	if firstNode == nil {
		return []individual.Individual{}
	}

	graphCopy := maps.Clone(graph)

	delete(graphCopy, firstNode.Name)

	missing := maps.Keys(graphCopy)

	population := make([]individual.Individual, size)

	currNode := firstNode

	for i := 0; i < size; i++ {

		chromosome := individual.NewChromosome(currNode, missing)
		population[i] = individual.NewIndividual(chromosome)

	}

	return population

}

func validateSolutionGenetic(input map[string]*node.Node, output path.Path) bool {

	if len(output.Nodes) != len(input)+1 {
		return false
	}

	nodesMap := map[string]struct{}{}

	for _, node := range output.Nodes {
		nodesMap[node.Name] = struct{}{}
	}

	return len(nodesMap) == len(input) && output.Nodes[0].Name == output.Nodes[len(output.Nodes)-1].Name

}
