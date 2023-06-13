package vector

import "sort"

// Аn analogue of a std::vector in C++
type Vector[typename any] []typename

// Get the number of elements in vector
func (vector *Vector[typename]) Size() int {
	return len(*vector)
}

// Get vector's capacity
func (vector *Vector[typename]) Capacity() int {
	return cap(*vector)
}

// Add a new element at the end of the vector
func (vector *Vector[typename]) PushBack(arg typename) {
	*vector = append(*vector, arg)
}

// Get the last element of the vector
func (vector *Vector[typename]) Back(arg typename) typename {
	return (*vector)[len(*vector)-1]
}

// Stable sort
func (vector *Vector[typename]) StableSort(left, right int, comp func(int, int) bool) {
	if left == right-1 {
		return
	}
	sort.SliceStable([]typename(*vector)[left:right], comp)
}

// Ordinary sort
func (vector *Vector[typename]) Sort(left, right int, comp func(int, int) bool) {
	if left == right-1 {
		return
	}
	sort.Slice([]typename(*vector)[left:right], comp)
}
