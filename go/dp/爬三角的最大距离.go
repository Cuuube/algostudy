package main

import (
	"algostudy/go/util"
	"fmt"
)

// 动态规划，爬三角。三角从上到下，每一步只能向左或者向右。问到终点时的最大数字
func triangle() {
	tri := [][]int{
		{7},
		{3, 8},
		{8, 1, 0},
		{2, 7, 4, 4},
		{4, 5, 2, 6, 5},
	}

	resultMap := [5][5]int{}
	for n, line := range tri {
		if n == 0 {
			resultMap[0][0] = 7
		}
		for x, v := range line {
			left := 0
			right := 0
			if x > 0 {
				left = resultMap[n-1][x-1]
			}
			if x < n {
				right = resultMap[n-1][x]
			}
			resultMap[n][x] = util.Max(left+v, right+v)
		}
	}

	fmt.Printf("%+v", resultMap)
}
