package main

/**
面试题 08.09. 括号
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
https://leetcode-cn.com/problems/bracket-lcci/
*/

// 4层for。。感觉可以优化
func generateParenthesis(n int) []string {
	raw := n
	// 根据n来递进，每次都用上一次的结果来插
	ret := map[string]bool{}
	for n > 0 {
		if n == raw {
			ret["()"] = true
			n--
			continue
		}
		newRet := map[string]bool{}
		// 从ret的结果插入
		for val := range ret {
			length := len(val)
			// 对每个字符串都确定左弧和右弧的位置
			for i := 0; i < length; i++ {
				for j := i; j < length; j++ {
					// 把左弧右弧插进字符串
					newStr := val[0:i] + "(" + val[i:j] + ")" + val[j:]
					newRet[newStr] = true
				}
			}

		}
		ret = newRet
		n--
	}
	retArr := make([]string, 0)
	for str := range ret {
		retArr = append(retArr, str)
	}
	return retArr
}
