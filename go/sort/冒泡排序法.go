package main

// 冒泡排序
func BubbleSort(arr []int) []int {
	// 1、从前向后遍历，不断把大的冒泡到后面（或小的冒泡到前面）
	for i := len(arr) - 1; i > 0; i-- {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
