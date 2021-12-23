package main

import "fmt"

/**
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


*/
func climbStairs() {
	// 爬楼梯，可以选择走一步，或者走两步。
	// 近似的，可以看作一棵二叉树。向左是走一步，向右是走两步。
	// 接到一个终点值，可以遍历这棵二叉树，直到终点距离n为止，获取次数。
	// 但是这种解法效率很低，是2的n方复杂度。
	// 仔细看树的形状，可以发现这棵树的局部和整体是自相似的。随着n的增加，树不断向上增长，实质长成了f(n-1)与f(n-2)的组合。也就是说，可以得出规律：f(n) = f(n-1) + f(n-2)

	fmt.Println(solve(8))
	fmt.Println(solve(9))
	fmt.Println(solve(50))
}

func solve(target int) int {
	dp := make([]int, 0)

	for i := 0; i <= target+1; i++ {
		switch i {
		case 0:
			dp = append(dp, 0)
		case 1:
			dp = append(dp, 1)
		case 2:
			dp = append(dp, 2)
		default:
			dp = append(dp, dp[i-1]+dp[i-2])
		}
	}

	return dp[target]
}

/**
面试题 08.01. 三步问题
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:

 输入：n = 3
 输出：4
 说明: 有四种走法
示例2:

 输入：n = 5
 输出：13
提示:

n范围在[1, 1000000]之间
https://leetcode-cn.com/problems/three-steps-problem-lcci/
*/
func waysToStep(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			dp[i] = 0
		case 1:
			dp[i] = 1
		case 2:
			dp[i] = 2
		case 3:
			dp[i] = 4
		default:
			dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
		}
	}
	return dp[n]
}
