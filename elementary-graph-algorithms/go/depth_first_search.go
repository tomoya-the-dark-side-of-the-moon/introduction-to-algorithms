package main

import (
	"fmt"
	"slices"
	"strings"
)

var time = 0

func depthFirstSearch(graph *Graph) {
	for _, u := range graph.vertices {
		u.color = white
		u.predecessor = nil
	}

	for _, u := range graph.vertices {
		if u.color == white {
			depthFirstSearchInternal(graph, u)
		}
	}
}

func depthFirstSearchInternal(graph *Graph, u *Vertex) {
	fmt.Printf("%s is discovered.\n", u.name)

	time++
	u.discoveryTime = time
	u.color = gray

	adjacencyList := graph.adjacencyLists[u.name]
	for _, adjacency := range adjacencyList {
		v := graph.vertices[adjacency]
		if v.color == white {
			v.predecessor = u
			depthFirstSearchInternal(graph, v)
		}
	}

	fmt.Printf("%s is finished.\n", u.name)

	time++
	u.finishTime = time
	u.color = black
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
	for _, name := range "uvxywz" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "u", finish: "v"})
	edges = append(edges, &Edge{start: "u", finish: "x"})
	edges = append(edges, &Edge{start: "v", finish: "y"})
	edges = append(edges, &Edge{start: "x", finish: "v"})
	edges = append(edges, &Edge{start: "y", finish: "x"})
	edges = append(edges, &Edge{start: "w", finish: "y"})
	edges = append(edges, &Edge{start: "w", finish: "z"})
	edges = append(edges, &Edge{start: "z", finish: "z"})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- depth first search ---")
	depthFirstSearch(graph)

	fmt.Println("--- result ---")
	for _, v := range vertices {
		fmt.Printf("%s (%d / %d) : %s\n", v.name, v.discoveryTime, v.finishTime, createPath(v))
	}
}
