package main

import "errors"

type MyArrayList struct {
	array []int
	currSize int
}


func (list *MyArrayList) size() int {
	return list.currSize
}

func (list *MyArrayList) isEmpty() bool {
	return list.currSize == 0;
}

func (list *MyArrayList) contains(val int) bool {
	for i := 0; i< list.currSize; i++ {
		if list.array[i] == val {
			return true
		}
	}
	return false
}

func extend(list *MyArrayList) {
	newArray := make([] int, list.currSize+10)
	for i := 0; i < list.currSize; i++ {
		newArray[i] = list.array[i]
	}
	list.array = newArray
}

func (list *MyArrayList) add(val int) {
	//If there is not enough space in the array then we'll extend it, copy the values and then insert
	if list.currSize == len(list.array) {
		extend(list)
	}
	list.array[list.currSize] = val
	list.currSize++
}

func (list *MyArrayList) remove(val int) bool{
	found := false
	for i:= 0; i< list.currSize; i++  {
		if list.array[i] == val {
			found = true
		} else if found {
			list.array[i-1] = list.array[i] //Shift all by 1 left. AFTER you found the element
		}
	}
	if found {
		list.currSize--
	}
	return found
}

func (list *MyArrayList) containsAll(collection []int) bool {
	for _,v := range collection {
		if list.contains(v) == false {
			return false
		}
	}
	return true
}

func (list *MyArrayList) addAll(collection []int) bool {
	for _,v := range collection {
		list.add(v)
	}
	return true
}

func (list *MyArrayList) clear() {
	list.currSize = 0
}

func (list *MyArrayList) get(idx int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}
	return list.array[idx], nil
}

func (list *MyArrayList) set(idx int, value int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}
	list.array[idx] = value
	return value, nil
}
// Inserts val so it is the element at index idx
func (list *MyArrayList) addAtIndex(idx int, val int) (int, error) {
	if idx > list.currSize {
		return 0, errors.New("index out of bounds")
	}

	if list.currSize ==  len(list.array) {
		extend(list)
	}
	for i := list.currSize; i > idx ; i-- {
		list.array[i] = list.array[i-1]
	}
	list.array[idx] = val
	list.currSize++
	return val, nil
}

//removes the value at index idx
func (list *MyArrayList) removeAtIndex(idx int) (int, error) {
	if idx >= list.currSize {
		return 0, errors.New("index out of bounds")
	}
	oldVal := list.array[idx]
	for i := idx; i< (list.currSize -1); i++ {
		list.array[i] = list.array[i+1]
	}
	list.currSize--
	return oldVal, nil
}

func (list *MyArrayList) indexOf(val int) int {
	for i := 0; i< list.currSize ; i++ {
		if list.array[i] == val {
			return i
		}
	}
	return -1
}