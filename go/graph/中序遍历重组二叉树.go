package main

/**
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

相关的leecode：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
https://zhuanlan.zhihu.com/p/40025912
*/

func BuildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 特点：后序遍历的最后一个节点，是根节点；而中序遍历的根节点在中间
	rootVal := postorder[len(postorder)-1]

	// 根据根节点，从inOrder中切出左右节点
	inorderLeft, inorderRight := splitByVal(inorder, rootVal)
	var postorderLeft, postorderRight []int = []int{}, []int{}
	// 如果没有左节点，说明postOrder中剩余部分都是右节点
	if len(inorderLeft) == 0 {
		postorderRight = postorder[0 : len(postorder)-1]
	} else if len(inorderRight) == 0 {
		// 如果没有右节点，说明postOrder中剩余部分都是左节点
		postorderLeft = postorder[0 : len(postorder)-1]
	} else {
		// 如果都有，要再切分一次。注意先用右节点的第一个元素切分
		postorderLeft, postorderRight = splitByVal(postorder[0:len(postorder)-1], inorderRight[0])
		// 然后再补进去
		temp := make([]int, 0)
		temp = append(temp, inorderRight[0])
		temp = append(temp, postorderRight...)
		postorderRight = temp
	}

	node := &TreeNode{
		Val:   rootVal,
		Left:  BuildTree(inorderLeft, postorderLeft),
		Right: BuildTree(inorderRight, postorderRight),
	}

	return node
}

func splitByVal(arr []int, val int) ([]int, []int) {
	var pre, after []int = make([]int, 0), make([]int, 0)
	flag := false
	for _, v := range arr {
		if v == val {
			flag = true
			continue
		}
		if !flag {
			pre = append(pre, v)
		} else {
			after = append(after, v)
		}
	}
	return pre, after
}

// var (
// 	inOrderIndex   = map[int]int{}
// 	postOrderIndex = map[int]int{}
// )

// func getIndexInOrder(val int, inorder []int) int {
// 	if len(inOrderIndex) == 0 {
// 		for idx, v := range inorder {
// 			inOrderIndex[v] = idx
// 		}
// 	}
// 	return inOrderIndex[val]
// }
// func getIndexPostOrder(val int, postorder []int) int {
// 	if len(postOrderIndex) == 0 {
// 		for idx, v := range postorder {
// 			postOrderIndex[v] = idx
// 		}
// 	}
// 	return postOrderIndex[val]
// }
