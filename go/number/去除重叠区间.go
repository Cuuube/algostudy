package main

import (
	"fmt"
	"sort"
)

/**
from 字节面试题
给定一组x轴上的n组闭区间，去掉尽可能少的闭区间，令其他闭区间都不相交。输出去掉的最少闭区间数
如：intput：3，[[10, 20], [15, 10], [20, 15]]
	必须去掉两个，才能让剩余的不相交。所以output：2
*/

// 去除重叠区间
func checkRanges(nums int, ranges [][]int) int {
	// 有几个点：
	// 1、区间左右值，不一定是小的在左，大的在右。因此需要先整理区间
	// 2、区间之间的排序有些混乱，最好可以先按照左端点、右端点来排序
	// 3、每一步都执行检查，保证自己的右端点小于下一个区间的左端点。
	newRanges := make([][]int, len(ranges))
	for i, arr := range ranges {
		if arr[0] < arr[1] {
			newRanges[i] = []int{arr[0], arr[1]}
		} else {
			newRanges[i] = []int{arr[1], arr[0]}
		}
	}

	var sortedNewRanges SortByStart = newRanges
	sort.Sort(sortedNewRanges)

	removeList := make([][]int, 0)
	idx := 0
	// 从头像尾部遍历
	for idx < sortedNewRanges.Len() {
		// 记录当前指针和操作到的指针
		current := sortedNewRanges[idx]
		nextIdx := idx + 1
		// 保证操作指针在数组内，让操作指针不断向后。直到遇到有效区间时，让当前指针指向操作指针
		for nextIdx < sortedNewRanges.Len() {
			next := sortedNewRanges[nextIdx]
			// 如果后一个区间的起点小于前区间的终点，移除后区间
			if next[0] <= current[1] {
				removeList = append(removeList, next)
			}
			nextIdx += 1
		}
		idx = nextIdx
	}
	fmt.Printf("sortedList: %+v\nremoveList: %+v\n", sortedNewRanges, removeList)
	return len(removeList)
}

type SortByStart [][]int

func (a SortByStart) Len() int      { return len(a) }
func (a SortByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByStart) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}
