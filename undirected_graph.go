package algorithmsingo

import (
	"math/rand"
)

type Edge struct {
	V1     int
	V2     int
	weight int
}

type UndirectedGraph struct {
	Vertices      []int
	Edges         []Edge
	edgesByVertex map[int][]Edge
}

func (graph *UndirectedGraph) AddVertex(vertex int) {
	graph.Vertices = append(graph.Vertices, vertex)
	graph.edgesByVertex[vertex] = []Edge{}
}

func (graph *UndirectedGraph) AddEdge(vertex1 int, vertex2 int, weight int) {
	newEdge := Edge{
		V1:     vertex1,
		V2:     vertex2,
		weight: weight,
	}

	graph.Edges = append(graph.Edges, newEdge)
	graph.edgesByVertex[vertex1] = append(graph.edgesByVertex[vertex1], newEdge)
	graph.edgesByVertex[vertex2] = append(graph.edgesByVertex[vertex2], newEdge)
}

func NewUndriectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		edgesByVertex: make(map[int][]Edge),
	}
}

func GenerateRandomUndirectedGraph(numOfVertices int, numberOfEdges int, minWeight float64, maxWeight float64, isConnected bool) UndirectedGraph {
	graph := *NewUndriectedGraph()
	//	if numOfVertices < numberOfEdges - 1 {
	//		return graph, errors.New("numberOfEdges must be greater than numOfVertices")
	//	}

	vertices := make([]int, numOfVertices)

	for i := 0; i < numOfVertices; i++ {
		graph.AddVertex(i)
		vertices[i] = i
	}
	availableVerticesByVertex := map[int][]int{}
	for i, vertex := range vertices {
		t := append([]int(nil), vertices[:i]...)
		availableVerticesByVertex[vertex] = append(t, vertices[i+1:]...)
	}

	if isConnected {
		initialVerticesNeedingConnection := append([]int(nil), vertices...)

		for len(initialVerticesNeedingConnection) != 1 {
			v1 := initialVerticesNeedingConnection[0]
			indexOfV2 := rand.Intn(len(initialVerticesNeedingConnection)-1) + 1
			v2 := initialVerticesNeedingConnection[indexOfV2]
			graph.AddEdge(v1, v2, 0)
			availableVerticesByVertex[v1] = getSliceWithOutElement(availableVerticesByVertex[v1], v2)
			availableVerticesByVertex[v2] = getSliceWithOutElement(availableVerticesByVertex[v2], v1)
			initialVerticesNeedingConnection = initialVerticesNeedingConnection[1:]
		}
	}
	var remaningEdges int
	if isConnected {
		remaningEdges = numberOfEdges - (numOfVertices - 1)
	} else {
		remaningEdges = numberOfEdges
	}

	for i := 0; i < remaningEdges; i++ {
		v1ToSelectFrom := getVerticesNotConnectedToEverything(availableVerticesByVertex)
		indexOfV1 := rand.Intn(len(v1ToSelectFrom))
		v1 := v1ToSelectFrom[indexOfV1]
		v2ToSelectFrom := availableVerticesByVertex[v1]
		indexOfV2 := rand.Intn(len(v2ToSelectFrom))
		v2 := v2ToSelectFrom[indexOfV2]
		graph.AddEdge(v1, v2, 0)
		availableVerticesByVertex[v1] = getSliceWithOutElement(availableVerticesByVertex[v1], v2)
		availableVerticesByVertex[v2] = getSliceWithOutElement(availableVerticesByVertex[v2], v1)
	}

	return graph
}

func getVerticesNotConnectedToEverything(availableVerticesByVertex map[int][]int) []int {
	vertices := []int{}
	for vertex, availableVertices := range availableVerticesByVertex {
		if len(availableVertices) > 0 {
			vertices = append(vertices, vertex)
		}
	}
	return vertices
}

func getSliceWithOutElement(slice []int, item int) []int {
	indexOfElement := -1
	for i, currentItem := range slice {
		if currentItem == item {
			indexOfElement = i
			break
		}
	}
	if indexOfElement != -1 {
		return append(slice[:indexOfElement], slice[indexOfElement+1:]...)
	} else {
		return slice
	}
}
