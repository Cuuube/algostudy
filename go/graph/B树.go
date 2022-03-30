package main

import (
	"sort"
	"unsafe"
)

// B树节点：范围节点+区间节点
type BTreeNode struct {
	id       uint64
	value    uintptr
	children []*BTreeNode
}

func (t *BTreeNode) isLeaf() bool {
	return t.children == nil
}

func (t *BTreeNode) getMinRangeNode(id uint64) *BTreeNode {
	// id小于所有，或者B树是空的，直接返回nil
	if len(t.children) == 0 || t.children[0].id < id {
		return nil
	}

	var currentNode *BTreeNode = t
	for currentNode != nil {
		var rangeNode *BTreeNode
		for idx := 0; idx < len(currentNode.children); idx++ {
			node := currentNode.children[idx]
			// 记录前一个不是leaf节点的节点
			if !node.isLeaf() {
				rangeNode = node
				continue
			}
			if node.id == id {
				// 返回父节点
				return currentNode
			} else if node.id > id {
				// id超出，返回之前的范围节点
				currentNode = rangeNode
				break
			}
		}
	}
	return currentNode
}

// 查找数据
func (t *BTreeNode) Get(id uint64) *BTreeNode {
	minRangeNode := t.getMinRangeNode(id)
	for idx, node := range minRangeNode.children {
		if node.id == id {
			return minRangeNode.children[idx]
		}
	}
	return nil
}

func (t *BTreeNode) Set(id uint64, data uintptr) {
	// 查找id是否存在，如果存在，直接更改
	minRangeNode := t.getMinRangeNode(id)

	for _, node := range minRangeNode.children {
		if node.id == id {
			node.value = data
			return
		}
	}

	// 如果不存在，新建
	// 1、寻找底层区间
	// 2、插入数据
	newChildren := make([]*BTreeNode, 0)
	flag := false
	newNode := BTreeNode{id: id, value: data}

	for idx, node := range minRangeNode.children {
		if node.id > id && !flag {
			// 插入
			flag = true
			newChildren = append(newChildren, &newNode)
		}
		newChildren = append(newChildren, minRangeNode.children[idx])
	}
	if !flag {
		newChildren = append(newChildren, &newNode)
	}
	// 判断是否要分裂:TODO
	// minRangeNode.checkSplit()
}

func buildBTree(data []int) *BTreeNode {
	sort.Slice(data, func(i, j int) bool {
		return i < j
	})

	node := BTreeNode{}
	node.children = make([]*BTreeNode, 0)

	for idx := range data {
		ptr := unsafe.Pointer(&data[idx])
		node.Set(uint64(idx), uintptr(ptr))
	}
	// 如果元素过多，提出更高一级

	return &node
}
