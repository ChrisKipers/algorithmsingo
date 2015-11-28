package algorithmsingo

import (
	"github.com/deckarep/golang-set"
)

func TopologicalSort(graph Graph) []int {
	// Assume graph is connected
	exploredVertices := mapset.NewSet()
	sortedVertices := []int{}

	for _, vertex := range graph.Vertices {
		if exploredVertices.Contains(vertex) {
			continue
		} else {
			tsExploreVertex(vertex, graph, exploredVertices, &sortedVertices)
		}
	}
	return sortedVertices
}

func tsExploreVertex(vertex int, graph Graph, exploredVertices mapset.Set, sortedVertices *[]int) {
	vertexEdges := graph.GetEdgesForVertex(vertex)
	for _, edge := range vertexEdges {
		if !exploredVertices.Contains(edge.V2) {
			tsExploreVertex(edge.V2, graph, exploredVertices, sortedVertices)
		}
	}
	*sortedVertices = append([]int{vertex}, *sortedVertices...)
	exploredVertices.Add(vertex)
}