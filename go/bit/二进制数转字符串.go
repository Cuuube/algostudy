package main

import "fmt"

/**
面试题 05.02. 二进制数转字符串
二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。

示例1:

 输入：0.625
 输出："0.101"
示例2:

 输入：0.1
 输出："ERROR"
 提示：0.1无法被二进制准确表示
提示：

32位包括输出中的"0."这两位。
*/

func testPrintBin() {
	fmt.Println(printBin(0.625))
}

func printBin(num float64) string {
	// 浮点数转二进制，根据浮点数在二进制的储存定义，第10位bit表示1，在之后每个bit，都表示前一bit/2
	// 也就是在小数位，第一位表示0.5，第二位表示0.25，第三位表示0.125。。。。。。。。
	// 浮点的无限除法细分，可以视为无限的乘法取余
	// 最后，题目要求如果32-2个字符装不下，就返回error
	ret := "0."
	cnt := 32 - 2
	for cnt > 0 && num != 0 {
		num = num * 2
		if num >= 1 {
			num -= 1
			ret += "1"
		} else {
			ret += "0"
		}
		cnt -= 1
	}
	if num == 0 {
		return ret
	}
	return "ERROR"
}
