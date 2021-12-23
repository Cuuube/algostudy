package main

import "fmt"

func testsubsets() {
	a := []int{9, 0, 3, 5, 7}
	fmt.Printf("%+v", subsets(a))
}

func subsets(nums []int) [][]int {
	// 子集问题：遍历数组：每个新元素加入，都执行这两步：
	// 1、复制一遍旧数组，把自己加到旧数组中去
	// 2、 把复制组合的这些新数组加到大数组列表里
	// 因为每个新元素加入，子列表数量都翻倍。也就是说，最终长度会是2的length次方个元素
	ret := make([][]int, 0)
	ret = append(ret, make([]int, 0))
	for _, v := range nums {
		for _, arr := range ret {
			ar := copyArr(arr)
			ar = append(ar, v)
			ret = append(ret, ar)
		}
	}
	return ret
}
func copyArr(arr []int) []int {
	ar := make([]int, 0)
	ar = append(ar, arr...)
	return ar
}

func subsetsError(nums []int) [][]int {
	// 子集问题：遍历数组：每个新元素加入，都执行这两步：
	// 1、复制一遍旧数组，把自己加到旧数组中去
	// 2、 把复制组合的这些新数组加到大数组列表里
	// 因为每个新元素加入，子列表数量都翻倍。也就是说，最终长度会是2的length次方个元素
	ret := make([][]int, 0)
	ret = append(ret, make([]int, 0))
	for _, v := range nums {
		for _, arr := range ret {
			// TODO 案发现场。golang的range数组坑。这里可能会报错
			// 原因：range ret时，ret虽然是一个快照，但是arr是实打实的指针。所以无论是推入v，还是把arr存入ret，都可能会被改变。
			// https://juejin.cn/post/6844903773559619591
			arr = append(arr, v)
			ret = append(ret, arr)
		}
	}
	return ret
}
