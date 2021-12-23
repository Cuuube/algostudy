package main

/**
面试题 02.05. 链表求和
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。



示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
https://leetcode-cn.com/problems/sum-lists-lcci/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 两个链表，可以一次遍历，直到两个链表都到了终点。然后加它们的值
	var result *ListNode
	c1 := l1
	c2 := l2
	weight := 1
	needPlus := false
	for c1 != nil || c2 != nil {
		value := 0
		if c1 != nil {
			value += (c1.Val * weight)
			c1 = c1.Next
		}
		if c2 != nil {
			value += (c2.Val * weight)
			c2 = c2.Next
		}
		if needPlus {
			value += 1
			needPlus = false
		}
		if value >= 10 {
			value = value % 10
			needPlus = true
		}
		// 组成链表
		result = appendLinkedList(result, value)
	}
	if needPlus {
		result = appendLinkedList(result, 1)
	}
	return result
}

func appendLinkedList(root *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value, Next: nil}
	if root == nil {
		return newNode
	}
	c := root
	for c.Next != nil {
		c = c.Next
	}
	c.Next = newNode
	return root
}

func getRealNumbers(l *ListNode) int {
	if l == nil {
		return 0
	}
	c := l
	weight := 1
	sum := 0
	for c != nil {
		sum += (c.Val * weight)
		weight = weight * 10
		c = c.Next
	}
	return sum
}
