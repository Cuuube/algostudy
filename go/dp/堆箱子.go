package main

import "sort"

/**
面试题 08.13. 堆箱子
堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。

输入使用数组[wi, di, hi]表示每个箱子。

示例1:

 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
 输出：6
示例2:

 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
 输出：10
提示:

箱子的数目不大于3000个。
*/

func pileBox(box [][]int) int {
	// 优化：box先排序，递归时不管前面的。就可以少了很多计算了
	var b Boxes = box
	sort.Sort(b)
	return do([]int{0, 0, 0}, box)
}

func do(maxBox []int, box [][]int) int {
	ret := 0

	// 一次遍历+递归。得到所有可能的N叉树，然后比较N叉结果。
	for i, currentBox := range box {
		if !canAdd(maxBox, currentBox) {
			continue
		}
		// 剩余的箱子，必须比自己大。优化：先排序，后直接用box[i+1:]就行
		otherBox := box[i+1:]
		// otherBox := make([][]int, 0)
		// for j, b := range box {
		// 	if i != j && canAdd(currentBox, b) {
		// 		otherBox = append(otherBox, b)
		// 	}
		// }
		// otherBox = append(otherBox, box[0:i]...)
		// otherBox = append(otherBox, box[i+1:]...)

		max := do(currentBox, otherBox)
		if max > ret {
			ret = max
		}
	}
	return ret + getHeight(maxBox)
}

func canAdd(flagBox, box []int) bool {
	return box[0] > flagBox[0] && box[1] > flagBox[1] && box[2] > flagBox[2]
}

func getHeight(box []int) int {
	return box[2]
}

type Boxes [][]int

func (a Boxes) Len() int      { return len(a) }
func (a Boxes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Boxes) Less(i, j int) bool { // 按长宽高最小的排序
	if a[i][0] == a[j][0] {
		if a[i][1] == a[j][1] {
			return a[i][2] < a[j][2]
		}
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}
