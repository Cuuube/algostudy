package main

import "fmt"

/**
面试题 08.06. 汉诺塔问题
在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
(1) 每次只能移动一个盘子;
(2) 盘子只能从柱子顶端滑出移到下一根柱子;
(3) 盘子只能叠在比它大的盘子上。

请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。

你需要原地修改栈。

示例1:

 输入：A = [2, 1, 0], B = [], C = []
 输出：C = [2, 1, 0]
示例2:

 输入：A = [1, 0], B = [], C = []
 输出：C = [1, 0]
提示:

A中盘子的数目不大于14个。
https://leetcode-cn.com/problems/hanota-lcci/
*/

func testhanota() {
	fmt.Println(hanota([]int{2, 1, 0}, []int{}, []int{}))
}

func hanota(A []int, B []int, C []int) []int {
	hannotaMove(&A, &C, &B, len(A))
	return C
}

/**
对于长度为n的汉诺塔，每一步，都是把n-1层先挪到中间层，然后底层放到目标层，再把中间层的挪回去的过程。
*/
func hannotaMove(source *[]int, target *[]int, third *[]int, n int) {
	if n == 1 {
		// 剩一个了，直接移动
		moveOne(source, target)
		return
	}

	// 先把顶层放到第三根上
	hannotaMove(source, third, target, n-1)
	// 把source放到target上
	moveOne(source, target)
	// 再把third的放到target上
	hannotaMove(third, target, source, n-1)
}

// 移动堆顶的1个盘子
func moveOne(source *[]int, target *[]int) {
	*target = append(*target, (*source)[len(*source)-1])
	*source = (*source)[0 : len(*source)-1]
}
