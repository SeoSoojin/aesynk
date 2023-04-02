package salesman

import (
	"fmt"
	"math/rand"
	"sort"

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

	currPopulation := initialPopulation

	for i := 0; i < g.generations; i++ {

		parents := g.selectParents(currPopulation)

		children := slices.Clone(parents)

		for j := 0; j < g.remainigGenerationSize; j += 2 {

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

		currPopulation = children

	}

	nodes := make([]*node.Node, len(currPopulation[0].Chromosome)+1)

	for i := 0; i < len(nodes)-1; i++ {

		nodes[i] = currPopulation[0].Chromosome[i]

	}

	nodes[len(nodes)-1] = currPopulation[0].Chromosome[0]

	out := path.Path{
		Nodes: nodes,
		Cost:  currPopulation[0].Fitness,
	}

	ok := validateSolutionGenetic(g.graph, out)

	if !ok {
		return path.Path{}, fmt.Errorf("invalid solution")
	}

	return out, nil

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

func randomPopulation(graph map[string]*node.Node, size int) []Individual {

	firstNode := generateRandomNode(graph)

	graphCopy := maps.Clone(graph)

	delete(graphCopy, firstNode.Name)

	missing := maps.Keys(graphCopy)

	population := make([]Individual, size)

	currNode := firstNode

	for i := 0; i < size; i++ {

		chromosome := newChromosome(currNode, missing)
		population[i] = Individual{Chromosome: chromosome, Fitness: chromosome.Fitness()}

	}

	return population

}

func newChromosome(initial *node.Node, missing []string) Chromosome {

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

func (g GeneticSolver) selectParents(population []Individual) []Individual {

	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})

	elitePopulation := population[:g.eliteGenerationSize]

	if g.remainigGenerationSize == 0 {
		return elitePopulation
	}

	remainingPopulation := population[g.eliteGenerationSize:]

	remainingParents := make([]Individual, g.remainigGenerationSize)

	for i := 0; i < g.remainigGenerationSize; i++ {

		remainingParents[i] = roundSelect(remainingPopulation[rand.Intn(len(remainingPopulation))], remainingPopulation[rand.Intn(len(remainingPopulation))])

	}

	return append(elitePopulation, remainingParents...)

}

func roundSelect(idv Individual, oidv Individual) Individual {

	if idv.Fitness < oidv.Fitness {
		return idv
	}

	return oidv

}

func (parent Individual) Breed(partner Individual, mutationProb float64) (Individual, Individual) {

	parentGenesSize := 1 + rand.Intn(len(parent.Chromosome)-1)
	parentIndexes := make(map[int]struct{}, parentGenesSize)
	parentGenesMap := make(map[string]struct{})

	childchromosome := make(Chromosome, len(parent.Chromosome))
	sChildchromosome := make(Chromosome, len(parent.Chromosome))

	for i := 0; i < parentGenesSize; i++ {
		random := 1 + rand.Intn(len(parent.Chromosome)-1)
		parentIndexes[random] = struct{}{}
	}

	for i := 0; i < len(parent.Chromosome); i++ {
		_, ok := parentIndexes[i]
		if ok {
			childchromosome[i] = parent.Chromosome[i]
			parentGenesMap[parent.Chromosome[i].Name] = struct{}{}
		}
		if !ok {
			sChildchromosome[i] = parent.Chromosome[i]
		}
	}

	j := 0
	k := 0
	for i := 0; i < len(childchromosome); i++ {

		if childchromosome[i] != nil {
			continue
		}

		_, ok := parentGenesMap[partner.Chromosome[j].Name]
		for ok {
			j++
			_, ok = parentGenesMap[partner.Chromosome[j].Name]
		}

		childchromosome[i] = partner.Chromosome[j]
		j++

	}

	for i := 0; i < len(childchromosome); i++ {

		if sChildchromosome[i] != nil {
			continue
		}

		_, ok := parentGenesMap[partner.Chromosome[k].Name]
		for !ok {
			k++
			_, ok = parentGenesMap[partner.Chromosome[k].Name]
		}

		sChildchromosome[i] = partner.Chromosome[k]
		k++

	}

	for i := 0; i < len(childchromosome); i++ {

		mutateChild := rand.Float64() < mutationProb
		mutateSChild := rand.Float64() < mutationProb

		if mutateChild {
			randomIndex := rand.Intn(len(childchromosome))

			aux := childchromosome[i]

			childchromosome[i] = childchromosome[randomIndex]
			childchromosome[randomIndex] = aux

		}

		if mutateSChild {
			randomIndex := rand.Intn(len(sChildchromosome))

			aux := sChildchromosome[i]

			sChildchromosome[i] = sChildchromosome[randomIndex]
			sChildchromosome[randomIndex] = aux
		}

	}

	return Individual{Chromosome: childchromosome, Fitness: childchromosome.Fitness()}, Individual{Chromosome: sChildchromosome, Fitness: sChildchromosome.Fitness()}

}
