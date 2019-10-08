package main

import "fmt"
import "crypto/rand"
import "math/big"

func main() {
	var list List
	list = &MyLinkedList{}
	testList(list)

	list = &MyArrayList{}
	testList(list)
}

func testList(list List) {
	//Add 100 random elements to the list.
	for i := 0; i < 100; i++ {
		p := new(big.Int).SetInt64(100)
		num, _ := rand.Int(rand.Reader, p)
		list.add(int(num.Int64()))
	}

	assertEqual(100, list.size());
	assertEqual(false, list.isEmpty());

	// Clear the list out
	list.clear();
	assertEqual(0, list.size());

	// Add in some specific elements
	list.add(1); // [1]
	list.add(2); // [1,2]
	list.addAtIndex(0, 3); // [3,1,2]
	val, _ := list.get(1)
	assertEqual(1, val);
	assertEqual(true, list.contains(3));
	assertEqual(false, list.contains(7));

	// Test removing an element
	removed, _ := list.removeAtIndex(0); //remove element at index zero.
	assertEqual(removed, 3); //3 should have been removed
	assertEqual(2, list.size());

	// Add again after removal.
	list.add(4); //[1,2,4]
	assertEqual(2, list.indexOf(4));
}

func assertEqual(a, b interface {}) {
	if (a == b) == false {
		fmt.Printf("FAIL! %v != %v\n", a, b)
	} else {
		fmt.Printf("PASS, %v == %v\n", a, b)
	}
}


type List interface {
	size() int
	get(idx int) (int, error)
	contains(val int) bool
	add(val int)
	remove(val int) bool
	containsAll(collection []int) bool
	addAll(collection []int) bool
	clear()
	set(idx int, val int) (int, error)
	addAtIndex(idx int, val int) (int, error)
	removeAtIndex(idx int) (int, error)
	indexOf(val int) (int)
	isEmpty() bool
}

func printList(list List) {
	for i := 0; i< list.size(); i++ {
		val, _ := list.get(i)
		fmt.Printf("%v, ", val)
	}
	fmt.Println()
}
