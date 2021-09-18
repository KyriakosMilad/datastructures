//implement dynamic array

package dyanmicarray

import (
	"errors"
)

type DynamicArray struct {
	capacity int
	size     int
	elements []interface{}
}

func New(capacity int, elements ...interface{}) (error, *DynamicArray) {
	if capacity < 0 {
		err := errors.New("capacity must be greater than or equal to 0")
		return err, nil
	}

	if capacity < len(elements) {
		err := errors.New("capacity must be greater than or equal to the number of elements you want to add")
		return err, nil
	}

	dynamicArray := DynamicArray{
		capacity: capacity,
		size:     len(elements),
		elements: make([]interface{}, capacity),
	}

	for idx, val := range elements {
		dynamicArray.elements[idx] = val
	}

	return nil, &dynamicArray
}

func (dynamicArray *DynamicArray) Append(element interface{}) int {
	if dynamicArray.capacity <= dynamicArray.size+1 {
		if dynamicArray.capacity == 0 {
			dynamicArray.capacity = 1
		} else {
			dynamicArray.capacity = dynamicArray.capacity * 2
		}

		newArray := make([]interface{}, dynamicArray.capacity)
		for idx, val := range dynamicArray.elements {
			newArray[idx] = val
		}

		dynamicArray.elements = newArray
	}

	dynamicArray.elements[dynamicArray.size] = element
	dynamicArray.size += 1

	return dynamicArray.size - 1
}

func (dynamicArray *DynamicArray) InsertAt(index int, element interface{}) error {
	if index > (dynamicArray.size) {
		return errors.New("index you want to add the element at is out of range, if you want to add the element at the end of the array try Append() method")
	}

	if index < 0 {
		return errors.New("index you want to add element at must be greater than or equal 0")
	}

	if dynamicArray.capacity <= dynamicArray.size+1 {
		if dynamicArray.capacity == 0 {
			dynamicArray.capacity = 1
		} else {
			dynamicArray.capacity = dynamicArray.capacity * 2
		}

		newArray := make([]interface{}, dynamicArray.capacity)
		for idx, val := range dynamicArray.elements {
			newArray[idx] = val
		}

		dynamicArray.elements = newArray
	}

	tempElement := element
	for i := index; i < dynamicArray.size; i++ {
		tempVal := dynamicArray.elements[i]
		dynamicArray.elements[i] = tempElement
		tempElement = tempVal
	}
	dynamicArray.elements[dynamicArray.size] = tempElement

	dynamicArray.size += 1

	return nil
}
