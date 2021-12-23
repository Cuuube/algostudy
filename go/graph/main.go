package main

func main() {
	// testFindWhetherExistsPath()
	testpPathSum()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
