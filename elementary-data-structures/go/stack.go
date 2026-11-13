package main

import "fmt"

type Element struct {
	data int
	next *Element
}

type Stack struct {
	size int
	head *Element
}

func (s *Stack) push(z *Element) {
	s.size++

	z.next = s.head
	s.head = z
}

func (s *Stack) pop() int {
	if s.size == 0 {
		return -1
	}

	s.size--

	popped := s.head
	result := popped.data
	s.head = popped.next
	popped = nil

	return result
}

func (s *Stack) printStack() {
	fmt.Printf("size = %d : ", s.size)

	for p := s.head; p != nil; p = p.next {
		fmt.Print(p.data)
		if p.next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	s := &Stack{size: 0, head: nil}

	for {
		fmt.Print("1:push 2:pop 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)

			s.push(&Element{data: data})
		case 2:
			result := s.pop()
			if result == -1 {
				fmt.Println("stack is empty")
			} else {
				fmt.Printf("result = %d\n", result)
			}
		case 3:
			s.printStack()
		default:
			fmt.Println("invalid operation")
		}
	}
}
