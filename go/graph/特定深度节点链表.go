package main

/**
面试题 04.03. 特定深度节点链表
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。



示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
https://leetcode-cn.com/problems/list-of-depth-lcci/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	// 其实质是树的广度优先遍历
	ret := make([]*ListNode, 0)

	nodes := []*TreeNode{tree}
	for len(nodes) > 0 {
		nextNodes := make([]*TreeNode, 0)
		listNode := new(ListNode)
		ret = append(ret, listNode)

		// 遍历树
		for idx, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}

			// 把node装入链表
			if idx == 0 {
				listNode.Val = node.Val
			} else {
				listNode.Next = &ListNode{Val: node.Val}
				listNode = listNode.Next
			}
		}
		nodes = nextNodes
	}
	return ret
}
