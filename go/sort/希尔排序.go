package main

// 希尔排序：插排优化版
func ShellSort(arr []int) []int {
	length := len(arr)
	// 核心：分批进行插入排序。让step不停/2来分批，然后每批内部执行插排，减少交换次数
	for steps := len(arr) / 2; steps > 0; steps /= 2 {
		// 实质上是个分批执行的插入排序
		for i := steps; i < length; i++ {
			// 内圈从后往已排序的地方进行检查、交换。
			for j := i; j >= steps; j -= steps {
				if arr[j] < arr[j-steps] {
					arr[j], arr[j-steps] = arr[j-steps], arr[j]
				} else {
					break
				}
			}
		}
	}
	return arr
}

// 抄的 https://www.cnblogs.com/chengxiao/p/6104371.html
func ShellSortRaw(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && arr[j] < arr[j-key] {
				arr[j], arr[j-key] = arr[j-key], arr[j]
				j = j - key
			}
		}
		key = key / 2
	}
	return arr
}
