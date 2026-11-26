package main

import (
	"fmt"
	"slices"
	"strings"
)

func breadthFirstSearch(graph *Graph, s *Vertex) {
	graph.InitializeVertices()

	s.color = gray
	s.distance = 0
	s.predecessor = nil

	queue := make([]*Vertex, 0)
	queue = append(queue, s)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		adjacencyList := graph.adjacencyLists[u.name]
		for _, adjacency := range adjacencyList {
			v := graph.vertices[adjacency]
			if v.color == white {
				fmt.Printf("%s is discovered.\n", v.name)

				v.color = gray
				v.distance = u.distance + 1
				v.predecessor = u

				queue = append(queue, v)
			}
		}

		fmt.Printf("%s is finished.\n", u.name)
		u.color = black
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
	for _, name := range "srvtuywxz" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "s", finish: "r"})
	edges = append(edges, &Edge{start: "s", finish: "v"})
	edges = append(edges, &Edge{start: "s", finish: "u"})
	edges = append(edges, &Edge{start: "r", finish: "s"})
	edges = append(edges, &Edge{start: "r", finish: "t"})
	edges = append(edges, &Edge{start: "r", finish: "w"})
	edges = append(edges, &Edge{start: "v", finish: "s"})
	edges = append(edges, &Edge{start: "v", finish: "y"})
	edges = append(edges, &Edge{start: "v", finish: "w"})
	edges = append(edges, &Edge{start: "t", finish: "r"})
	edges = append(edges, &Edge{start: "t", finish: "u"})
	edges = append(edges, &Edge{start: "u", finish: "s"})
	edges = append(edges, &Edge{start: "u", finish: "t"})
	edges = append(edges, &Edge{start: "u", finish: "y"})
	edges = append(edges, &Edge{start: "y", finish: "v"})
	edges = append(edges, &Edge{start: "y", finish: "u"})
	edges = append(edges, &Edge{start: "y", finish: "x"})
	edges = append(edges, &Edge{start: "w", finish: "r"})
	edges = append(edges, &Edge{start: "w", finish: "v"})
	edges = append(edges, &Edge{start: "w", finish: "x"})
	edges = append(edges, &Edge{start: "w", finish: "z"})
	edges = append(edges, &Edge{start: "x", finish: "y"})
	edges = append(edges, &Edge{start: "x", finish: "w"})
	edges = append(edges, &Edge{start: "x", finish: "z"})
	edges = append(edges, &Edge{start: "z", finish: "w"})
	edges = append(edges, &Edge{start: "z", finish: "x"})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- breadth first search ---")
	breadthFirstSearch(graph, vertices["s"])

	fmt.Println("--- result ---")
	for _, v := range vertices {
		fmt.Printf("%s (%d) : %s\n", v.name, v.distance, createPath(v))
	}
}
