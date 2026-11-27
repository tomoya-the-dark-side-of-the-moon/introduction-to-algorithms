package main

import (
	"container/heap"
	"fmt"
	"math"
)

type VertexHeap []*Vertex

func (h VertexHeap) Len() int {
	return len(h)
}

func (h VertexHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h VertexHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *VertexHeap) Push(x any) {
	*h = append(*h, x.(*Vertex))
}

func (h VertexHeap) Contains(target *Vertex) bool {
	for _, v := range h {
		if v == target {
			return true
		}
	}

	return false
}

func (h *VertexHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func mstPrim(graph *Graph, r *Vertex) {
	vertices := graph.vertices
	for _, v := range vertices {
		v.weight = math.MaxInt64
		v.predecessor = nil
	}

	r.weight = 0

	vertexHeap := &VertexHeap{}
	heap.Init(vertexHeap)

	for _, v := range vertices {
		heap.Push(vertexHeap, v)
	}

	for vertexHeap.Len() != 0 {
		u := heap.Pop(vertexHeap).(*Vertex)
		fmt.Printf("target vertex: %s, weight = %d\n", u.name, u.weight)

		adjacencyList := graph.adjacencyLists[u.name]
		for _, direction := range adjacencyList {
			v := vertices[direction.direction]
			weight := direction.weight
			if vertexHeap.Contains(v) && v.weight > weight {
				fmt.Printf("Update vertex: %s, weight = %d -> %d\n", v.name, v.weight, weight)

				v.weight = weight
				v.predecessor = u

				heap.Init(vertexHeap)
			}
		}
	}
}

func main() {
	vertices := make(map[string]*Vertex)
	for _, name := range "abcdefghi" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "a", finish: "b", weight: 4})
	edges = append(edges, &Edge{start: "a", finish: "h", weight: 8})
	edges = append(edges, &Edge{start: "b", finish: "a", weight: 4})
	edges = append(edges, &Edge{start: "b", finish: "c", weight: 8})
	edges = append(edges, &Edge{start: "b", finish: "h", weight: 11})
	edges = append(edges, &Edge{start: "c", finish: "b", weight: 8})
	edges = append(edges, &Edge{start: "c", finish: "d", weight: 7})
	edges = append(edges, &Edge{start: "c", finish: "f", weight: 4})
	edges = append(edges, &Edge{start: "c", finish: "i", weight: 2})
	edges = append(edges, &Edge{start: "d", finish: "c", weight: 7})
	edges = append(edges, &Edge{start: "d", finish: "e", weight: 9})
	edges = append(edges, &Edge{start: "d", finish: "f", weight: 14})
	edges = append(edges, &Edge{start: "e", finish: "d", weight: 9})
	edges = append(edges, &Edge{start: "e", finish: "f", weight: 10})
	edges = append(edges, &Edge{start: "f", finish: "c", weight: 4})
	edges = append(edges, &Edge{start: "f", finish: "d", weight: 14})
	edges = append(edges, &Edge{start: "f", finish: "e", weight: 10})
	edges = append(edges, &Edge{start: "f", finish: "g", weight: 2})
	edges = append(edges, &Edge{start: "g", finish: "f", weight: 2})
	edges = append(edges, &Edge{start: "g", finish: "h", weight: 1})
	edges = append(edges, &Edge{start: "g", finish: "i", weight: 6})
	edges = append(edges, &Edge{start: "h", finish: "a", weight: 8})
	edges = append(edges, &Edge{start: "h", finish: "b", weight: 11})
	edges = append(edges, &Edge{start: "h", finish: "g", weight: 1})
	edges = append(edges, &Edge{start: "h", finish: "i", weight: 7})
	edges = append(edges, &Edge{start: "i", finish: "c", weight: 2})
	edges = append(edges, &Edge{start: "i", finish: "g", weight: 6})
	edges = append(edges, &Edge{start: "i", finish: "h", weight: 7})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- Prim's algorithm ---")
	mstPrim(graph, vertices["a"])

	fmt.Println("--- result ---")
	total := 0
	for _, v := range vertices {
		fmt.Printf("name = %s, weight = %d", v.name, v.weight)
		if v.predecessor != nil {
			fmt.Printf(", predecessor = %s", v.predecessor.name)
		}
		fmt.Println()

		total += v.weight
	}
	fmt.Printf("total = %d\n", total)
}
