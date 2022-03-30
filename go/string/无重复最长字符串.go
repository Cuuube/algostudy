package main

/*

3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetc-2/
*/

// 滑动窗口法
func LengthOfLongestSubstring(s string) int {
	set := map[byte]bool{}
	chars := []byte(s)
	right := 0
	max := 0
	for left := 0; left < len(chars); left++ {
		if left != 0 {
			delete(set, chars[left-1]) // 每次进入，移除之前的一位
		}

		// 右侧移动
		for right < len(chars) && !set[chars[right]] {
			set[chars[right]] = true
			right++
		}
		// 更新max
		if right-left > max {
			max = right - left
		}
	}
	return max
}

// func LengthOfLongestSubstring(s string) int {
// 	// 移动右指针，如果依然不存在重复，则继续移动。如果存在重复，移动左指针
// 	chars := []byte(s)
// 	set := map[byte]int{}
// 	right := 0
// 	max := 0
// 	for left := 0; left < len(chars); left++ {
// 		if left != 0 {
// 			delete(set, chars[left-1])
// 		}
// 		// 右指针不断向右移动，直到遇到重复的
// 		for right < len(chars) && set[chars[right]] == 0 {
// 			set[chars[right]]++ // 放入set
// 			right++             // 右指针右移
// 		}

// 		// 替换掉最长的结果
// 		if right-left > max {
// 			max = right - left
// 		}
// 	}

// 	return max
// }

// O(N^2)会超时，滑动遍历，使用set判断是否已经存在新字符串
func lengthOfLongestSubstringBad(s string) int {
	ret := ""
	for head := 0; head < len(s)-len(ret); head++ {
		clips := ""
		set := map[string]struct{}{}
		for idx := head; idx < len(s); idx++ {
			c := s[idx : idx+1]
			if _, existed := set[c]; existed {
				clips = ""
				set = map[string]struct{}{}
				continue
			}
			set[c] = struct{}{}
			clips += c
			if len(clips) > len(ret) {
				ret = clips
			}
		}
	}
	return len(ret)
}
