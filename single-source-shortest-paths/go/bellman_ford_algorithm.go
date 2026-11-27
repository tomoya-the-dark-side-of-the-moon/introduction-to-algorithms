package main

import (
	"fmt"
	"slices"
	"strings"
)

func bellmanFord(graph *Graph, s *Vertex) bool {
	InitializeSingleSource(graph, s)

	vertices := graph.vertices
	edges := graph.edges

	for i := 0; i < len(vertices)-1; i++ {
		for _, edge := range edges {
			u, v, weight := vertices[edge.start], vertices[edge.finish], edge.weight
			Relax(u, v, weight)
		}
	}

	for _, edge := range edges {
		u, v, weight := vertices[edge.start], vertices[edge.finish], edge.weight
		if v.distance > u.distance+weight {
			return false
		}
	}

	return true
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
	edges = append(edges, &Edge{start: "s", finish: "t", weight: 6})
	edges = append(edges, &Edge{start: "s", finish: "y", weight: 7})
	edges = append(edges, &Edge{start: "t", finish: "y", weight: 8})
	edges = append(edges, &Edge{start: "t", finish: "x", weight: 5})
	edges = append(edges, &Edge{start: "t", finish: "z", weight: -4})
	edges = append(edges, &Edge{start: "y", finish: "x", weight: -3})
	edges = append(edges, &Edge{start: "y", finish: "z", weight: 9})
	edges = append(edges, &Edge{start: "x", finish: "t", weight: -2})
	edges = append(edges, &Edge{start: "z", finish: "s", weight: 2})
	edges = append(edges, &Edge{start: "z", finish: "x", weight: 7})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- Bellman-Ford algorithm ---")
	hasNoNegativeWeightCycle := bellmanFord(graph, vertices["s"])

	fmt.Println("--- result ---")
	if !hasNoNegativeWeightCycle {
		fmt.Println("Graph has a negative weight cycle.")
	} else {
		for _, v := range vertices {
			fmt.Printf("name = %s / weight = %d (%s)\n", v.name, v.distance, createPath(v))
		}
	}
}
