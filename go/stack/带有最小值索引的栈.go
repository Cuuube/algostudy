package main

/**
面试题 03.02. 栈的最小值
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。


示例：

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
https://leetcode-cn.com/problems/min-stack-lcci/
*/

// 题目的需求：push、pop、min的时间复杂度为O(1)，也就是不能遍历
// 有什么办法可以快速获取最小值呢？
// 利用栈的后入先出特性，可以每次加入数据，都记录当前最小值。然后把最小值同步存入另一个栈
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	min := this.GetMin()
	if x < min {
		min = x
	}
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, min)
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	// value := this.stack[length - 1]
	this.stack = this.stack[0 : length-1]
	this.minStack = this.minStack[0 : length-1]
}

func (this *MinStack) Top() int {
	length := len(this.stack)
	value := this.stack[length-1]
	return value
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return 0xfffffffffffffff
	}
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
