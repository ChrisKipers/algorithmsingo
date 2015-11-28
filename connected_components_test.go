package algorithmsingo

import (
//	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestConnectedComponents(t *testing.T)  {
	t.Log("When calculating connected components")
	{
		graph := NewUndriectedGraph()
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

		graph.AddEdge(5, 6, 0)
		graph.AddEdge(6, 7, 0)
		graph.AddEdge(7, 8, 0)

		connectedComponents := CalculateConnectedComponents(*graph)

		t.Log("calculates the correct components")
		{
			fmt.Println(connectedComponents)
		}
	}
}