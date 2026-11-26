package main

import (
	"fmt"
	"math"
)

type Color int

const (
	white Color = iota
	gray
	black
)

type Vertex struct {
	name        string
	distance    int
	predecessor *Vertex
	color       Color
}

type Edge struct {
	start  string
	finish string
}

type Graph struct {
	vertices       map[string]*Vertex
	edges          []*Edge
	adjacencyLists map[string][]string
}

func (graph *Graph) InitializeVertices() {
	vertices := graph.vertices
	for _, v := range vertices {
		v.color = white
		v.distance = math.MaxInt64
		v.predecessor = nil
	}
}

func (graph *Graph) InitializeGraph() {
	vertices := graph.vertices
	edges := graph.edges
	graph.adjacencyLists = make(map[string][]string)
	for k := range vertices {
		graph.adjacencyLists[k] = make([]string, 0)
	}

	for _, edge := range edges {
		graph.adjacencyLists[edge.start] = append(graph.adjacencyLists[edge.start], edge.finish)
	}
}

func (graph *Graph) PrintGraph() {
	vertices := graph.vertices
	for k := range vertices {
		fmt.Print(k, " : ")
		adjacencyList := graph.adjacencyLists[k]
		for i, adjacency := range adjacencyList {
			fmt.Print(adjacency)
			if i < len(adjacencyList)-1 {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
	}
}
