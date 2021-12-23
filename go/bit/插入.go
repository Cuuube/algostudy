package main

/**
面试题 05.01. 插入
给定两个整型数字 N 与 M，以及表示比特位置的 i 与 j（i <= j，且从 0 位开始计算）。

编写一种方法，使 M 对应的二进制数字插入 N 对应的二进制数字的第 i ~ j 位区域，不足之处用 0 补齐。具体插入过程如图所示。



题目保证从 i 位到 j 位足以容纳 M， 例如： M = 10011，则 i～j 区域至少可容纳 5 位。



示例1:

 输入：N = 1024(10000000000), M = 19(10011), i = 2, j = 6
 输出：N = 1100(10001001100)
示例2:

 输入： N = 0, M = 31(11111), i = 0, j = 4
 输出：N = 31(11111)
https://leetcode-cn.com/problems/insert-into-bits-lcci/
*/
func insertBits1(N int, M int, i int, j int) int {
	// 第一种方法、先且上0来清空指定位置
	// 再把目标字串插入指定位置
	for idx := i; idx <= j; idx++ {
		N = clearBit(N, idx)
	}
	return N + (M << i)
}

func insertBits2(N int, M int, i int, j int) int {
	// 第二种方法、把字符串分割为左边，右边和中间
	leftMoveSteps := j + 1
	left := N >> leftMoveSteps << leftMoveSteps // 利用先右移再左移，来清空右侧数据，从而获得左侧数据

	rightMoveSteps := i
	right := N - (N >> i << i) // 利用两者的差，来获得右侧数据

	insertNum := M << rightMoveSteps
	return left + right + insertNum
}
