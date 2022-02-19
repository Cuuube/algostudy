package main

func main() {
	// testPrintBin()
	// testReverseBits()
	// testfindClosedNumbers()
	// fmt.Println(strconv.ParseInt("111", 2, 32))
	testdrawLine()

}

//-------------- 一些bit操作总结 --------------

// 清除固定bit位为0
func clearBit(target int, index int) int {
	musk := 1 << index // 00001000
	musk = ^musk       // 11110111
	return target & musk
}

// 填充固定bit位为1
func fillBit(target int, index int) int {
	musk := 1 << index // 00001000
	return target | musk
}

// 翻转bit
func reversedBit(num int, i int) int {
	toTrue := 1<<i | num
	if toTrue != num {
		return toTrue
	}
	return ^(1 << i) & num
}

// 查找第n位是true还是false
func getBit(num int, i int) bool {
	musk := 1 << i
	return num|musk == num
}

// 查找bit中有几位是1
func getBitTrueCnt(num int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if getBit(num, i) {
			cnt++
		}
	}
	return cnt
}
