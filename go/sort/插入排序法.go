package main

// 插入排序法
func InsertionSort(arr []int) []int {
	// 对待排序的执行一次遍历，已排序的执行n次排序
	// 从第二个开始遍历
	for i := 1; i < len(arr); i++ {
		// current := arr[i]

		// 遍历已排序的
		for position := i; position >= 1; position-- {
			// 如果比之前的小，就执行交换
			if arr[position] < arr[position-1] {
				arr[position], arr[position-1] = arr[position-1], arr[position]
			} else {
				break
			}
		}
	}
	return arr
}
