//implement dynamic array

package dyanmicarray

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	capacity int
	size     int
	elements []interface{}
}

func New(capacity int, elements ...interface{}) (error, *DynamicArray) {
	if capacity < 1 {
		err := errors.New("capacity must be greater than 0")
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
		dynamicArray.capacity = dynamicArray.capacity * 2

		newArray := make([]interface{}, dynamicArray.capacity)
		for idx, val := range dynamicArray.elements {
			newArray[idx] = val
		}

		dynamicArray.elements = newArray
	}
	fmt.Println(dynamicArray.capacity, dynamicArray.size)

	dynamicArray.elements[dynamicArray.size] = element
	dynamicArray.size += 1

	return dynamicArray.size - 1
}
