package main

import (
	"fmt"
)

/**
面试题 05.08. 绘制直线
绘制直线。有个单色屏幕存储在一个一维数组中，使得32个连续像素可以存放在一个 int 里。屏幕宽度为w，且w可被32整除（即一个 int 不会分布在两行上），屏幕高度可由数组长度及屏幕宽度推算得出。请实现一个函数，绘制从点(x1, y)到点(x2, y)的水平线。

给出数组的长度 length，宽度 w（以比特为单位）、直线开始位置 x1（比特为单位）、直线结束位置 x2（比特为单位）、直线所在行数 y。返回绘制过后的数组。

示例1:

 输入：length = 1, w = 32, x1 = 30, x2 = 31, y = 0
 输出：[3]
 说明：在第0行的第30位到第31为画一条直线，屏幕表示为[0b000000000000000000000000000000011]
示例2:

 输入：length = 3, w = 96, x1 = 0, x2 = 95, y = 0
 输出：[-1, -1, -1]

 https://leetcode-cn.com/problems/draw-line-lcci/
*/

func testdrawLine() {
	fmt.Println(drawLine(1, 32, 30, 31, 0))
	fmt.Println(drawLine(3, 96, 0, 95, 0))
}

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	// 求出x1,y，x2，y落在哪个int
	c := 0
	temp := 0
	ret := make([]int, 0)
	for c < (length * 32) {
		offset := c % 32

		x0 := c % w
		y0 := c / w
		if y0 == y && x0 >= x1 && x0 <= x2 {
			temp = 1<<(31-offset) | temp
		}
		if (c+1)%32 == 0 {
			ret = append(ret, temp)
			// ret = append(ret, int(int32(temp))) // 这里丑陋的转换，强行把最大值溢出为-1
			temp = 0
		}
		c += 1
	}

	return ret
}

// func drawLine(length int, w int, x1 int, x2 int, y int) []int32 {
// 	// wBits := w/32
// 	// h := length/wBits

// 	// 求出x1,y，x2，y落在哪个int
// 	c := 0
// 	var temp int32 = 0
// 	ret := make([]int32, 0)
// 	for c < (length * 32) {
// 		offset := c % 32

// 		x0 := c % w
// 		y0 := c / w
// 		if y0 == y && x0 >= x1 && x0 <= x2 {
// 			temp = 1<<(31-offset) | temp
// 		}
// 		if (c+1)%32 == 0 {
// 			ret = append(ret, temp)
// 			temp = 0
// 		}
// 		c += 1
// 	}

// 	return ret
// }
