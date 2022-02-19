package main

func HeapSort(arr []int) []int {
	// 执行堆化
	arr = buildMaxHeap(arr, 0, len(arr))

	for i := 0; i < len(arr)-1; i++ {
		end := len(arr) - 1 - i
		// 交换头尾
		arr[0], arr[end] = arr[end], arr[0]

		// 将排序之外的元素继续堆化
		arr = buildMaxHeap(arr, 0, end)
	}
	return arr
}

// 构建大顶堆
func buildMaxHeap(arr []int, start, end int) []int {
	// 遍历每一个父节点，保证父节点比子节点大
	for i := end/2 - 1; i >= 0; i-- { // 父节点为[0, end/2-1]
		leftChild := i*2 + 1  // 左子节点
		rightChild := i*2 + 2 // 右子节点
		// 如果左右子节点合法，且大于父节点，就调换两者位置
		if rightChild < end && arr[rightChild] > arr[i] {
			arr[rightChild], arr[i] = arr[i], arr[rightChild]
		}
		if leftChild < end && arr[leftChild] > arr[i] {
			arr[leftChild], arr[i] = arr[i], arr[leftChild]
		}
	}
	return arr
}
