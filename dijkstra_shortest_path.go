package algorithmsingo
import (
	"github.com/deckarep/golang-set"
	"math"
)

type ShortestPath struct {
	Cost float64
	Path []int
}

func DijkstraShortestPath(startVertex int, endVertex int, graph Graph) ShortestPath {
	exploredVertices := mapset.NewSet()
	exploredVertices.Add(startVertex)
	edgesOnBorder := []Edge{}
	for _, edge := range graph.GetEdgesForVertex(startVertex) {
		edgesOnBorder = append(edgesOnBorder, edge)
	}
	shortestPathByVertex := make(map[int]ShortestPath)
	shortestPathByVertex[startVertex] = ShortestPath{
		Cost: 0,
		Path: []int{startVertex},
	}

	for !exploredVertices.Contains(endVertex) {
		indexOfNextShortestEdge := indexOfShortestEdgeOnBorder(edgesOnBorder, shortestPathByVertex)
		nextShortestEdge := edgesOnBorder[indexOfNextShortestEdge]
		precedingPath := shortestPathByVertex[nextShortestEdge.V1]
		edgesPath := append([]int(nil), precedingPath.Path...)
		edgesPath = append(edgesPath, nextShortestEdge.V2)
		shortestPathForVertex := ShortestPath{
			Cost: precedingPath.Cost + nextShortestEdge.weight,
			Path: edgesPath,
		}
		shortestPathByVertex[nextShortestEdge.V2] = shortestPathForVertex
		edgesOnBorder = append(edgesOnBorder[:indexOfNextShortestEdge], edgesOnBorder[indexOfNextShortestEdge + 1:]...)
		exploredVertices.Add(nextShortestEdge.V2)

		for _, edge := range graph.GetEdgesForVertex(nextShortestEdge.V2) {
			if !exploredVertices.Contains(edge.V2) {
				edgesOnBorder = append(edgesOnBorder, edge)
			}
		}
	}

	return shortestPathByVertex[endVertex]
}

func indexOfShortestEdgeOnBorder(edges []Edge, shortestPathByVertex map[int]ShortestPath) int {
	indexOfLeastExpensive := -1
	costOfLeastExpensive := math.MaxFloat64

	for i, edge := range edges {
		costOfPrecedingPath := shortestPathByVertex[edge.V1].Cost
		if costOfPrecedingPath + edge.weight < costOfLeastExpensive {
			indexOfLeastExpensive = i
		}
	}

	return indexOfLeastExpensive
}