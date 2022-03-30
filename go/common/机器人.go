package main

import (
	"fmt"
	"strconv"
)

func main() {
	cmd1 := "LFFF"
	cmd2 := "10(LFF)RB"
	fmt.Println(robotGo([2]int{0, 0}, cmd1))
	fmt.Println(robotGo([2]int{0, 0}, cmd2))
}

/**
问题描述

有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？

*/

type Point [2]int

func robotGo(initPosition Point, cmd string) Point {
	// 初始化机器人状态
	// 解析指令，并根据指令前进
	fullCmd := unzipCmd(cmd)
	fmt.Println(fullCmd)

	// 表示位置
	var currentPosition Point = initPosition
	// 表示方向
	var currentDirection Point = [2]int{0, 1}

	for _, step := range fullCmd {
		// fmt.Printf("%+v, %+v\n", currentDirection, currentPosition)
		switch step {
		case 'R':
			// 一定要同时赋值，不然就要使用temp变量
			currentDirection[0], currentDirection[1] = currentDirection[1], -currentDirection[0]
		case 'L':
			currentDirection[0], currentDirection[1] = -currentDirection[1], currentDirection[0]
		// 对算子进行增减
		case 'F':
			currentPosition[0] += currentDirection[0]
			currentPosition[1] += currentDirection[1]
		case 'B':
			currentPosition[0] -= currentDirection[0]
			currentPosition[1] -= currentDirection[1]
		}
	}

	return currentPosition
}

// 将压缩命令转换为正常单次执行的命令
func unzipCmd(rawCmd string) string {
	ret := make([]rune, 0)

	times := make([]rune, 0)     // 重复次数
	repeatCmd := make([]rune, 0) // 重复指令
	inClosure := false
	for _, char := range rawCmd {
		if isNumber(char) {
			times = append(times, char)
			continue
		} else if char == '(' {
			inClosure = true
			continue
		} else if char == ')' {
			// 执行repeat
			leftTimes, _ := strconv.Atoi(string(times))
			for ; leftTimes > 0; leftTimes-- {
				ret = append(ret, repeatCmd...)
			}

			// reset vars
			inClosure = false
			repeatCmd = make([]rune, 0)
			times = make([]rune, 0)
			continue
		}
		if inClosure {
			repeatCmd = append(repeatCmd, char)
		} else {
			ret = append(ret, char)
		}
	}
	return string(ret)
}

const (
	ASCII_ZERO = 48
	ASCII_NINE = 57
)

func isNumber(char rune) bool {
	return char >= ASCII_ZERO && char <= ASCII_NINE
}
