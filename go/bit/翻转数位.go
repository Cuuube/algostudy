package main

import "fmt"

/**
面试题 05.03. 翻转数位
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

示例 1：

输入: num = 1775(110111011112)
输出: 8
示例 2：

输入: num = 7(01112)
输出: 4

*/

func testReverseBits() {
	fmt.Println(reverseBits(0))
	fmt.Println(reverseBits(-1))
	fmt.Println(reverseBits(45725232))
}

func reverseBits(num int) int {
	// 转变为相邻元素的和:
	// 遍历数组，记录从1开始到1结束的长度
	arr := make([]int, 0)
	cnt := 0
	// going := false
	for i := 0; i < 32; i++ {
		bitIsTrue := num|(1<<i) == num
		if bitIsTrue {
			cnt += 1
			continue
		}
		arr = append(arr, cnt)
		cnt = 0
	}
	if cnt != 0 {
		arr = append(arr, cnt)
	}

	// 检查数组相邻最大值
	max := 0
	for i, v := range arr {
		sum := 0
		if i == len(arr)-1 {
			sum = v
		} else {
			sum = v + arr[i+1]
		}
		if sum > max {
			max = sum
		}
	}
	// 边界检查
	if max >= 32 {
		return max
	}
	return max + 1
}

func reverseBitsBad(num int) int {
	// 找到最多1的联排数
	// 遍历bits，挨个变成1
	max := 0
	for i := 0; i < 32; i++ {
		changed := fillBit(num, i)
		if changed == num {
			continue
		}
		cnt := findMaxCnt(changed)
		if cnt > max {
			max = cnt
		}
	}
	return max
}

func findMaxCnt(num int) int {
	max := 0
	cnt := 0
	flag := false
	for i := 0; i < 32; i++ {
		if fillBit(num, i) == num {
			if flag {
				cnt++
				if cnt > max {
					max = cnt
				}
			} else {
				flag = true
				cnt = 1
				if cnt > max {
					max = cnt
				}
			}
		} else {
			cnt = 0
			flag = false
		}
	}
	return max
}
