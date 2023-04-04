package individual

import "math/rand"

func NewIndividual(chromosome Chromosome) Individual {

	return Individual{
		Chromosome: chromosome,
		Fitness:    chromosome.Fitness(),
	}

}

func RoundSelect(idv Individual, oidv Individual) Individual {

	if idv.Fitness < oidv.Fitness {
		return idv
	}

	return oidv

}

func (parent Individual) Breed(partner Individual, mutationProb float64) (Individual, Individual) {

	childCh, sChildCh := parent.Chromosome.MappedCrossoverRandomIndex(partner.Chromosome)

	child := NewIndividual(childCh)
	sChild := NewIndividual(sChildCh)

	child.Mutate(mutationProb)
	sChild.Mutate(mutationProb)

	return child, sChild

}

func (idv *Individual) Mutate(mutationProb float64) {

	for i := 0; i < len(idv.Chromosome); i++ {

		mutate := rand.Float64() < mutationProb

		if mutate {
			randomIndex := rand.Intn(len(idv.Chromosome))

			aux := idv.Chromosome[i]

			idv.Chromosome[i] = idv.Chromosome[randomIndex]
			idv.Chromosome[randomIndex] = aux

		}

	}

}
