package algorithmsingo
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDijkstraShortestPath(t *testing.T) {
	t.Log("When using Dijkstra's algorithm to find the shortest path")
	{
		graph := *NewGraph()
		graph.AddVertex(1)
		graph.AddVertex(2)
		graph.AddVertex(3)
		graph.AddVertex(4)
		graph.AddVertex(5)
		graph.AddVertex(6)

		graph.AddDirectedEdge(1, 3, 5)
		graph.AddDirectedEdge(1, 2, 1)
		graph.AddDirectedEdge(2, 3, 1)
		graph.AddDirectedEdge(2, 5, 2)
		graph.AddDirectedEdge(3, 4, 1)
		graph.AddDirectedEdge(3, 6, 6)
		graph.AddDirectedEdge(5, 6, 3)

		shortestPath := DijkstraShortestPath(1, 6, graph)

		t.Log("it calculates the correct cost of the shortest path")
		{
			assert.Equal(t, shortestPath.Cost, 6.0)
		}

		t.Log("it calculates the correct shortest path")
		{
			assert.Equal(t, shortestPath.Path, []int{1, 2, 5, 6})
		}
	}
}
