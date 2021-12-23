package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
// 	// 方案：左树的最大值一定小于自身，右树的最小值大于自身
// 	flag := true
// 	BSTToArrayAndCheck(root, &flag)
// 	return flag

// }

// func BSTToArrayAndCheck(node *TreeNode, valid *bool) (max, min int) {
// 	if node == nil {
// 		return -0xfffffffffffffff, 0xfffffffffffffff
// 	}
// 	if node.Left != nil {
// 		leftMax, leftMin := BSTToArrayAndCheck(node.Left, valid)
// 		if leftMax >= node.Val {
// 			*valid = false
// 		}

// 		max = leftMax
// 		min = leftMin

// 	}
// 	if node.Right != nil {
// 		rightMax, rightMin := BSTToArrayAndCheck(node.Right, valid)
// 		if rightMin <= node.Val {
// 			*valid = false
// 		}
// 		if rightMax > max {
// 			max = rightMax
// 		}
// 		if rightMin < min {
// 			min = rightMin
// 		}
// 	}
// 	if node.Val > max {
// 		max = node.Val
// 	}
// 	if node.Val < min {
// 		min = node.Val
// 	}
// 	return
// }

func isValidBSTBad(root *TreeNode) bool {
	// 方案：先拍平为数组，再检查数组是否有序

	arr := BSTToArray(root)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] {
				return false
			}
		}
	}
	return true
}

func BSTToArray(node *TreeNode) []int {
	ret := make([]int, 0)
	if node == nil {
		return ret
	}
	if node.Left != nil {
		ret = append(ret, BSTToArray(node.Left)...)
	}
	ret = append(ret, node.Val)
	if node.Right != nil {
		ret = append(ret, BSTToArray(node.Right)...)
	}
	return ret
}
