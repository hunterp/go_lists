package main

import "errors"

type MyLinkedList struct {
	currSize int
	head *Node
	tail *Node
}

type Node struct {
	value int
	prev *Node
	next *Node
}

func (list *MyLinkedList) size() int {
	return list.currSize
}

func (list *MyLinkedList) isEmpty() bool {
	return list.currSize == 0;
}

func (list *MyLinkedList) contains(val int) bool {
	curr := list.head
	for curr != nil {
		if curr.value == val {
			return true
		}
		curr = curr.next
	}
	return false
}

func (list *MyLinkedList) add(val int) {
	newNode := Node{value: val}
	if list.tail == nil {
		list.head = &newNode
		list.tail = &newNode
	} else {
		list.tail.next = &newNode
		newNode.prev = list.tail
		list.tail = &newNode
	}
	list.currSize ++
}

func (list *MyLinkedList) remove(val int) bool{
	curr := list.head
	for curr != nil {
		if curr.value == val {
			if curr.prev == nil {
				list.head = curr.next
			} else {
				curr.prev.next = curr.next
			}
			if curr.next == nil {
				list.tail = curr.prev
			} else {
				curr.next.prev = curr.prev
			}
			list.currSize--
			return true
		}
		curr = curr.next
	}
	return false
}

func (list *MyLinkedList) containsAll(collection []int) bool {
	for _,v := range collection {
		if list.contains(v) == false {
			return false
		}
	}
	return true
}

func (list *MyLinkedList) addAll(collection []int) bool {
	for _,v := range collection {
		list.add(v)
	}
	return true
}

func (list *MyLinkedList) clear() {
	list.head = nil
	list.tail = nil
	list.currSize = 0
}

func (list *MyLinkedList) get(idx int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}
	curr := list.head
	for p := 0; p < idx; p++ {
		curr = curr.next
	}
	return curr.value, nil
}

func (list *MyLinkedList) set(idx int, value int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}
	curr := list.head
	for p := 0; p < idx; p++ {
		curr = curr.next
	}
	oldVal := curr.value
	curr.value = value
	return oldVal, nil
}
// Inserts val so it is the element at index idx
func (list *MyLinkedList) addAtIndex(idx int, val int) (int, error) {
	if idx > list.currSize {
		return 0, errors.New("index out of bounds")
	}

	newNode := Node{value: val}
	curr := list.head
	for pos := 0; pos < idx; pos ++ {
		curr = curr.next
	}

	//curr being null means that EITHER the list is currently empty OR I'm inserting at the end of the list
	if curr == nil {
		if list.currSize > 0 {
			newNode.prev = list.tail
			list.tail.next = &newNode
		}
		list.tail = &newNode //Either way this is the new tail
	} else {
		newNode.prev = curr.prev
		newNode.next = curr
		if newNode.prev != nil {
			newNode.prev.next = &newNode
		}
	}

	if idx == 0 {
		list.head = &newNode
	}

	list.currSize ++
	return val, nil
}

//removes the value at index idx
func (list *MyLinkedList) removeAtIndex(idx int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}

	curr := list.head
	for pos := 0; pos < idx ; pos ++ {
		curr = curr.next
	}
	//curr is now the element that we want to remove.
	if curr.prev == nil {
		list.head = curr.next
	} else {
		curr.prev.next = curr.next
	}
	if curr.next == nil {
		list.tail = curr.prev
	} else {
		curr.next.prev = curr.prev
	}
	list.currSize--
	return curr.value, nil
}

func (list *MyLinkedList) indexOf(val int) int {
	curr := list.head
	for i:= 0; curr != nil; i++ {
		if (curr.value == val) {
			return i
		}
		curr = curr.next
	}
	return -1
}