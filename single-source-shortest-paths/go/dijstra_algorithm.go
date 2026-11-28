package main

import (
	"container/heap"
	"fmt"
	"slices"
	"strings"
)

type VertexHeap []*Vertex

func (h VertexHeap) Len() int {
	return len(h)
}

func (h VertexHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h VertexHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *VertexHeap) Push(x any) {
	item := x.(*Vertex)
	*h = append(*h, item)
}

func (h *VertexHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func dijkstra(graph *Graph, s *Vertex) {
	InitializeSingleSource(graph, s)

	vertexHeap := &VertexHeap{}
	for _, v := range graph.vertices {
		vertexHeap.Push(v)
	}

	heap.Init(vertexHeap)

	for vertexHeap.Len() != 0 {
		u := heap.Pop(vertexHeap).(*Vertex)
		adjacencyList := graph.adjacencyLists[u.name]
		for _, direction := range adjacencyList {
			v := graph.vertices[direction.direction]
			weight := direction.weight
			Relax(u, v, weight)

			heap.Init(vertexHeap)
		}
	}
}

func createPath(v *Vertex) string {
	path := make([]string, 0)

	for v != nil {
		path = append(path, v.name)

		v = v.predecessor
	}

	slices.Reverse(path)

	return strings.Join(path, " -> ")
}

func main() {
	vertices := make(map[string]*Vertex)
	for _, name := range "styxz" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "s", finish: "t", weight: 10})
	edges = append(edges, &Edge{start: "s", finish: "y", weight: 5})
	edges = append(edges, &Edge{start: "t", finish: "y", weight: 2})
	edges = append(edges, &Edge{start: "t", finish: "x", weight: 1})
	edges = append(edges, &Edge{start: "y", finish: "t", weight: 3})
	edges = append(edges, &Edge{start: "y", finish: "x", weight: 9})
	edges = append(edges, &Edge{start: "y", finish: "z", weight: 2})
	edges = append(edges, &Edge{start: "x", finish: "z", weight: 4})
	edges = append(edges, &Edge{start: "z", finish: "s", weight: 7})
	edges = append(edges, &Edge{start: "z", finish: "x", weight: 6})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- Dijkstra's algorithm ---")
	dijkstra(graph, vertices["s"])

	fmt.Println("--- result ---")
	for _, v := range vertices {
		fmt.Printf("name = %s / weight = %d (%s)\n", v.name, v.distance, createPath(v))
	}
}
