package algorithmsingo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCut(t *testing.T) {
	graph := *NewGraph()
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)
	graph.AddVertex(8)

	graph.AddEdge(1, 2, 0)
	graph.AddEdge(1, 3, 0)
	graph.AddEdge(1, 4, 0)
	graph.AddEdge(2, 3, 0)
	graph.AddEdge(2, 4, 0)
	graph.AddEdge(3, 4, 0)
	graph.AddEdge(3, 5, 0)
	graph.AddEdge(4, 6, 0)
	graph.AddEdge(5, 7, 0)
	graph.AddEdge(5, 6, 0)
	graph.AddEdge(5, 8, 0)
	graph.AddEdge(6, 7, 0)
	graph.AddEdge(6, 8, 0)
	graph.AddEdge(7, 8, 0)

	t.Log("When getting the miniumum cut for a graph")
	{
		minCut := CalculateMinCut(graph, 2)
		fmt.Println(minCut)

		t.Log("it returns the correct minimum cuts")
		{
			var group1 VertexGroup
			var group2 VertexGroup

			if groupContainsVertex(minCut.Groups[0], 1) {
				group1 = minCut.Groups[0]
				group2 = minCut.Groups[1]
			} else {
				group1 = minCut.Groups[1]
				group2 = minCut.Groups[0]
			}
			assert.True(t, groupContainsVertex(group1, 1))
			assert.True(t, groupContainsVertex(group1, 2))
			assert.True(t, groupContainsVertex(group1, 3))
			assert.True(t, groupContainsVertex(group1, 4))

			assert.True(t, groupContainsVertex(group2, 5))
			assert.True(t, groupContainsVertex(group2, 6))
			assert.True(t, groupContainsVertex(group2, 7))
			assert.True(t, groupContainsVertex(group2, 8))

			assert.Equal(t, 2, minCut.NumberOfEdges)
		}
	}

	t.Log("When getting the miniumum cut for a graph in parallel")
	{
		minCut := CalculateMinCutPar(graph, 2)
		fmt.Println(minCut)

		t.Log("it returns the correct minimum cuts")
		{
			var group1 VertexGroup
			var group2 VertexGroup

			if groupContainsVertex(minCut.Groups[0], 1) {
				group1 = minCut.Groups[0]
				group2 = minCut.Groups[1]
			} else {
				group1 = minCut.Groups[1]
				group2 = minCut.Groups[0]
			}
			assert.True(t, groupContainsVertex(group1, 1))
			assert.True(t, groupContainsVertex(group1, 2))
			assert.True(t, groupContainsVertex(group1, 3))
			assert.True(t, groupContainsVertex(group1, 4))

			assert.True(t, groupContainsVertex(group2, 5))
			assert.True(t, groupContainsVertex(group2, 6))
			assert.True(t, groupContainsVertex(group2, 7))
			assert.True(t, groupContainsVertex(group2, 8))

			assert.Equal(t, 2, minCut.NumberOfEdges)
		}
	}
}

func groupContainsVertex(group VertexGroup, vertex int) bool {
	contains := false
	for _, curVertex := range group.Vertices {
		if vertex == curVertex {
			contains = true
			break
		}
	}
	return contains
}
