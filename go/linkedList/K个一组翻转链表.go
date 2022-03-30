package main

/**
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：


输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
示例 3：

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]
示例 4：

输入：head = [1], k = 1
输出：[1]
提示：

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz
通过次数304,383提交次数457,745

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 前进K步，同时将当前链表翻转

	c := head
	var start, end *ListNode = head, nil

	// 每次要检查之后是否有k个
	for checkValid(c, k) {
		// 如果有，翻转前k个元素，并获得start节点，end节点和others的头节点
		s, e, others := revertK(c, k)
		// 拼接之前链表和这次的链表
		if end != nil {
			end.Next = s
		}
		// 刷新end，指向最后节点
		end = e
		// 把other赋予当前节点，继续循环
		c = others
	}
	// 最后把剩余的拼接到最后
	if end != nil && c != nil {
		end.Next = c
	}
	return start
}

// 将head链表的前k个元素翻转，返回翻转后链表
func revertK(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	var ret *ListNode
	var end *ListNode
	current := head

	for k > 0 {
		// 拼接节点到ret头部
		tempNext := current.Next
		current.Next = ret
		ret = current

		// 赋值
		if end == nil {
			end = current
		}
		// 递进
		current = tempNext
		k--
	}
	return ret, end, current
}

// 检查是否可以继续
func checkValid(head *ListNode, k int) bool {
	for k > 0 {
		if head == nil {
			return false
		}
		head = head.Next
		k--
	}
	return true
}

// 可行，但是不优雅
func reverseKGroupBad(head *ListNode, k int) *ListNode {
	// 前进K步，同时将当前链表翻转

	var newHead, newEnd *ListNode = nil, nil
	var retHead, retEnd *ListNode = nil, nil
	current := head

	idx := 0
	for current != nil {
		// 缓存next
		tempNext := current.Next

		// 把当前节点取下来，放到new链表上
		current.Next = newHead
		newHead = current
		if newEnd == nil {
			newEnd = newHead
		}

		// 递进步骤
		current = tempNext
		idx++
		// 完全一组k个的操作
		if idx%k == 0 {
			if retHead == nil && retEnd == nil {
				// 赋初值
				retHead = newHead
				retEnd = newEnd
			} else {
				// 拼接到尾部
				retEnd.Next = newHead
				retEnd = newEnd
			}
			// 新建slice
			newHead, newEnd = nil, nil

			if !checkValid(current, k) {
				retEnd.Next = current
				return retHead
			}
		}
	}

	return retHead
}
