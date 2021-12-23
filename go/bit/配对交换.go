package main

/**
面试题 05.07. 配对交换
配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。

示例1:

 输入：num = 2（或者0b10）
 输出 1 (或者 0b01)
示例2:

 输入：num = 3
 输出：3
提示:

num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。

https://leetcode-cn.com/problems/exchange-lcci/
*/

func exchangeBits(num int) int {
	// 检测奇数位和偶数位是否相同，相同则不变，不相同则各自取反
	for i := 0; i < 32; i += 2 {
		rightIsTrue := 1<<i|num == num
		leftIsTrue := 1<<(i+1)|num == num
		if leftIsTrue != rightIsTrue {
			num = reversedBit(num, i)
			num = reversedBit(num, i+1)
		}
	}
	return num
}
