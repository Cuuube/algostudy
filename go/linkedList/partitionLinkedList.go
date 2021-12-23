package main

/**
面试题 02.04. 分割链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你不需要 保留 每个分区中各节点的初始相对位置。



示例 1：


输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]
示例 2：

输入：head = [2,1], x = 2
输出：[1,2]


提示：

链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200

https://leetcode-cn.com/problems/partition-list-lcci/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 维护两个分离的节点，分别是小于x的和大于等于x的
	// 遍历链表，如果当前节点的值小于x，把自己拼到小节点中去，然后继续遍历，直到终点

	var smallNode *ListNode // 小链表
	var largeNode *ListNode // 大链表

	var currentNode = head
	for currentNode != nil { // 不断遍历
		// 把当前节点处理为独立节点：记录next，然后砍掉自己的next。
		nextNode := currentNode.Next
		currentNode.Next = nil

		// 按照value分别拼接到大链表或者小链表
		if currentNode.Val < x {
			if smallNode == nil {
				smallNode = currentNode
			} else {
				smallNode = joinLinkedList(smallNode, currentNode)
			}
		} else {
			if largeNode == nil {
				largeNode = currentNode
			} else {
				largeNode = joinLinkedList(largeNode, currentNode)
			}
		}

		// 继续循环
		currentNode = nextNode
	}
	return joinLinkedList(smallNode, largeNode)
}

// 链表拼接
func joinLinkedList(linkedList1 *ListNode, linkedList2 *ListNode) *ListNode {
	if linkedList1 == nil && linkedList2 != nil {
		return linkedList2
	}
	if linkedList1 != nil && linkedList2 == nil {
		return linkedList1
	}
	if linkedList1 == nil && linkedList2 == nil {
		return nil
	}

	// 找链表终点，把后面这个链表挂上去
	n := linkedList1
	for n.Next != nil {
		n = n.Next
	}
	n.Next = linkedList2
	return linkedList1
}
