package main

// --- 普通Stack
type Stack struct {
	stack []int
}

func ConstructorStack() Stack {
	return Stack{
		stack: make([]int, 0),
	}
}

func (this *Stack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *Stack) Pop() int {
	length := len(this.stack)
	if length <= 0 {
		return -1
	}
	value := this.stack[length-1]
	this.stack = this.stack[0 : length-1]
	return value
}

func (this *Stack) Peek() int {
	length := len(this.stack)
	if length <= 0 {
		return -1
	}
	value := this.stack[length-1]
	return value
}

func (this *Stack) top() int {
	return this.Peek()
}

func (this *Stack) IsEmpty() bool {
	return this.Size() == 0
}

func (this *Stack) Size() int {
	return len(this.stack)
}
