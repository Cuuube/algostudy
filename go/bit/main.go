package main

func main() {
	// testPrintBin()
	// testReverseBits()
	// testfindClosedNumbers()
	// fmt.Println(strconv.ParseInt("111", 2, 32))
	testdrawLine()

}

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

func getBitTrueCnt(num int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		turnTrue := 1<<i | num
		if turnTrue == num {
			cnt++
		}
	}
	return cnt
}

func reversedBit(num int, i int) int {
	toTrue := 1<<i | num
	if toTrue != num {
		return toTrue
	}
	return ^(1 << i) & num
}
