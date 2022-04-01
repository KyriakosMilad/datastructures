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
	if dynamicArray.capacity < (dynamicArray.size + 1) {
		if dynamicArray.capacity == 0 {
			dynamicArray.capacity = 1
		} else if dynamicArray.capacity == dynamicArray.size {
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

	if dynamicArray.capacity < (dynamicArray.size + 1) {
		if dynamicArray.capacity == 0 {
			dynamicArray.capacity = 1
		} else if dynamicArray.capacity == dynamicArray.size {
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

func (dynamicArray *DynamicArray) RemoveAt(index int) error {
	if dynamicArray.size == 0 {
		return errors.New("cannot remove from empty array")
	}

	if index > (dynamicArray.size-1) || index < 0 {
		return errors.New("index out of range")
	}

	newArray := make([]interface{}, dynamicArray.size-1)

	counter := 0
	for _, val := range dynamicArray.elements {
		if counter == index {
			continue
		}
		newArray[counter] = val
		counter += 1
	}

	dynamicArray.elements = newArray
	dynamicArray.size -= 1
	dynamicArray.capacity = dynamicArray.size

	return nil
}

func (dynamicArray *DynamicArray) RemoveAllWhere(element interface{}) error {
	if dynamicArray.size == 0 {
		return errors.New("cannot remove from empty array")
	}

	tempArray := make([]interface{}, dynamicArray.capacity)
	capacityCounter := dynamicArray.capacity
	counter := 0
	for _, val := range dynamicArray.elements {
		if val == element || val == nil {
			capacityCounter -= 1
			continue
		}
		tempArray[counter] = val
		counter += 1
	}

	newArray := make([]interface{}, capacityCounter)
	counter = 0
	for _, val := range tempArray {
		if val == nil {
			continue
		}
		newArray[counter] = val
		counter += 1
	}

	dynamicArray.elements = newArray
	dynamicArray.size = len(newArray)
	dynamicArray.capacity = dynamicArray.size
	return nil
}

func (dynamicArray *DynamicArray) RemoveFirstWhere(element interface{}) error {
	if dynamicArray.size == 0 {
		return errors.New("cannot remove from empty array")
	}

	newArray := make([]interface{}, dynamicArray.capacity)

	counter := 0
	found := false
	for _, val := range dynamicArray.elements {
		if element == val && found == false {
			found = true
			continue
		}
		newArray[counter] = val
		counter += 1
	}

	dynamicArray.elements = newArray
	if found == true {
		dynamicArray.size -= 1
	}
	dynamicArray.capacity = dynamicArray.size

	return nil
}

func (dynamicArray *DynamicArray) RemoveLast() error {
	if dynamicArray.size == 0 {
		return errors.New("cannot remove from empty array")
	}

	if dynamicArray.size > 0 {
		dynamicArray.elements = dynamicArray.elements[:len(dynamicArray.elements)-1]
	} else {
		dynamicArray.elements = []interface{}{}
	}

	dynamicArray.size -= 1
	dynamicArray.capacity = dynamicArray.size

	return nil
}

func (dynamicArray *DynamicArray) Elements() []interface{} {
	return dynamicArray.elements
}

func (dynamicArray *DynamicArray) Size() int {
	return dynamicArray.size
}

func (dynamicArray *DynamicArray) IsEmpty() bool {
	return dynamicArray.size == 0
}

func (dynamicArray *DynamicArray) Capacity() int {
	return dynamicArray.capacity
}

func (dynamicArray *DynamicArray) Get(index int) (error, interface{}) {
	if index < 0 || index > (dynamicArray.size-1) {
		return errors.New("index out of range"), nil
	}
	return nil, dynamicArray.elements[index]
}

func (dynamicArray *DynamicArray) Set(index int, value interface{}) error {
	if index < 0 || index > dynamicArray.capacity-1 {
		return errors.New("index out of range")
	}
	dynamicArray.elements[index] = value

	return nil
}

func (dynamicArray *DynamicArray) IndexOf(element interface{}) int {
	for idx, val := range dynamicArray.elements {
		if val == element {
			return idx
		}
	}

	return -1
}

func (dynamicArray *DynamicArray) Contains(element interface{}) bool {
	return dynamicArray.IndexOf(element) != -1
}

func (dynamicArray *DynamicArray) Clear() {
	dynamicArray.elements = make([]interface{}, dynamicArray.capacity)
	dynamicArray.size = 0
}
