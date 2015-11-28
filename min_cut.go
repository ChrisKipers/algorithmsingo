package algorithmsingo

import (
	"math/rand"
	"sync"
	"time"
)

type VertexGroup struct {
	Vertices []int
}

type MinCut struct {
	Groups        []VertexGroup
	NumberOfEdges int
}

type groupEdge struct {
	g1 *VertexGroup
	g2 *VertexGroup
}

func CalculateMinCut(graph Graph, numberOfCuts int) MinCut {
	var bestMinCut *MinCut = nil
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numOfTrials := getNumberOfTrials(len(graph.Vertices))
	for i := 0; i < numOfTrials; i++ {
		minCut := calculateMinCut(graph, numberOfCuts, myRand)
		if bestMinCut == nil || minCut.NumberOfEdges < bestMinCut.NumberOfEdges {
			bestMinCut = &minCut
		}
	}
	return *bestMinCut
}

func CalculateMinCutPar(graph Graph, numberOfCuts int) MinCut {
	numVertices := len(graph.Vertices)
	numOfTrials := getNumberOfTrials(len(graph.Vertices))
	minCuts := make(chan MinCut, numVertices)
	wg := &sync.WaitGroup{}
	for i := 0; i < numOfTrials; i++ {
		wg.Add(1)
		go func() {
			myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
			defer wg.Done()
			min := calculateMinCut(graph, numberOfCuts, myRand)
			minCuts <- min
		}()
	}
	go watchWg(wg, minCuts)
	bestMinCut := <-minCuts
	for minCut := range minCuts {
		if minCut.NumberOfEdges < bestMinCut.NumberOfEdges {
			bestMinCut = minCut
		}
	}
	return bestMinCut
}

func getNumberOfTrials(numVertices int) int {
	return numVertices * numVertices
}

func watchWg(wg *sync.WaitGroup, channel chan MinCut) {
	wg.Wait()
	close(channel)
}

func calculateMinCut(graph Graph, numberOfCuts int, myRand *rand.Rand) MinCut {
	vertexToGroup := make(map[int]*VertexGroup)
	groups := []*VertexGroup{}

	for _, vertex := range graph.Vertices {
		group := VertexGroup{
			Vertices: []int{vertex},
		}
		vertexToGroup[vertex] = &group
		groups = append(groups, &group)
	}

	groupEdges := make([]*groupEdge, len(graph.Edges))

	for i, edge := range graph.Edges {
		v1Group := vertexToGroup[edge.V1]
		v2Group := vertexToGroup[edge.V2]
		gEdge := groupEdge{
			g1: v1Group,
			g2: v2Group,
		}
		groupEdges[i] = &gEdge
	}
	numVertices := len(graph.Vertices)
	for i := 0; i < numVertices-numberOfCuts; i++ {
		// Randomly select edge to remove
		selectedEdgeIndex := myRand.Intn(len(groupEdges))
		selectedEdge := groupEdges[selectedEdgeIndex]

		// Merge groups related to edge
		g1 := selectedEdge.g1
		g2 := selectedEdge.g2
		newGroup := g1.merge(*g2)

		// Filter out groups that were merged
		newGroups := make([]*VertexGroup, len(groups)-1)
		newGroupIndex := 0
		for _, group := range groups {
			if group != g1 && group != g2 {
				newGroups[newGroupIndex] = group
				newGroupIndex++
			}
		}

		newGroups[len(newGroups)-1] = &newGroup
		groups = newGroups

		// Update all edges
		for _, edge := range groupEdges {
			if edge.g1 == g1 || edge.g1 == g2 {
				edge.g1 = &newGroup
			}
			if edge.g2 == g1 || edge.g2 == g2 {
				edge.g2 = &newGroup
			}
		}

		newGroupEdges := []*groupEdge{}
		// Filter out edges
		for _, edge := range groupEdges {
			if edge.g1 != edge.g2 {
				newGroupEdges = append(newGroupEdges, edge)
			}
		}

		groupEdges = newGroupEdges
	}
	formattedGroups := make([]VertexGroup, len(groups))
	for i, group := range groups {
		formattedGroups[i] = *group
	}
	return MinCut{
		Groups:        formattedGroups,
		NumberOfEdges: len(groupEdges),
	}
}

func (group1 VertexGroup) merge(group2 VertexGroup) VertexGroup {
	vertices := append(group1.Vertices, group2.Vertices...)
	return VertexGroup{
		Vertices: vertices,
	}
}
