package main

type Stack []string

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// PushN a new slice of values onto the stack
func (s *Stack) PushN(str []string) {
	*s = append(*s, str...) // Simply append the new value to the end of the stack
}

// Pop Removes and return stop element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// PopN Removes and returns top N elements of stack. Return false if stack is empty.
func (s *Stack) PopN(n int) ([]string, bool) {
	if len(*s) < n {
		return nil, false
	} else {
		element := (*s)[len(*s)-n:] // Index into the slice and obtain the element.
		*s = (*s)[:len(*s)-n]       // Remove it from the stack by slicing it off.
		return element, true
	}
}
