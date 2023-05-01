package tasks

import "fmt"

/*
Task 30. Развернуть односвязный список
*/

type node struct {
	value int
	next  *node
}

func Task30() {
	start := initNodes(5)
	printNodes(start)
	printNodes(reverseNodes(start))
}

func reverseNodes(start *node) *node {
	if start == nil {
		return nil
	}

	if start.next == nil {
		return start
	}

	var elem, beforeElem *node
	if start.next.next == nil {
		elem = start.next
		start.next.next, start.next = start, nil
		return elem
	}

	beforeElem = start.next
	elem = beforeElem.next
	start.next = elem
	beforeElem.next = start

	for elem.next != nil {
		start.next = elem.next
		elem.next = beforeElem
		beforeElem = elem
		elem = start.next
	}

	elem.next = beforeElem
	start.next = nil
	return elem
}

func initNodes(n int) *node {
	if n <= 0 {
		return nil
	}

	start := node{
		value: 0,
	}

	elem := &start
	for i := 1; i < n; i++ {
		newElem := new(node)
		newElem.value = i

		elem.next = newElem
		elem = newElem
	}

	return &start
}

func printNodes(start *node) {
	if start == nil {
		fmt.Println("no nodes to show -> 0")
		return
	}

	fmt.Println()
	elem := start
	for {
		fmt.Print(elem.value, " ")
		if elem.next == nil {
			break
		}
		elem = elem.next
	}
}
