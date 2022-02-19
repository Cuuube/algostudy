package main

// 选择排序法
func SelectionSort(arr []int) []int {
	// 不断找最小值，放到最左侧
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
