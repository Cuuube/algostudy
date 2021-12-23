package main

import (
	"fmt"
)

func testSolveNQueens() {
	// chess := []string{"    ", "    ", "    ", "    "}
	// putQueen(0, 1, chess)
	// js, _ := json.Marshal(chess)
	// fmt.Printf("%s\n", string(js))

	// fmt.Printf("%+v\n", solveNQueens(4))
	fmt.Printf("%+v\n", solveNQueens(8))
}

/**
面试题 08.12. 八皇后
设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展

示例:

 输入：4
 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
https://leetcode-cn.com/problems/eight-queens-lcci/
*/

/** 思路：
	每行、每列只能放一个皇后。
	先初始化棋盘为空值（ ）。每放一个皇后(Q)，将米字形格子翻转为不可放置（.）
	因为要找到所有情况，所以必须遍历到底，不能提前return
**/
func solveNQueens(n int) [][]string {
	// init chess
	chess := make([]string, n)
	for idx := range chess {
		row := ""
		for i := 0; i < n; i++ {
			row += " "
		}
		chess[idx] = row
	}

	// solve
	return dfs(0, chess)
}

func dfs(rowIndex int, chess []string) [][]string {
	// 最后一层，返回布置好的棋局
	if rowIndex == len(chess) {
		return [][]string{chess}
	}

	queenInserted := false // 判读啊是否这一行放了Q。如果没放Q，之后路线不再计算
	ret := make([][]string, 0)
	row := chess[rowIndex]

	// 提醒：dfs一颗节点为n的树，可以转换为一次遍历，一次递归。
	// 这里视为一颗从原点，第一行任意位置的树。
	for cIdx, cell := range row {
		// 剪枝：必须是空白值，才继续操作。
		if string(cell) != " " {
			continue
		}
		// 注意：chess需要拷贝
		newChess := make([]string, len(chess))
		copy(newChess, chess) // 调用copy，数组长度必须一致

		// 递归：当前位置放Q
		putQueen(rowIndex, cIdx, newChess)
		ret = append(ret, dfs(rowIndex+1, newChess)...)
		queenInserted = true
	}
	if !queenInserted {
		return [][]string{}
	}

	return ret
}

// 往棋盘里放Q，以及写入Q的势力范围
func putQueen(rowIndex, cellIndex int, chess []string) {
	// q的米字形都更换为0；只从rowIndex开始更换即可
	// 优化：可以避免执行O(n2)的复杂度，遍历横线，竖线，和两个角线就可以
	for rIdx := rowIndex; rIdx < len(chess); rIdx++ {
		row := chess[rIdx]
		newRow := ""
		for cIdx, cell := range row {
			// 是目标节点，就存入Q
			if rIdx == rowIndex && cIdx == cellIndex {
				newRow += "Q"
				continue
			}
			// 如果已经有内容，就放入原内容
			char := string(cell)
			if char != " " {
				newRow += char
				continue
			}
			// 计算是否在那个Q的十字方向，或者对角线
			distance := rIdx - rowIndex
			if cIdx == cellIndex || rIdx == rowIndex || cIdx-distance == cellIndex || cIdx+distance == cellIndex {
				// 是的话存入不可放置节点
				newRow += "."
			} else {
				// 不是的话，继续保持空白
				newRow += " "
			}
		}
		chess[rIdx] = newRow
	}
}

// func copy(arr []string) []string {
// 	ret := make([]string, len(arr))
// 	for i, v := range arr {
// 		ret[i] = v
// 	}
// 	return ret
// }

// func solveNQueens(n int) [][]string {
// 	return setQueen(0, n, nil)
// }

// func setQueen(rowIndex int, n int, chess []string) [][]string {
// 	if chess == nil {
// 		chess = []string{"....", "....", "....", "...."}
// 	}
// 	ret := make([][]string, 0)
// 	for i := 0; i < n; i++ {
// 		// newChess := []string{chess...}
// 		// row := newChess[rowIndex]
// 		// row[i] = "Q"
// 		// newChess[rowIndex] = row

// 		// 检查这个地方是不是可以放Q
// 		if checkValid(chess, rowIndex, i) {
// 			row := chess[rowIndex]
// 			row = row[0:i] + "Q" + row[i+1:]
// 			chess[rowIndex] = row
// 		} else {
// 			continue
// 		}
// 		if rowIndex < n-1 {
// 			// 下一行放Q
// 			result := setQueen(rowIndex+1, n, chess)
// 			if len(result) != 0 {
// 				ret = append(ret, result...)
// 			}
// 		}
// 	}
// 	return ret
// }

// func checkValid(chess []string, rowIndex, cellIndex int) bool {
// 	// 同行有q，同列有q，或者对角线有Q，都不行
// 	for rIdx, row := range chess {
// 		for cIdx, cell := range row {
// 			if cell != "Q" {
// 				continue
// 			}
// 			if cIdx == cellIndex || rIdx == rowIndex ||
// 		}
// 	}
// 	return true
// }
