package vector

type Vector[typename comparable] []typename

func (vector *Vector[typename]) Size() int {
	return len(*vector)
}

func (vector *Vector[typename]) Capacity() int {
	return cap(*vector)
}

func (vector *Vector[typename]) Push_back(arg typename) {
	*vector = append(*vector, arg)
}

func (vector *Vector[typename]) Back(arg typename) typename {
	return (*vector)[len(*vector)-1]
}
