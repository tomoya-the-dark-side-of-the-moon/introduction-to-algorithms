package main

import "fmt"

type Element struct {
	key        int
	next, prev *Element
}

type Hashtable struct {
	bucketSize int
	table      []*Element
}

func (ht *Hashtable) put(z *Element) {
	hashVal := ht.calculateHashValue(z.key)

	z.next = ht.table[hashVal]
	if ht.table[hashVal] != nil {
		ht.table[hashVal].prev = z
	}
	ht.table[hashVal] = z
}

func (ht *Hashtable) contains(key int) *Element {
	hashVal := ht.calculateHashValue(key)

	for p := ht.table[hashVal]; p != nil; p = p.next {
		if p.key == key {
			return p
		}
	}

	return nil
}

func (ht *Hashtable) remove(z *Element) {
	hashVal := ht.calculateHashValue(z.key)

	if z.next != nil {
		z.next.prev = z.prev
	}
	if z.prev != nil {
		z.prev.next = z.next
	} else {
		ht.table[hashVal] = z.next
	}
	z = nil
}

func (ht *Hashtable) printHashValue() {
	for i := range ht.bucketSize {
		fmt.Printf("%d : ", i)
		for p := ht.table[i]; p != nil; p = p.next {
			fmt.Printf("%d", p.key)
			if p.next != nil {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
	}
}

func (ht *Hashtable) calculateHashValue(key int) int {
	return key % ht.bucketSize
}

func main() {
	table := make([]*Element, 13)
	ht := &Hashtable{bucketSize: 13, table: table}
	for i := 1; i <= 100; i++ {
		ht.put(&Element{key: i})
	}

	for {
		fmt.Print("1:put 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			ht.put(&Element{key: key})
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)

			z := ht.contains(key)
			if z == nil {
				fmt.Printf("%d is not found in the hash table.\n", key)
			} else {
				ht.remove(z)
			}
		case 3:
			ht.printHashValue()
		default:
			fmt.Println("invalid operation")
		}
	}
}
