package main

/**
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。



示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
https://leetcode-cn.com/problems/zero-matrix-lcci/
*/
func setZeroes(matrix [][]int) {
	// 分别记录需要清空的行和列，第一步先记录，第二步做清理
	clearX := map[int]bool{}
	clearY := map[int]bool{}
	for x, row := range matrix {
		for y, cell := range row {
			if cell == 0 {
				clearX[x] = true
				clearY[y] = true
			}
		}
	}

	for x, row := range matrix {
		clearRow := false
		if clearX[x] {
			clearRow = true
		}
		for y := range row {
			if clearRow {
				matrix[x][y] = 0
				continue
			}
			if clearY[y] {
				matrix[x][y] = 0
			}
		}
	}
}
