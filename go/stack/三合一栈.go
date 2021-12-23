package main

import "fmt"

/**
面试题 03.01. 三合一
三合一。描述如何只用一个数组来实现三个栈。

你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。

构造函数会传入一个stackSize参数，代表每个栈的大小。

示例1:

 输入：
["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
 输出：
[null, null, null, 1, -1, -1, true]
说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
示例2:

 输入：
["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, -1, -1]


提示：
0 <= stackNum <= 2

https://leetcode-cn.com/problems/three-in-one-lcci/
*/

func doStackConstruct() {
	s := ConstructorTripleInOne(2, 3)
	s.Push(0, 1)
	s.Push(0, 2)
	s.Push(0, 3)
	fmt.Println(s.Pop(0))
	fmt.Println(s.Pop(0))
	fmt.Println(s.Pop(0))
	fmt.Println(s.Peek(0))
}

type TripleInOne struct {
	data []int
	// 0号存stackSize，此后存stackSize个元素，然后是本stack的length，然后其他栈
}

func ConstructorTripleInOne(stackSize int, stackNum int) TripleInOne {
	totalLength := 1 + (stackSize+1)*stackNum
	data := make([]int, totalLength)
	data[0] = stackSize
	return TripleInOne{
		data: data,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	stackSize := this.data[0]

	// 找到各项下表
	arrStart := this.getStart(stackNum)
	currentLength := this.getStackLength(stackNum)

	if currentLength >= stackSize {
		return
	}
	this.data[arrStart+currentLength] = value

	this.setStackLength(stackNum, currentLength+1)
}

func (this *TripleInOne) Pop(stackNum int) int {
	stackSize := this.getStackSize(stackNum)

	// 找到各项下表
	arrEnd := this.getEnd(stackNum)
	currentLength := this.getStackLength(stackNum)

	if currentLength == 0 {
		return -1
	}

	this.setStackLength(stackNum, currentLength-1)
	return this.data[arrEnd-1-(stackSize-currentLength)]
}

func (this *TripleInOne) Peek(stackNum int) int {
	stackSize := this.getStackSize(stackNum)
	arrEnd := this.getEnd(stackNum)
	currentLength := this.getStackLength(stackNum)

	if currentLength == 0 {
		return -1
	}
	return this.data[arrEnd-1-(stackSize-currentLength)]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.getStackLength(stackNum) == 0
}

func (this *TripleInOne) getStackSize(stackNum int) int {
	return this.data[0]
}

func (this *TripleInOne) getStackLength(stackNum int) int {
	arrEnd := this.getEnd(stackNum)
	return this.data[arrEnd]
}
func (this *TripleInOne) setStackLength(stackNum int, length int) {
	arrEnd := this.getEnd(stackNum)
	this.data[arrEnd] = length
}

func (this *TripleInOne) getStart(stackNum int) int {
	stackSize := this.getStackSize(stackNum)
	return 1 + stackNum*(stackSize+1)
}

func (this *TripleInOne) getEnd(stackNum int) int {
	stackSize := this.getStackSize(stackNum)
	return 1 + stackNum*(stackSize+1) + stackSize
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
