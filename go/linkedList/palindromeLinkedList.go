package main

/**
面试题 02.06. 回文链表
编写一个函数，检查输入的链表是否是回文的。



示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true


进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
https://leetcode-cn.com/problems/palindrome-linked-list-lcci/
*/

func isPalindrome(head *ListNode) bool {
	// 如果要O(1)的空间复杂度，也就是说不能使用栈，或其他可变列表。只能在链表本身上操作了。但是会弄脏链表。
	// 1、找到链表中点。奇数：链表中点的位置。偶数：链表后的位置。
	// 2、翻转链表后半部分
	// 3、双指针比较前半和后半是否相等
	if head == nil {
		return true
	}

	midNode := getMidNode(head)

	midNode = reverseLinkedList(midNode)

	// 比较
	left := head
	right := midNode
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func getMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 快慢指针寻找中点
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow会停在中(奇数)点，或者中点之前（偶数），这里调整偶数的话，让slow进入到后半截
	if fast.Next != nil && fast.Next.Next == nil {
		slow = slow.Next
	}
	return slow
}

func reverseLinkedList(head *ListNode) *ListNode {
	var ret *ListNode
	c := head

	for c != nil {
		// 记录下一节点
		next := c.Next

		// 把当前节点c摘出来，放到ret前，并且成为新的ret
		c.Next = ret
		ret = c

		// 继续遍历
		c = next
	}
	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeNormal(head *ListNode) bool {
	// 构建一个栈来检查。O(n) 时间复杂度和 O(n) 空间复杂度
	var stack IntStack = []int{}
	c := head
	for c != nil {
		stack.push(c.Val)
		c = c.Next
	}
	c = head
	for c != nil {
		if stack.pop() != c.Val {
			return false
		}
		c = c.Next
	}
	return true
}

type IntStack []int

func (s *IntStack) push(v int) {
	(*s) = append(*s, v)
}
func (s *IntStack) pop() int {
	stack := *s
	v := stack[len(stack)-1]
	(*s) = stack[0 : len(stack)-1]
	return v
}
