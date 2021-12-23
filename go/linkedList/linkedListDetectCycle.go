package main

/**
面试题 02.08. 环路检测
给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。若环不存在，请返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。



示例 1：



输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：



输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：



输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。


进阶：

你是否可以不用额外空间解决此题？
*/

func detectCycle(head *ListNode) *ListNode {
	// 如果要空间O(1)，不用哈希表，只能用指针。
	// 先用快慢指针，让快指针走2步，慢指针走1步。最终快慢指针会在环上相遇。
	// 当快指针和慢指针相遇，相遇点和入环的距离正好是起点到入环的整数倍距离。这时候再启动一个指针从头向后走，等到慢指针和新指针相遇，该节点就是环口
	// 证明：https://leetcode-cn.com/problems/linked-list-cycle-lcci/solution/guan-jian-zai-yu-ru-he-zheng-ming-by-qia-lfiw/
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head.Next
	delay := head
	delayStarted := false
	for fast != nil && fast.Next != nil {
		if fast == slow && !delayStarted {
			delayStarted = true
			continue
		}
		if slow == delay {
			return delay
		}

		slow = slow.Next
		if delayStarted {
			delay = delay.Next
		} else {
			fast = fast.Next.Next
		}
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycleNormal(head *ListNode) *ListNode {
	// 用hash表记录每个节点的地址，是否出现过
	unique := make(map[*ListNode]bool)
	c := head
	for c != nil {
		_, found := unique[c]
		if found {
			return c
		}
		unique[c] = true
		c = c.Next
	}
	return nil
}
