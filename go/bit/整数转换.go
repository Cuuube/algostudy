package main

/**
面试题 05.06. 整数转换
整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。

示例1:

 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
 输出：2
示例2:

 输入：A = 1，B = 2
 输出：2
提示:

A，B范围在[-2147483648, 2147483647]之间
*/

func convertInteger(A int, B int) int {
	// 本质是判断两个数里有几个位是不同的
	cnt := 0
	for i := 0; i < 32; i++ {
		mask := 1 << i
		isTrueA := mask|A == A
		isTrueB := mask|B == B
		if isTrueA != isTrueB {
			cnt += 1
		}
	}
	return cnt
}
