package main

/**
回文排列
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func canPermutePalindrome(s string) bool {
	dict := make(map[rune]int)
	for _, r := range s {
		cnt, existed := dict[r]
		if !existed {
			cnt = 0
		}
		dict[r] = cnt + 1
	}

	// 检查dict
	singleExisted := false
	for _, cnt := range dict {
		if cnt%2 == 1 {
			if !singleExisted {
				singleExisted = true
				continue
			} else {
				return false
			}
		}
	}
	return true
}
