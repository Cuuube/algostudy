package main

/**
给定一组数字，寻找任意相邻数字串中，和最大的结果。
例如：input：[2,-3,4,5,-2]；output：9
需求：使用时间复杂度O(n)的算法。
*/

func findMaxSlice(arr []int) int {
	max := 0
	// 1、相邻和最大，最简单的思路：需要知道每一个子串的和，然后比较大小。但这样一来，复杂度是O(n2)。有什么简化方式呢？
	// 2、我们把所有切片拿出来，会发现，如果开头是负数的，或切片前几个和为负数，那可以直接去掉。去掉负数区间，有助于增大总值。
	// 3、那么发现，如果我们直接用原数组来检查呢？从头遍历，如果遇到前n个相加为负数，就直接舍弃这n个结果。
	// 4、每一步都记录当前切片的value，不断和历史max比较。更大则更新max。如此，一直到结尾，就可以算出max的值了。
	current := 0
	for _, v := range arr {
		current += v
		if current < 0 {
			current = 0
			continue
		}
		if current > max {
			max = current
		}
	}
	return max
}
