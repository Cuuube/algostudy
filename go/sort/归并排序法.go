package main

// 归并排序法
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	// 永远将元素从中间分成两组，比较&合并交给其他函数
	return merge(MergeSort(arr[0:mid]), MergeSort(arr[mid:]))
}

// 负责接收两个有序数组，将两个有序数组合并为一个有序数组
func merge(arr1, arr2 []int) []int {
	ret := make([]int, 0)
	var i, j int = 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}
	// 放入剩余元素
	for i < len(arr1) {
		ret = append(ret, arr1[i])
		i++
	}
	for j < len(arr2) {
		ret = append(ret, arr2[j])
		j++
	}
	return ret
}
