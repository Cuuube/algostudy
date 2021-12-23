package main

/**
面试题 02.01. 移除重复节点
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	// 遍历链表，用哈希表记录当前值的首节点。
	// 不断询问下一节点的值是否已经存在，如果存在，把下下节点拼到下一节点。如果存在，记录下一节点unique，指针继续遍历
	if head == nil {
		return nil
	}

	var uniqueMap = map[int]bool{head.Val: true}

	var currentNode *ListNode = head
	for currentNode.Next != nil {
		_, found := uniqueMap[currentNode.Next.Val]
		// 如果下一节点的value是独立的，则继续顺延爬链表
		if !found {
			uniqueMap[currentNode.Next.Val] = true
			currentNode = currentNode.Next
			continue
		}
		// 如果下一节点value是重复的，把下下节点
		currentNode.Next = currentNode.Next.Next
	}
	return head
}
