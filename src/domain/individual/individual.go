package individual

import "math/rand"

func RoundSelect(idv Individual, oidv Individual) Individual {

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

	k := 0
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

	child := Individual{Chromosome: childchromosome, Fitness: childchromosome.Fitness()}
	sChild := Individual{Chromosome: sChildchromosome, Fitness: sChildchromosome.Fitness()}

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
