//implement dynamic array

package dyanmicarray

import "errors"

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
