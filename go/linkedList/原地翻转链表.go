package main

/**
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000


进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？



*/

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	// 定义返回值
	var ret *ListNode = nil
	// 当前节点，按顺序向后便利
	c := head
	for c != nil {
		temp := c.Next // 把后面链表部分摘出来
		c.Next = ret   // 当前节点单独摘出来，next指向反转好的链表
		ret = c        // 然后把反转好的赋值回ret
		c = temp       // 将之前缓存的剩余链表部分继续遍历
	}
	return ret
}

// 递归方式翻转链表
// pre:已翻转的部分，cur：未翻转的部分
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	if cur.Next == nil {
		cur.Next = pre
		return cur
	}
	// 记录next
	next := cur.Next
	cur.Next = pre // 把头节点拼到前面
	pre = cur      // 重新制定头节点
	return reverse(pre, next)
}
