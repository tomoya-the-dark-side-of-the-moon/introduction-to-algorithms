package main

import (
	"fmt"
)

type Vertex struct {
	name        string
	weight      int
	predecessor *Vertex
}

type Edge struct {
	start  string
	finish string
	weight int
}

type Direction struct {
	direction string
	weight    int
}

type Graph struct {
	vertices       map[string]*Vertex
	edges          []*Edge
	adjacencyLists map[string][]Direction
}

func (graph *Graph) InitializeGraph() {
	vertices := graph.vertices
	edges := graph.edges
	graph.adjacencyLists = make(map[string][]Direction)
	for k := range vertices {
		graph.adjacencyLists[k] = make([]Direction, 0)
	}

	for _, edge := range edges {
		graph.adjacencyLists[edge.start] = append(graph.adjacencyLists[edge.start], Direction{direction: edge.finish, weight: edge.weight})
	}
}

func (graph *Graph) PrintGraph() {
	vertices := graph.vertices
	for k := range vertices {
		fmt.Print(k, " : ")
		adjacencyList := graph.adjacencyLists[k]
		for i, adjacency := range adjacencyList {
			fmt.Printf("%s(%d)", adjacency.direction, adjacency.weight)
			if i < len(adjacencyList)-1 {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
	}
}
