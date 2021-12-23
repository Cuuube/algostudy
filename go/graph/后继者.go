package main

/**
面试题 04.06. 后继者
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

示例 1:

输入: root = [2,1,3], p = 1

  2
 / \
1   3

输出: 2
示例 2:

输入: root = [5,3,6,2,4,null,null,1], p = 6

      5
     / \
    3   6
   / \
  2   4
 /
1

输出: null

https://leetcode-cn.com/problems/successor-lcci/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	// 概念：中序后继。相当于找到树内比自己大，且val差的最小的元素。
	// 1、O(logn)的方法：遍历树，找到最接近的元素。（略）
	// 2、O(n)方法：拍平为数组后遍历。（略）
	// 3、特殊方法：从根节点往下探，有如下特征：
	// 我们需要的节点总是比目标节点大一点。所以：
	// 若当前节点小于目标节点，向右走
	// 若当前节点大于目标节点，记录当前节点，和ret比较，记录最接近的，然向左走
	// https://www.cnblogs.com/FannyChung/p/bst-de-zhong-xu-hou-ji.html
	var ret *TreeNode // 这里要从比自己大的开始记录，不能记录跟节点！万一不大，就会一直是错的！
	cur := root
	for cur != nil {
		if cur.Val <= p.Val {
			cur = cur.Right
			continue
		}
		if ret == nil || cur.Val < ret.Val {
			ret = cur
		}
		cur = cur.Left
	}
	return ret
}
