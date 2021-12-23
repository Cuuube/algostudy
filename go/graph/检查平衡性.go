package main

/**
面试题 04.04. 检查平衡性
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。


示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
      1
     / \
    2   2
   / \
  3   3
 / \
4   4
返回 false 。
https://leetcode-cn.com/problems/check-balance-lcci/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 深度遍历法，分别记录遍历到底的最长depth。
	// 如果当前到底的路径depth和最长相差2以上，说明不平衡
	// 需要用一个flag，保证子树里面没有差距为2的。
	isBalance := true
	searchTree(root, 0, &isBalance)

	return isBalance
}
func searchTree(node *TreeNode, currentDeepth int, isBalance *bool) (max int) {
	if node == nil {
		return currentDeepth - 1
	}
	currentDeepth += 1
	// 递归，实质是倒序遍历
	leftMax := searchTree(node.Left, currentDeepth, isBalance)
	rightMax := searchTree(node.Right, currentDeepth, isBalance)
	max = leftMax
	min := rightMax
	if rightMax > max {
		max = rightMax
		min = leftMax
	}
	// 关键的判断条件
	if max-min >= 2 {
		*isBalance = false
	}
	return
}

// 	if node == nil {
// 		return currentDeepth - 1, currentDeepth - 1
// 	}
// 	currentDeepth += 1
// 	// 左侧节点
// 	leftMax1, rightMax1 := searchTree(node.Left, currentDeepth)
// 	leftMax2, rightMax2 := searchTree(node.Right, currentDeepth)

// 	leftMax = leftMax1
// 	rightMax = rightMax1
// 	if leftMax2 > leftMax {
// 		leftMax = leftMax2
// 	}
// 	if rightMax2 > rightMax {
// 		rightMax = rightMax2
// 	}
// 	return
// }
// func searchTree(node *TreeNode, currentDeepth int) (max, min int) {
// 	if node == nil {
// 		return currentDeepth - 1, currentDeepth - 1
// 	}
// 	currentDeepth += 1
// 	// 左侧节点
// 	leftMax, leftMin := searchTree(node.Left, currentDeepth)
// 	rightMax, rightMin := searchTree(node.Right, currentDeepth)

// 	max = rightMax
// 	min = rightMin
// 	if leftMax > rightMax {
// 		max = leftMax
// 	}
// 	if leftMin < rightMin {
// 		min = leftMin
// 	}
// 	return
// }

// func isBalanced(root *TreeNode) bool {
// if root == nil {
// 	return true
// }
// return compareBalanced(root.Left, root.Right)
// }
// func compareBalanced(left *TreeNode, right *TreeNode) bool {
//     // 检查左右是否相似
//     // 当一侧是null，另一侧不是null且有孩子，说明不平衡
//     if left == nil && right != nil && (right.Left != nil || right.Right != nil) {
//         return false
//     }
//     if right == nil && left != nil && (left.Left != nil || left.Right != nil) {
//         return false
//     }
//     if left == nil && right == nil  {
//         return true
//     }
//     return compareBalanced(left.Left, right.Right)
// }
