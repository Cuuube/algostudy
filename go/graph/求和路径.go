package main

import "fmt"

/*
面试题 04.12. 求和路径
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

3
解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
提示：

节点总数 <= 10000

https://leetcode-cn.com/problems/paths-with-sum-lcci/
*/

func testpPathSum() {
	root := &TreeNode{Val: -2, Right: &TreeNode{Val: -3}}
	fmt.Println(pathSum(root, -5))
}

func pathSum(root *TreeNode, sum int) int {
	// 深度优先遍历，每走一步，就记录当前节点的sum值
	// 并且对左右节点执行同样的深度优先遍历
	// 相当于深度优先遍历的平方
	cnt := 0
	if root == nil {
		return cnt
	}
	cnt = dfs(root, sum, cnt)
	cnt += pathSum(root.Left, sum)
	cnt += pathSum(root.Right, sum)
	return cnt
}

func dfs(node *TreeNode, sum int, cnt int) int {
	if node == nil {
		return cnt
	}
	sum -= node.Val
	if sum == 0 {
		cnt += 1
	}
	cnt = dfs(node.Left, sum, cnt)
	cnt = dfs(node.Right, sum, cnt)
	return cnt
}

// 别人的
// var dfs func(*TreeNode, int)
// func pathSum(root *TreeNode, sum int) (ans int) {
//     dfs = func(node *TreeNode, num int) {
//         if node == nil {return}
//         num -= node.Val
//         if num == 0 {ans++}
//         dfs(node.Left, num)
//         dfs(node.Right, num)
//     }
//     if root == nil {return 0}
//     dfs(root, sum)
//     l := pathSum(root.Left, sum)
//     r := pathSum(root.Right, sum)
//     return l + r + ans
// }
