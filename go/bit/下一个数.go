package main

import (
	"fmt"
)

/**
面试题 05.04. 下一个数
下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。

示例1:

 输入：num = 2（或者0b10）
 输出：[4, 1] 或者（[0b100, 0b1]）
示例2:

 输入：num = 1
 输出：[2, -1]
提示:

num的范围在[1, 2147483647]之间；
如果找不到前一个或者后一个满足条件的正数，那么输出 -1。
https://leetcode-cn.com/problems/closed-number-lcci/
*/

func testfindClosedNumbers() {
	// fmt.Println(findClosedNumbers(34))
	fmt.Println(findClosedNumbers(67))
}

func findClosedNumbers(num int) []int {
	// TODO有问题！！！
	if num == 0xfffffff {
		return []int{-1, -1}
	}
	// 暴力枚举吧
	big := -1
	small := -1
	cnt := getBitTrueCnt(num)
	cur := num
	for cur < 0xfffffff {
		cur += 1
		currentCnt := getBitTrueCnt(cur)
		if currentCnt == cnt {
			big = cur
			break
		}
	}
	cur = num
	for cur >= 0 {
		cur -= 1
		currentCnt := getBitTrueCnt(cur)
		if currentCnt == cnt {
			small = cur
			break
		}
	}
	return []int{big, small}
}

func findClosedNumbersError(num int) []int {
	// 思路有问题！！
	// 同样1的数量，找到更大和更小的。
	// 更大的：把最右边的1尽可能往左移一位
	// 更小的：把最右边的1尽可能往右移一位
	// num为32位

	const length = 32
	big := -1
	small := -1
	// 找到最右侧的1，翻转为0.像左遍历：若遇到0，翻转。
	reversed := false
	temp := num
	for i := 0; i < length; i++ {
		turnTrue := 1<<i | temp
		bitIsTrue := turnTrue == temp
		if !reversed && bitIsTrue {
			temp = reversedBit(temp, i)
			reversed = true
			continue
		}
		turnFalse := ^(1 << i) & temp
		bitIsFalse := turnFalse == temp
		if bitIsFalse && reversed {
			temp = reversedBit(temp, i)
			reversed = false
			break
		}
	}
	if !reversed && temp != num {
		big = temp
	}

	// 找到最右侧的0，翻转为1.像左遍历：若遇到1，翻转。
	reversed = false
	temp = num
	for i := 0; i < length; i++ {
		turnFalse := ^(1 << i) & temp
		bitIsFalse := turnFalse == temp
		if !reversed && bitIsFalse {
			temp = reversedBit(temp, i)
			reversed = true
			continue
		}
		turnTrue := 1<<i | temp
		bitIsTrue := turnTrue == temp
		if bitIsTrue && reversed {
			temp = reversedBit(temp, i)
			reversed = false
			break
		}
	}
	if !reversed && temp != num {
		small = temp
	}
	return []int{big, small}
}
