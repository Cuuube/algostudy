package main

/**
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-unique-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isUniqueByBitMap(astr string) bool {
	var bitmap int
	for _, r := range astr {
		char := int(r)
		var exec int = 1 << (char - 97)
		if bitmap|exec == bitmap {
			return false
		}
		bitmap = bitmap | exec
	}
	return true
}

// func isUniqueBad(astr string) bool {
// 	uniqueMap := map[string]bool{}
// 	for _, val := range astr {
// 		char := fmt.Sprintf("%c", val)
// 		if _, exist := uniqueMap[char]; exist {
// 			return false
// 		}
// 		uniqueMap[char] = true
// 	}
// 	return true
// }
