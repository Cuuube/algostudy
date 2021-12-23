package main

import "fmt"

func testmultiply() {
	// fmt.Println(multiply(1, 10))
	fmt.Println(multiply(3, 5))
	fmt.Println(multiply(2, 7))
	fmt.Println(multiply(3, 4))
	fmt.Println(multiply(73807517, 14))
}

func multiply(A int, B int) int {
	// A乘以B，在二进制世界可以无限化简为：A每次乘以2,B每次整除2（如果有余，A再加一次自己）
	// 一直如此，直到B等于1.
	// 否则，A左移，B右移。如果B右移前末尾存在1，那么A+自身。（因为下一次移位会抹掉这个1）
	ret := A
	addon := 0
	for {
		// 判断B末端是不是1，如果是，A要加一次自己
		needPlus := B|1 == B
		if needPlus {
			addon += A
		}

		B = B >> 1
		if B == 0 {
			break
		} else {
			ret = ret << 1
		}
	}
	return ret + addon
}

func multiply2(A int, B int) int {
	// 乘法就相当于加x次
	// 调整一下顺序，让A大B小，然后用B当次数
	if B > A {
		temp := A
		A = B
		B = temp
	}
	if B == 0 {
		return 0
	}
	return A + multiply2(A, B-1)
}
