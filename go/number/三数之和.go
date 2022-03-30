package main

import "sort"

/*

15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105



https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/

*/

// 排序+左右双指针
func SumThreeNumberEqualZero(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		right := n - 1
		// 枚举 b: [first, n]
		for left := first + 1; left < n; left++ {
			// 需要和上一次枚举的数不相同
			if left > first+1 && nums[left] == nums[left-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧。枚举right：[left, n]，并且保证不能过小
			for left < right && nums[first]+nums[left]+nums[right] > 0 {
				right--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if left == right {
				break
			}
			if nums[first]+nums[left]+nums[right] == 0 {
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
			}
		}
	}
	return ans
}

// 有bug
func SumThreeNumberEqualZeroBad(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	ret := make([][]int, 0)

	for start := 0; start < len(nums)-2; start++ {
		if start-1 >= 0 && sortedNums[start-1] == sortedNums[start] {
			continue
		}

		// 将剩余的范围使用左右双指针，向中间移动
		left := start + 1
		right := len(nums) - 1

	LOOP:
		for left+1 < right && sortedNums[left] == sortedNums[left+1] {
			left++
		}
		for right-1 > left && sortedNums[right] == sortedNums[right-1] {
			right--
		}
		if left >= right {
			continue
		}
		sum := sortedNums[start] + sortedNums[left] + sortedNums[right]
		if sum == 0 {
			// append
			ret = append(ret, []int{sortedNums[start], sortedNums[left], sortedNums[right]})
		} else if sum > 0 {
			right--
			goto LOOP
		} else {
			left++
			goto LOOP
		}
	}
	return ret
}
