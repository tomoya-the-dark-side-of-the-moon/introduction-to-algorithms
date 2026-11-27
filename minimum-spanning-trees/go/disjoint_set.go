package main

type Node struct {
	data   string
	rank   int
	parent *Node
}

func CreateNode(ch string) *Node {
	x := &Node{data: ch}
	x.parent = x

	return x
}

func UnionSet(x, y *Node) {
	link(FindSet(x), FindSet(y))
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

func FindSet(x *Node) *Node {
	if x.parent != x {
		x.parent = FindSet(x.parent)
	}

	return x.parent
}
