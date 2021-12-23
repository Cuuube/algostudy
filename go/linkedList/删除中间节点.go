package main

/**
面试题 02.02. 返回倒数第 k 个节点
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	// 使用快慢指针，把倒数k个，变成两个指针的固定距离。就可以通过一次遍历找到目标值
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if k > 0 {
			k -= 1
		} else {
			slow = slow.Next
		}
	}
	return slow.Val
}

func kthToLastBad(head *ListNode, k int) int {
	// 先计算链表长度length，然后通过length - k再次遍历
	length := 0
	current := head
	for current != nil {
		length += 1
		current = current.Next
	}

	i := length - k
	idx := 0
	current = head
	for idx != i {
		idx += 1
		current = current.Next
	}
	return current.Val
}

/**
面试题 02.03. 删除中间节点
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。

假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。

例如，传入节点 c（位于单向链表 a->b->c->d->e->f 中），将其删除后，剩余链表为 a->b->d->e->f



示例：

输入：节点 5 （位于单向链表 4->5->1->9 中）
输出：不返回任何数据，从链表中删除传入的节点 5，使链表变为 4->1->9
https://leetcode-cn.com/problems/delete-middle-node-lcci/
*/
func deleteNode(node *ListNode) {
	// 只能获取到当前节点，那么也就是无法把节点重新挂载。
	// 可以改变自己的val为下一节点，然后拿掉下一节点
	current := node
	current.Val = current.Next.Val
	current.Next = current.Next.Next
}
