package main

import "strings"

/*
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False
提示：

字符串长度在[0, 100000]范围内。
说明:

你能只调用一次检查子串的方法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-rotation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isFlipedString(s1 string, s2 string) bool {
	// 轮转字符串的特征是一个字符串通过一次分割重组，就可以组成另一个字符串。
	// 因为是一次分割重组，那么相同的字串首尾相连，必定包含旋转后的字符串
	// abcxyz => cxyzab；重复cxyzab为「cxyzabcxyzav」，里面包含目标字串
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s2+s2, s1)
}

func isFlipedStringBad(s1 string, s2 string) bool {
	// 思路：字符串旋转有几个特征：
	// 1、长度必须完全相同
	// 2、所有包含的字母相同
	// 如果只包含asicc，用[]byte，否则用[]rune
	chars1 := []byte(s1)
	chars2 := []byte(s2)

	if len(chars1) != len(chars2) {
		return false
	}

	cntMap := make(map[byte]int)
	for _, char := range chars1 {
		cnt, existed := cntMap[char]
		if !existed {
			cnt = 0
		}
		cntMap[char] = cnt + 1
	}

	// 检查第二个字串
	for _, char := range chars2 {
		cnt, existed := cntMap[char]
		if !existed {
			return false
		}
		cntMap[char] = cnt - 1
		if cntMap[char] == 0 {
			delete(cntMap, char)
			continue
		}
	}

	return len(cntMap) == 0
}
