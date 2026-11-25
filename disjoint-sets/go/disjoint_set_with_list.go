package main

import "fmt"

type Element struct {
	data string
	next *Element
}

type LinkedList struct {
	head *Element
	tail *Element
}

func (ll *LinkedList) append(z *Element) {
	if ll.head == nil {
		ll.head = z
	} else {
		ll.tail.next = z
	}
	ll.tail = z
}

func (ll *LinkedList) contains(data string) bool {
	for p := ll.head; p != nil; p = p.next {
		if p.data == data {
			return true
		}
	}

	return false
}

func unionSet(sets map[string]*LinkedList, x, y string) {
	headOfX := findSet(sets, x)
	headOfY := findSet(sets, y)

	if headOfX == nil || headOfY == nil {
		return
	}

	if headOfX.data == headOfY.data {
		return
	}

	llOfX := sets[headOfX.data]
	llOfY := sets[headOfY.data]
	if headOfX.data < headOfY.data {
		llOfX.tail.next = llOfY.head
		llOfX.tail = llOfY.tail

		delete(sets, headOfY.data)
	} else {
		llOfY.tail.next = llOfX.head
		llOfY.tail = llOfX.tail

		delete(sets, headOfX.data)
	}
}

func findSet(sets map[string]*LinkedList, data string) *Element {
	for _, ll := range sets {
		if ll.contains(data) {
			return ll.head
		}
	}

	return nil
}

func printSet(sets map[string]*LinkedList) {
	for k, ll := range sets {
		fmt.Print(k, " : ")
		for p := ll.head; p != nil; p = p.next {
			fmt.Print(p.data)
			if p.next != nil {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()
	}
}

func main() {
	sets := make(map[string]*LinkedList)

	for _, ch := range "abcdefghij" {
		e := &Element{data: string(ch)}
		ll := &LinkedList{}
		ll.append(e)
		sets[string(ch)] = ll
	}

	fmt.Println("--- disjoint set after initializing ---")
	printSet(sets)
	fmt.Println("--- merge b and d ---")
	unionSet(sets, "b", "d")
	printSet(sets)
	fmt.Println("--- merge e and f ---")
	unionSet(sets, "e", "f")
	printSet(sets)
	fmt.Println("--- merge a and c ---")
	unionSet(sets, "a", "c")
	printSet(sets)
	fmt.Println("--- merge h and i ---")
	unionSet(sets, "h", "i")
	printSet(sets)
	fmt.Println("--- merge a and b ---")
	unionSet(sets, "a", "b")
	printSet(sets)
	fmt.Println("--- merge f and g ---")
	unionSet(sets, "f", "g")
	printSet(sets)
	fmt.Println("--- merge b and c ---")
	unionSet(sets, "b", "c")
	printSet(sets)
}
