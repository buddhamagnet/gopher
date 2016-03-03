package stack

import "errors"

type Stack []interface{}

var ErrStackEmpty = errors.New("stack is empty")

func (stack *Stack) Pop() (interface{}, error) {
	if len(*stack) == 0 {
		return nil, ErrStackEmpty
	}
	copied := *stack

	x := copied[len(copied)-1]
	*stack = copied[:len(copied)-1]
	return x, nil
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, ErrStackEmpty
	}
	return stack[len(stack)-1], nil
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
