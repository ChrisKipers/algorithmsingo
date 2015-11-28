package algorithmsingo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	t.Log("When topologically sorting a graph")
	{
		graph := NewGraph()
		graph.AddVertex(5)
		graph.AddVertex(6)
		graph.AddVertex(7)
		graph.AddVertex(8)
		graph.AddVertex(1)
		graph.AddVertex(2)
		graph.AddVertex(3)
		graph.AddVertex(4)

		graph.AddDirectedEdge(1, 2, 0)
		graph.AddDirectedEdge(1, 3, 0)
		graph.AddDirectedEdge(2, 4, 0)
		graph.AddDirectedEdge(3, 4, 0)
		graph.AddDirectedEdge(4, 5, 0)
		graph.AddDirectedEdge(5, 6, 0)
		graph.AddDirectedEdge(5, 7, 0)
		graph.AddDirectedEdge(6, 8, 0)
		graph.AddDirectedEdge(7, 8, 0)

		sortedVertices := TopologicalSort(*graph)

		t.Log("it orders the vertices based on dependencies")
		{
			assert.True(t, isVertexBeforOtherVertex(1, 2, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(1, 3, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(2, 4, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(3, 4, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(4, 5, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(5, 6, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(5, 7, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(6, 8, sortedVertices))
			assert.True(t, isVertexBeforOtherVertex(7, 8, sortedVertices))
		}
	}
}


func isVertexBeforOtherVertex(vertex int, otherVertex int, sortedVertices []int) bool {
	for _, currentVertex := range sortedVertices {
		if currentVertex == vertex {
			return true
		} else if currentVertex == otherVertex {
			return false
		}
	}
	return true
}
