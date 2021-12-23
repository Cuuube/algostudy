package main

import "fmt"

func testKeyboard() {
	fmt.Println(keyboard(5, 50))
}

func keyboard(k int, n int) int {
	return doKeyboard(k, n) % 1000000007
}

func doKeyboard(k int, n int) int {
	if n <= 0 || k <= -25 {
		return 1
	}
	if k > 0 {
		return 26 * doKeyboard(k-1, n-1)
	} else {
		return (26 + k - 1) * doKeyboard(k-1, n-1)
	}
}
