package main

// 快速排序法
func QuickSort(arr []int) []int {
	length := len(arr)
	if length == 0 || length == 1 {
		return arr
	} else if length == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}

	mid := length / 2
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < length; i++ {
		v := arr[i]
		if i == mid {
			continue
		}
		if v <= arr[mid] {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	ret := QuickSort(left)
	ret = append(ret, arr[mid])
	ret = append(ret, QuickSort(right)...)
	return ret
}
