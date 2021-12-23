package main

/**
面试题 03.05. 栈排序
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。

示例1:

 输入：
["SortedStack", "push", "push", "peek", "pop", "peek"]
[[], [1], [2], [], [], []]
 输出：
[null,null,null,1,null,2]
示例2:

 输入：
["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
[[], [], [], [1], [], []]
 输出：
[null,null,null,null,null,true]
说明:

栈中的元素数目在[0, 5000]范围内。

https://leetcode-cn.com/problems/sort-of-stacks-lcci/
*/

type SortedStack struct {
	stack     *Stack
	tempStack *Stack
}

func Constructor() SortedStack {
	s := ConstructorStack()
	st := ConstructorStack()
	return SortedStack{stack: &s, tempStack: &st}
}

func (this *SortedStack) Push(val int) {
	if this.stack.IsEmpty() {
		this.stack.Push(val)
		return
	}
	var v int
	// 进来一个数，不断弹出栈，比较大小，如果栈内数比较小，就放右边，直到栈内的数比val大
	for !this.stack.IsEmpty() {
		v = this.stack.Pop()
		if v < val {
			this.tempStack.Push(v)
			continue
		}
		this.stack.Push(v)
		break
	}
	this.stack.Push(val)
	// 把temp的数字返回
	for !this.tempStack.IsEmpty() {
		this.stack.Push(this.tempStack.Pop())
	}
}

func (this *SortedStack) Pop() {
	this.stack.Pop()
}

func (this *SortedStack) Peek() int {
	return this.stack.top()
}

func (this *SortedStack) IsEmpty() bool {
	return this.stack.IsEmpty()
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
