package main

/**
一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。



示例 1:

输入:
first = "pale"
second = "ple"
输出: True


示例 2:

输入:
first = "pales"
second = "pal"
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/one-away-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func oneEditAway(first string, second string) bool {
	firstChars := []byte(first)
	secondChars := []byte(second)
	if distance(len(firstChars), len(secondChars)) > 1 {
		return false
	}

	if len(firstChars) == len(secondChars) {
		flag := true
		for idx, char := range firstChars {
			if char != secondChars[idx] {
				if flag {
					flag = false
					continue
				} else {
					return false
				}
			}
		}
		return true
	}

	if len(firstChars) > len(secondChars) {
		temp := firstChars
		firstChars = secondChars
		secondChars = temp
	}
	offset := 0
	for idx, char := range firstChars {
		if char != secondChars[idx+offset] {
			if offset != 0 || char != secondChars[idx+1] {
				return false
			}
			offset += 1
		}
	}
	return true
	// 思路：1、长短一样的话，判断替换。遍历字符串，查看是不是只有一个字符不同
	// 2、长短不齐的话，判断增删。短的放前面，长的放后面

}

func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
