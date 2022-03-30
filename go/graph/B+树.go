package main

// B+树节点：范围节点，非叶子节点不带数据。叶子节点具有指向后叶子节点的指针
type BplusTreeNode struct {
	id       uint64
	value    uintptr
	children []BplusTreeNode
	nextNode *BplusTreeNode // 指向下一个节点的指针
}

func buildBplusTree() {
	// TODO
}
