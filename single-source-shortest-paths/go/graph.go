package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	name        string
	distance    int
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

func InitializeSingleSource(graph *Graph, s *Vertex) {
	for _, v := range graph.vertices {
		v.distance = math.MaxInt64
		v.predecessor = nil
	}

	s.distance = 0
}

func Relax(u, v *Vertex, weight int) {
	if v.distance > u.distance+weight {
		fmt.Printf("Relax (u, v) = (%s, %s) / before = %d, after = %d\n", u.name, v.name, v.distance, u.distance+weight)

		v.distance = u.distance + weight
		v.predecessor = u
	}
}
