package main

import "fmt"

type Element struct {
	data int
	next *Element
}

type Queue struct {
	size       int
	head, tail *Element
}

func (q *Queue) enqueue(z *Element) {
	q.size++

	if q.head == nil {
		q.head = z
	} else {
		q.tail.next = z
	}
	q.tail = z
}

func (q *Queue) dequeue() int {
	if q.size == 0 {
		return -1
	}

	q.size--

	dequeued := q.head
	result := dequeued.data
	q.head = dequeued.next
	dequeued = nil

	return result
}

func (q *Queue) printQueue() {
	fmt.Printf("size = %d : ", q.size)

	for p := q.head; p != nil; p = p.next {
		fmt.Print(p.data)
		if p.next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	q := &Queue{size: 0}

	for {
		fmt.Print("1:enqueue 2:dequeue 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)

			q.enqueue(&Element{data: data})
		case 2:
			result := q.dequeue()
			if result == -1 {
				fmt.Println("queue is empty")
			} else {
				fmt.Printf("result = %d\n", result)
			}
		case 3:
			q.printQueue()
		default:
			fmt.Println("invalid operation")
		}
	}
}
