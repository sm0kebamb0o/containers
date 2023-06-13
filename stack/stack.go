package stack

// An analog to std::stack in C++
type Stack[typename any] []typename

// Get the last added element
func (stack *Stack[typename]) Top() typename {
	return (*stack)[len(*stack)-1]
}

// Get the size of the stack
func (stack *Stack[typename]) Size() int {
	return len(*stack)
}

// Erase the last element
func (stack *Stack[typename]) Pop() bool {
	switch stack.Size() {
	case 0:
		return false
	case 1:
		*stack = nil
	default:
		*stack = (*stack)[:stack.Size()-1]
	}
	return true
}

// Add a new element
func (stack *Stack[typename]) Push(arg typename) {
	*stack = append(*stack, arg)
}

// Test whether stack is empty
func (stack *Stack[typename]) Empty() bool {
	return stack.Size() == 0
}
