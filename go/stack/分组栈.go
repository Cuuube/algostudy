package main

import "fmt"

/**
面试题 03.03. 堆盘子
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。

当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.

示例1:

 输入：
["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
[[1], [1], [2], [1], [], []]
 输出：
[null, null, null, 2, 1, -1]
示例2:

 输入：
["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
[[2], [1], [2], [3], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, 3]

https://leetcode-cn.com/problems/stack-of-plates-lcci/
*/

func testStackOfPlates() {
	s := ConstructorStackOfPlates(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.PopAt(0))
	fmt.Println(s.PopAt(0))
	fmt.Println(s.PopAt(0))
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
}

// TODO 有bug
// 相当于把一整条栈给分片
type StackOfPlates struct {
	stacks []*Stack
	cap    int
}

func ConstructorStackOfPlates(cap int) StackOfPlates {
	return StackOfPlates{
		stacks: make([]*Stack, 0),
		cap:    cap,
	}
}

func (this *StackOfPlates) Push(val int) {
	// 特殊情况
	if this.cap <= 0 {
		return
	}

	// 初始值
	stackLength := len(this.stacks)
	if stackLength == 0 {
		newStack := ConstructorStack()
		this.stacks = append(this.stacks, &newStack)
		stackLength += 1
	}
	latestStack := this.stacks[stackLength-1]
	// 栈满了
	if latestStack.Size() >= this.cap {
		newStack := ConstructorStack()
		this.stacks = append(this.stacks, &newStack)
		stackLength += 1
		latestStack = &newStack
	}
	latestStack.Push(val)
}

func (this *StackOfPlates) Pop() int {
	stackLength := len(this.stacks)
	if stackLength <= 0 {
		return -1
	}
	latestStack := this.stacks[stackLength-1]
	if latestStack.Size() <= 0 {
		if stackLength <= 1 {
			return -1
		}
		this.stacks = this.stacks[0 : len(this.stacks)-1]
		return this.Pop()
	}
	return latestStack.Pop()
}

func (this *StackOfPlates) PopAt(index int) int {
	stackLength := len(this.stacks)
	if index >= stackLength {
		return -1
	}
	stack := this.stacks[index]
	val := stack.Pop()

	// 如果stack为空，删除本stack
	if stack.IsEmpty() {
		temps := this.stacks[index+1 : stackLength]
		this.stacks = append(this.stacks[0:index], temps...)
	}
	return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
