package main

import "fmt"

/**
面试题 08.11. 硬币
硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)

示例1:

 输入: n = 5
 输出：2
 解释: 有两种方式可以凑成总金额:
5=5
5=1+1+1+1+1
示例2:

 输入: n = 10
 输出：4
 解释: 有四种方式可以凑成总金额:
10=10
10=5+5
10=5+1+1+1+1+1
10=1+1+1+1+1+1+1+1+1+1
说明：

注意:

你可以假设：

0 <= n (总金额) <= 1000000
*/

/*
	正确解法：
	dfs N叉树时，屏蔽掉重复路径。
	或者构建dp表时，按面值，分四种情况来更新dp表：只有1块，有1块和5块，有1块和5块和10块、有1块和5块和10块和25块，，
*/
func waysToChange(targetValue int) int {
	dp := make([]int, targetValue+1)
	dp[0] = 1 // 无意义，视为value正好和面值相同的情况

	coins := []int{1, 5, 10, 25}

	// 分面值，每增加一种面值，只能用该面值和比这个面值小的币种。
	for _, coinValue := range coins {

		// 每增加一种面值，都更新一遍dp表
		i := 1
		for i <= targetValue {
			if i-coinValue >= 0 {
				dp[i] += dp[i-coinValue]
			}
			i++
		}
	}

	fmt.Println(dp)
	return dp[targetValue] % 1000000007
}

/*
	踩入陷阱的解法。错误原因：
	对于付钱方式而言，付6块，先付1再付5，和先付5再付1是一样的。所以这里会重复计算。
*/
func waysToChangeError(targetValue int) int {
	// 经典dp问题，用经典dp解法试试
	dp := make([]int, targetValue+1)
	dp[0] = 1

	coins := []int{1, 5, 10, 25}

	i := 1
	for i <= targetValue {
		for _, coinValue := range coins {
			if i-coinValue >= 0 {
				dp[i] += dp[i-coinValue]
			}
		}
		i++
	}
	fmt.Println(dp)
	return dp[targetValue] % 1000000007
}

// 引用答案：
// func waysToChange(targetValue int) int {
// dp := make([]int, targetValue+1)
// dp[0] = 1
// coins := []int{1, 5, 10, 25}
// for i := 0; i < len(coins); i++ {
// 	coinValue := coins[i]
// 	for value := 1; value <= targetValue; value++ {
// 		if value-coinValue >= 0 {
// 			dp[value] += dp[value-coinValue]
// 		}
// 	}
// }
// }
