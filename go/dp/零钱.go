package main

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

// 抄的
func waysToChange(targetValue int) int {
	dp := make([]int, targetValue+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for i := 0; i < len(coins); i++ {
		coinValue := coins[i]
		for value := 1; value <= targetValue; value++ {
			if value-coinValue >= 0 {
				dp[value] += dp[value-coinValue]
			}
		}
	}
	return dp[targetValue] % 1000000007
}

// TODO 有问题
// func waysToChange(n int) int {
// 	// 经典dp问题。
// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	dp[1] = 1

// 	i := 2
// 	for i <= n {
// 		sum := dp[i-1]
// 		if i == 5 || i == 10 || i == 25 {
// 			sum += 1
// 		}

// 		if i-5 > 0 {
// 			sum += dp[i-5]
// 		}
// 		if i-10 > 0 {
// 			sum += dp[i-10]
// 		}
// 		if i-25 > 0 {
// 			sum += dp[i-25]
// 		}
// 		dp[i] = sum
// 		i++
// 	}
// 	fmt.Println(dp)
// 	return dp[n]
// }
