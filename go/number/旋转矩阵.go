package main

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
注意：本题与主站 48 题相同：https://leetcode-cn.com/problems/rotate-image/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rotate(matrix [][]int) {
	// 顺时针旋转90度，可以拆解为先上下翻转，再沿从左上到右下的对角线翻转。
	// 因此分两步处理。
	edge := len(matrix)
	// 上下翻转
	// 注意取值范围是个一半的矩形，防止重复计算。
	for x := 0; x < edge/2; x++ {
		for y := 0; y < edge; y++ {
			xx := -x + edge - 1 // 上下翻转，实质是先取-x，翻转到坐标系下，再平移上来length-1的长度
			yy := y
			temp := matrix[x][y]
			matrix[x][y] = matrix[xx][yy]
			matrix[xx][yy] = temp
		}
	}

	// 对角交换
	// 注意取值范围是个一半的三角形，防止重复计算。
	for x := 0; x < edge; x++ {
		for y := 0; y <= x; y++ {
			xx := y
			yy := x
			temp := matrix[x][y]
			matrix[x][y] = matrix[xx][yy]
			matrix[xx][yy] = temp
		}
	}
}
