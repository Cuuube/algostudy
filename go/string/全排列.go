package main

import (
	"fmt"
	"strings"
)

func testpermutation() {
	// fmt.Println(permutationN(5))
	// fmt.Println(permutationNM(1, 5))
	// fmt.Println(permutation("abb"))
	// fmt.Printf("%+v\n", cache)
	// fmt.Println(permutationWithSameChar("abbcbcaa"))
	// fmt.Printf("%+v\n", uniquecCache)
	// fmt.Println(arrangementBySubset("abbcbcaa"))
	fmt.Println(arrangementByDp("abbc"))

}

/**
面试题 08.07. 无重复字符串的排列组合
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

示例1:

 输入：S = "qwe"
 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
示例2:

 输入：S = "ab"
 输出：["ab", "ba"]
提示:

字符都是英文字母。
字符串长度在[1, 9]之间。
https://leetcode-cn.com/problems/permutation-i-lcci/submissions/
*/
func permutation(S string) []string {
	// 不仅要排列数量，还要排列结果。可以在递归的时候，把结果加进去
	// fn(currentChar byte, otherChars []byte, result *[][]byte)
	// ret := make([]string, 0)

	// chars := []byte(S)
	// permutationResults := fn(chars)
	// for _, resultChars := range permutationResults {
	// 	ret = append(ret, string(resultChars))
	// }
	// return ret
	return cachedDfs(S)
}

func permutationWithSameChar(S string) []string {
	// 对包含重复数据的全排列
	ret := make([]string, 0)
	result := cachedDfsDeup(S)
	for str := range result {
		ret = append(ret, str)
	}
	return ret
}

// 获取字符全排列的结果
func fn(chars []byte) [][]byte {
	ret := make([][]byte, 0)
	if len(chars) <= 1 {
		// 递归终点，只有一个元素，就返回自己
		currentChars := append(make([]byte, 0), chars[0])
		ret = append(ret, currentChars)
		return ret
	}
	// 多个字符，就遍历+递归，获取每个剩余字符打头的全排列
	for i, char := range chars {
		// 组成char开头的各种字符串

		// 把获取剩余字符
		otherChars := make([]byte, 0)
		otherChars = append(otherChars, chars[0:i]...)
		if i != len(chars)-1 {
			otherChars = append(otherChars, chars[i+1:]...)
		}

		// 把剩余字符全排列
		permutationResults := fn(otherChars)
		for _, resultClips := range permutationResults {
			// 把自己加到剩余字符全排列的前头
			fullClips := append(make([]byte, 0), char)
			fullClips = append(fullClips, resultClips...)
			ret = append(ret, fullClips)
		}
	}
	return ret
}

// 实质上，用的方法是dfs+记忆化。可以看做一棵树，每个节点都是当前char，下个节点是剩余的char。以深度优先遍历。
// 优化1：获取字符串全排列的结果，是上面那个char的优化版。，时间空间减少20%
// 优化2：缓存全排列路由表。时间减少50%，空间增加30%
// 思考：如果用动态规划，会是什么样的？
var cache = make(map[string][]string)

func cachedDfs(str string) []string {
	// cache
	if ret, found := cache[str]; found {
		return ret
	}

	ret := make([]string, 0)
	if len(str) <= 1 {
		// 递归终点，只有一个元素，就返回自己
		ret = append(ret, str)
		cache[str] = ret
		return ret
	}
	// 多个字符，就遍历+递归，获取每个剩余字符打头的全排列
	for i := 0; i < len(str); i++ {
		currentChar := str[i : i+1]
		otherStr := strings.Replace(str, currentChar, "", 1)
		permutationResults := cachedDfs(otherStr)
		for _, result := range permutationResults {
			ret = append(ret, currentChar+result)
		}

	}
	cache[str] = ret
	return ret
}

// dp做效率不高
func arrangementByDp(str string) map[string]bool {
	// dp做全排列，如何构建dp表？递进关系是什么？
	// 递进关系：旧字符和新加字符。dp表idx：当前所有字符。特殊值：1个字符直接存。执行函数：字符和老结果交混

	// 把char混到list字符串里的任意位置
	mixin := func(list map[string]bool, char string) map[string]bool {
		ret := make(map[string]bool, 0)
		for subStr := range list {
			// 将currentChar插入subStr任意位置
			i := 0
			for i <= len(subStr) {
				newStr := subStr[0:i] + char + subStr[i:]
				ret[newStr] = true
				i++
			}
		}
		return ret
	}

	end := 1
	dp := make(map[string]map[string]bool, 0)
	for end <= len(str) {
		currentStr := str[0:end]
		// 当只有一位的时候，直接保存结果
		if len(currentStr) == 1 {
			dp[currentStr] = map[string]bool{currentStr: true}
			end += 1
			continue
		}
		newChar := str[end-1 : end]
		before := str[0 : end-1]
		// dp[n] = dp[n-1] * n
		dp[currentStr] = mixin(dp[before], newChar) // 拉出一个算子
		end += 1
	}
	return dp[str]
}

// 占用内存和执行时间都要更好一些
func arrangementBySubset(str string) []string {
	// 用子集的思路来写：
	// 1、遍历字符串
	// 2、每获取一个新字符，都遍历之前的结果，并以之前结果的任意位置来插入当前字符
	// 3、注意，可能有大量重复，注意去重
	// 执行速度O(n3)
	ret := make(map[string]bool, 0)
	end := 1
	for end <= len(str) {
		currentChar := str[end-1 : end]
		end += 1
		if len(ret) == 0 {
			ret[currentChar] = true
			continue
		}

		newRet := make(map[string]bool, 0)
		// 遍历之前结果
		for subStr := range ret {
			// 将currentChar插入subStr任意位置
			i := 0
			for i <= len(subStr) {
				newStr := subStr[0:i] + currentChar + subStr[i:]
				newRet[newStr] = true
				i++
			}
		}
		ret = newRet
	}
	retArr := make([]string, 0)
	for str := range ret {
		retArr = append(retArr, str)
	}
	return retArr
}

// 带上去重效果的，包含相同元素的全排列
var uniquecCache = make(map[string]map[string]bool)

func cachedDfsDeup(str string) map[string]bool {
	// cache
	if ret, found := uniquecCache[str]; found {
		return ret
	}

	ret := make(map[string]bool, 0)
	if len(str) <= 1 {
		// 递归终点，只有一个元素，就返回自己
		uniquecCache[str] = ret
		ret[str] = true
		return ret
	}
	// 多个字符，就遍历+递归，获取每个剩余字符打头的全排列
	for i := 0; i < len(str); i++ {
		currentChar := str[i : i+1]
		otherStr := strings.Replace(str, currentChar, "", 1)
		permutationResults := cachedDfsDeup(otherStr)
		for result := range permutationResults {
			ret[currentChar+result] = true
		}

	}
	uniquecCache[str] = ret
	return ret
}

// 问：n个字符全排列，有几种排列方式？
func permutationN(n int) int {
	// 边界检查
	// n的全排列，是自己的阶乘n!
	if n <= 1 {
		return n
	}
	return n * permutationN(n-1)
}

// 问：n个字符，排列m长度，全排列，有几种排列方式？
func permutationNM(n, m int) int {
	// 思路，m的加入，相当于在n的阶乘操作中，乘m次。可能提前截胡，所以要让m递减，如果m为0，就提前终止
	if n == 0 || m == 0 {
		return 0
	}
	if n <= 1 || m <= 1 {
		return n
	}
	return n * permutationNM(n-1, m-1)
}
