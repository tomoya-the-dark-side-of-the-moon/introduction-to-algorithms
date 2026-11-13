package main

import "fmt"

type Element struct {
	data       int
	next, prev *Element
}

type LinkedList struct {
	size       int
	head, tail *Element
}

func (ll *LinkedList) add(z *Element) {
	ll.size++

	if ll.head == nil {
		ll.head = z
	} else {
		ll.tail.next = z
	}
	z.prev = ll.tail
	ll.tail = z
}

func (ll *LinkedList) contains(data int) *Element {
	for p := ll.head; p != nil; p = p.next {
		if p.data == data {
			return p
		}
	}

	return nil
}

func (ll *LinkedList) remove(z *Element) {
	ll.size--

	if z.next != nil {
		z.next.prev = z.prev
	} else {
		ll.tail = z.prev
	}
	if z.prev != nil {
		z.prev.next = z.next
	} else {
		ll.head = z.next
	}
	z = nil
}

func (ll *LinkedList) printList() {
	fmt.Printf("size = %d\n", ll.size)

	fmt.Println("--- from head to tail ---")
	for p := ll.head; p != nil; p = p.next {
		fmt.Printf("%d", p.data)
		if p.next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
	}

	fmt.Println("--- from tail to head ---")
	for p := ll.tail; p != nil; p = p.prev {
		fmt.Printf("%d", p.data)
		if p.prev != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	ll := &LinkedList{size: 0}

	for {
		fmt.Print("1:add 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)

			ll.add(&Element{data: data})
		case 2:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)

			z := ll.contains(data)
			if z == nil {
				fmt.Printf("%d is not found in the list.\n", data)
			} else {
				ll.remove(z)
			}
		case 3:
			ll.printList()
		default:
			fmt.Println("invalid operation")
		}
	}
}
