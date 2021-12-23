package main

import "fmt"

/**
面试题 03.04. 化栈为队
实现一个MyQueue类，该类用两个栈来实现一个队列。


示例：

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false

说明：

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
https://leetcode-cn.com/problems/implement-queue-using-stacks-lcci/
*/

func testMyQueue() {
	s := ConstructorMyQueue()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}

// 思路：构造两个栈，一个正向一个反向。从正向栈推入数据时，立马构造反向栈。从反向栈弹出数据时，立马刷新正向栈
// 优化：其实不用搞那么多操作，只需要在推入是，把数据拿到正向栈里，推出时，把数据拿到反向栈里就行。
type MyQueue struct {
	stack        *Stack
	reverseStack *Stack
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	s := ConstructorStack()
	rs := ConstructorStack()
	return MyQueue{
		stack:        &s,
		reverseStack: &rs,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.stack.IsEmpty() && !this.reverseStack.IsEmpty() {
		this.swap()
	}

	this.stack.Push(x)

	// temp := ConstructorStack()
	// revesedStack := ConstructorStack()
	// for !this.stack.IsEmpty() {
	// 	v := this.stack.Pop()
	// 	temp.Push(v)
	// 	revesedStack.Push(v)
	// }
	// this.reverseStack = &revesedStack

	// for !temp.IsEmpty() {
	// 	v := temp.Pop()
	// 	this.stack.Push(v)
	// }
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.reverseStack.IsEmpty() && this.stack.IsEmpty() {
		return -1
	} else if this.reverseStack.IsEmpty() && !this.stack.IsEmpty() {
		this.swap()
	}

	return this.reverseStack.Pop()

	// value := this.reverseStack.Pop()

	// temp := ConstructorStack()
	// stack := ConstructorStack()
	// for !this.reverseStack.IsEmpty() {
	// 	v := this.reverseStack.Pop()
	// 	temp.Push(v)
	// 	stack.Push(v)
	// }
	// this.stack = &stack
	// for !temp.IsEmpty() {
	// 	v := temp.Pop()
	// 	this.reverseStack.Push(v)
	// }
	// return value
}

func (this *MyQueue) swap() {
	var (
		raw    *Stack
		target *Stack
	)
	if !this.stack.IsEmpty() && this.reverseStack.IsEmpty() {
		raw = this.stack
		target = this.reverseStack
	} else if this.stack.IsEmpty() && !this.reverseStack.IsEmpty() {
		raw = this.reverseStack
		target = this.stack
	} else {
		return
	}
	for !raw.IsEmpty() {
		target.Push(raw.Pop())
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.reverseStack.IsEmpty() && this.stack.IsEmpty() {
		return -1
	} else if this.reverseStack.IsEmpty() && !this.stack.IsEmpty() {
		this.swap()
	}
	return this.reverseStack.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.IsEmpty() && this.reverseStack.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
