package main

import (
	"cmp"
	"fmt"
	"slices"
)

func mstKruskal(graph *Graph) []*Edge {
	nodes := make(map[string]*Node)
	for k := range graph.vertices {
		nodes[k] = CreateNode(k)
	}

	edges := graph.edges
	slices.SortFunc(edges, func(x, y *Edge) int { return cmp.Compare(x.weight, y.weight) })

	results := make([]*Edge, 0)
	for _, edge := range edges {
		u, v, weight := edge.start, edge.finish, edge.weight
		if FindSet(nodes[u]) != FindSet(nodes[v]) {
			fmt.Printf("add edge (%s, %s) (%d)\n", u, v, weight)

			results = append(results, edge)
			UnionSet(nodes[u], nodes[v])
		}
	}

	return results
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

	fmt.Println("--- Kruskal's algorithm ---")
	results := mstKruskal(graph)

	fmt.Println("--- result ---")
	total := 0
	for _, edge := range results {
		fmt.Printf("start = %s, finish = %s, weight = %d\n", edge.start, edge.finish, edge.weight)
		total += edge.weight
	}
	fmt.Printf("total = %d\n", total)
}
