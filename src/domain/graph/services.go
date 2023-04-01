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

		coordinatesFields := line[1 : 1+g.Dimensions]

		coordinates := make([]float64, g.Dimensions)

		for index, value := range coordinatesFields {

			value := strings.Trim(value, " ")

			floatCoord, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return g, err
			}
			coordinates[index] = floatCoord

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
