package main

/**
面试题 04.08. 首个共同祖先
设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。

例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]

    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4
示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// 共同祖先，可以理解为倒序遍历，分别判断一个节点的左右两叉是否各包含一个目标节点
	// 在current节点，若本节点是目标节点，返回本节点。如果左侧包含目标节点，右侧也包含目标节点，返回本节点。否则，返回单侧节点。若一侧为空，返回另一侧节点。若两侧都为空，返回nil。

	return contains(root, p, q)
}

func contains(node, p, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node == p || node == q {
		return node
	}
	leftNode := contains(node.Left, p, q)
	rightNode := contains(node.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return node
	}
	if leftNode == nil {
		return rightNode
	}
	return leftNode
}

// func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
// 	// 共同祖先，可以理解为倒序遍历，分别判断一个节点的左右两叉是否各包含一个目标节点
// 	// 在current节点，leftContains() && rightContains()

// 	var target TreeNode
// 	contains(root, p, q, &target)
// 	return &target
// }

// func contains(node, p, q, target *TreeNode) (bool, bool) {
// 	if node == nil {
// 		return false, false
// 	}
// 	if node == p || node == q {
// 		return true, true
// 	}
// 	lc1, lc2 := contains(node.Left, p, q, target)
// 	rc1, rc2 := contains(node.Right, p, q, target)
// 	if lc1 || lc2 && rc1 || rc2 {
// 		*target = *node
// 	}
// 	return lc1 || lc2, rc1 || rc2
// }
