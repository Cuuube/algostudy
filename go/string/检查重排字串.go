package main

/**
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-permutation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func CheckPermutation(s1 string, s2 string) bool {
	dict := make(map[rune]int8)
	for _, charRune := range s1 {
		cnt, existed := dict[charRune]
		if !existed {
			cnt = 0
		}
		dict[charRune] = cnt + 1
	}

	for _, charRune := range s2 {
		cnt, existed := dict[charRune]
		if !existed || cnt <= 0 {
			return false
		}
		dict[charRune] -= 1
	}

	for _, cnt := range dict {
		if cnt != 0 {
			return false
		}
	}
	return true
}
