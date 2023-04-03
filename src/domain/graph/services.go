package graph

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/seosoojin/aesynk/src/domain/node"
)

type Graph struct {
	NodeMap    map[string]*node.Node
	Dimensions int
	Symmetric  bool
}

func NewGraph(symmetric bool, dimensions int) graph {

	return &Graph{
		NodeMap:    make(map[string]*node.Node),
		Symmetric:  symmetric,
		Dimensions: dimensions,
	}

}

func (g *Graph) FromCSV(path string) (graph, error) {

	f, err := os.Open(path)
	if err != nil {
		return g, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()
	if err != nil {
		return g, err
	}

	linesWithoutHeader := lines[1:]

	nodesMap := make(map[string]*node.Node)

	for _, line := range linesWithoutHeader {

		coordinatesString := line[1]

		coordinatesString = strings.Trim(coordinatesString, " []")

		cordinatesArray := strings.Split(coordinatesString, ",")

		if len(cordinatesArray) != g.Dimensions {
			return g, fmt.Errorf("invalid number of coordinates")
		}

		coordinates := make([]float64, g.Dimensions)

		for i, coordinate := range cordinatesArray {

			coordinate = strings.Trim(coordinate, " ")

			coordinates[i], err = strconv.ParseFloat(coordinate, 64)
			if err != nil {
				return g, err
			}

		}

		node := node.NewNode(line[0], coordinates)

		nodesMap[node.Name] = node

	}

	for _, line := range linesWithoutHeader {

		auxNode := nodesMap[line[0]]

		adjascents := line[len(line)-1]

		adjascents = strings.Trim(adjascents, " []")

		for _, adjacent := range strings.Split(adjascents, ",") {

			adjacent = strings.Trim(adjacent, " ")
			auxNode.AddAdjacent(nodesMap[adjacent])
			if g.Symmetric {
				nodesMap[adjacent].AddAdjacent(auxNode)
			}

		}

	}

	g.NodeMap = nodesMap

	return g, nil

}

func (g *Graph) Walk(start string) error {

	startNode, ok := g.NodeMap[start]
	if !ok {
		return fmt.Errorf("node %s not found", start)
	}

	visited := make(map[string]struct{})

	walk(startNode, visited, 0)

	return nil

}

func walk(curr *node.Node, visited map[string]struct{}, depth int) {

	if _, ok := visited[curr.Name]; ok {
		return
	}

	builder := strings.Builder{}

	for i := 0; i < depth; i++ {
		builder.WriteString(" ")
	}

	builder.WriteString(curr.Name)

	fmt.Println(builder.String())

	visited[curr.Name] = struct{}{}

	for _, adjacent := range curr.Adjacents {

		walk(adjacent.To, visited, depth+1)

	}

}

func (g *Graph) Nodes() map[string]*node.Node {
	return g.NodeMap
}

func (g *Graph) GenerateCompleteGraph(size int, randomize bool) graph {

	nodesMap := make(map[string]*node.Node)

	for i := 0; i < size; i++ {

		coordinates := make([]float64, g.Dimensions)

		for j := range coordinates {

			if randomize {
				coordinates[j] = rand.Float64()
				continue
			}

			coordinates[j] = float64(i) + (float64(j) / 10)

		}

		node := node.NewNode(strconv.Itoa(i), coordinates)

		nodesMap[node.Name] = node

	}

	for _, node := range nodesMap {

		for _, otherNode := range nodesMap {

			if node.Name == otherNode.Name {
				continue
			}

			node.AddAdjacent(otherNode)

		}

	}

	g.NodeMap = nodesMap

	return g

}

func (g *Graph) ToCSV(path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)

	header := []string{"Name", "Coordinates", "Adjacents"}

	lines := [][]string{header}

	for _, node := range g.NodeMap {

		line := []string{node.Name}
		builder := strings.Builder{}

		for i := 0; i < len(node.Coordinates)-1; i++ {

			builder.WriteString(fmt.Sprintf("%f,", node.Coordinates[i]))

		}

		builder.WriteString(fmt.Sprintf("%f", node.Coordinates[len(node.Coordinates)-1]))

		line = append(line, builder.String())

		builder.Reset()

		for i := 0; i < len(node.Adjacents)-1; i++ {

			builder.WriteString(fmt.Sprintf("%s,", node.Adjacents[i].To.Name))

		}

		builder.WriteString(node.Adjacents[len(node.Adjacents)-1].To.Name)

		line = append(line, builder.String())

		lines = append(lines, line)

	}

	return writer.WriteAll(lines)

}

func (g *Graph) ValidateCompleteGraph() bool {

	for _, node := range g.NodeMap {

		if len(node.Adjacents) != len(g.NodeMap)-1 {
			return false
		}

		for _, adjacent := range node.Adjacents {

			if adjacent.To.Name == node.Name {
				return false
			}

		}

	}

	return true

}
