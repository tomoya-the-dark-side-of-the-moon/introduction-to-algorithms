package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data   string
	rank   int
	parent *Node
}

func createNode(ch string) *Node {
	x := &Node{data: ch}
	x.parent = x

	return x
}

func unionSet(x, y *Node) {
	link(findSet(x), findSet(y))
}

func link(x, y *Node) {
	if x.rank > y.rank {
		y.parent = x
	} else {
		x.parent = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

func findSet(x *Node) *Node {
	if x.parent != x {
		x.parent = findSet(x.parent)
	}

	return x.parent
}

func printSets(nodes map[string]*Node) {
	groups := make(map[string][]string)
	for _, v := range nodes {
		s := findSet(v)
		groups[s.data] = make([]string, 0)
	}

	for _, v := range nodes {
		s := findSet(v)
		groups[s.data] = append(groups[s.data], v.data)
	}

	for k, v := range groups {
		fmt.Print(k, " : ", strings.Join(v, ", "), "\n")
	}
}

func main() {
	nodes := make(map[string]*Node)

	for _, ch := range "abcdefghij" {
		nodes[string(ch)] = createNode(string(ch))
	}

	fmt.Println("--- disjoint set after initializing ---")
	printSets(nodes)
	fmt.Println("--- merge b and d ---")
	unionSet(nodes["b"], nodes["d"])
	printSets(nodes)
	fmt.Println("--- merge e and f ---")
	unionSet(nodes["e"], nodes["f"])
	printSets(nodes)
	fmt.Println("--- merge a and c ---")
	unionSet(nodes["a"], nodes["c"])
	printSets(nodes)
	fmt.Println("--- merge h and i ---")
	unionSet(nodes["h"], nodes["i"])
	printSets(nodes)
	fmt.Println("--- merge a and b ---")
	unionSet(nodes["a"], nodes["b"])
	printSets(nodes)
	fmt.Println("--- merge f and g ---")
	unionSet(nodes["f"], nodes["g"])
	printSets(nodes)
	fmt.Println("--- merge b and c ---")
	unionSet(nodes["b"], nodes["c"])
	printSets(nodes)
}
