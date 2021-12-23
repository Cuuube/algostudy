package main

/**
面试题 04.02. 最小高度树
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
https://leetcode-cn.com/problems/minimum-height-tree-lcci/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// 一开始想的是能不能设计一棵树，让无论怎么推入数据，都可以保持平衡。但是对树的知识太薄弱，没想出来。
	// 题目因为是已经排好序了，那么只需要按照一定顺序构建树就好了。
	// 从中间开始，递归左右两侧
	length := len(nums)
	if length == 0 {
		return nil
	} else if length == 1 {
		return &TreeNode{Val: nums[0]}
	}
	// 找到中间节点
	mid := length / 2
	root := &TreeNode{Val: nums[mid]}
	// 左右两侧用递归
	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1 : length])
	return root
}
