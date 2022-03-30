package main

import (
	"strconv"
)

/**
670. 最大交换
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 108]



*/

// 思路：先把数字专为string，方便遍历
// 如果要交换，肯定是最左边小的，希望和右边最大的进行交换，得到最高效率。
// 关键点：如何找到每一位自己右边最大的？可以从右边遍历，找自己右边历史最大极值
func MaxSwap(number int) int {
	strArr := strconv.Itoa(number)
	numberArr := make([]int, 0)
	for idx := range strArr {
		char := strArr[idx : idx+1]
		num, _ := strconv.Atoi(char)
		numberArr = append(numberArr, num)
	}
	// 构造最大表：自己idx：自己右边最大值的idx
	idxToMaxVal := make(map[int]int, 0)
	maxValKey := len(numberArr) - 1
	for idx := len(numberArr) - 1; idx >= 0; idx-- {
		num := numberArr[idx]
		if num > numberArr[maxValKey] {
			maxValKey = idx
		}
		idxToMaxVal[idx] = maxValKey
	}
	// 进行交换检查，从高位向低位找，找到交换最优的
	for idx := 0; idx < len(numberArr); idx++ {
		num := numberArr[idx]
		targetMaxIdx := idxToMaxVal[idx]
		targetMaxNum := numberArr[targetMaxIdx]
		if targetMaxNum > num {
			// 执行交换
			numberArr[idx], numberArr[targetMaxIdx] = numberArr[targetMaxIdx], numberArr[idx]
			return createNumber(numberArr)
		}
	}
	return number
}

func createNumber(arr []int) int {
	str := ""
	for idx := range arr {
		str += strconv.Itoa(arr[idx])
	}
	ret, _ := strconv.Atoi(str)
	return ret
}
