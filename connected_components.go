package algorithmsingo

import (
	"github.com/deckarep/golang-set"
	"fmt"
)

func CalculateConnectedComponents(graph UndirectedGraph) [][]int {
	components := [][]int{}
	currentComponent := []int{}
	unexploredVertices := mapset.NewSet()
	for _, vertex := range graph.Vertices {
		unexploredVertices.Add(vertex)
	}
	unexploredVerticesAsSlice := unexploredVertices.ToSlice()
	for len(unexploredVerticesAsSlice) > 0 {
		vertex, _ := unexploredVerticesAsSlice[0].(int)
		exploreVertex(vertex, unexploredVertices, &currentComponent, graph)
		components = append(components, currentComponent)
		currentComponent = []int{}
		unexploredVerticesAsSlice = unexploredVertices.ToSlice()
	}

	return components
}

func exploreVertex(vertex int, unexploredVertices mapset.Set, currentComponent *[]int, graph UndirectedGraph) {
	fmt.Println(*currentComponent)
	if !unexploredVertices.Contains(vertex) {
		return
	}
	unexploredVertices.Remove(vertex)
	*currentComponent = append(*currentComponent, vertex)

	for _, edge := range graph.GetEdgesForVertex(vertex) {
		var otherVertex int
		if edge.V1 != vertex {
			otherVertex = edge.V1
		} else {
			otherVertex = edge.V2
		}
		exploreVertex(otherVertex, unexploredVertices, currentComponent, graph)
	}
}