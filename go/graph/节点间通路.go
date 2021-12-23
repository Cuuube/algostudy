package main

import "fmt"

/**
面试题 04.01. 节点间通路
节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。

示例1:

 输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
 输出：true
示例2:

 输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
 输出 true
提示：

节点数量n在[0, 1e5]范围内。
节点编号大于等于 0 小于 n。
图中可能存在自环和平行边。
https://leetcode-cn.com/problems/route-between-nodes-lcci/
*/

func testFindWhetherExistsPath() {
	graph := [][]int{{0, 1}, {0, 3}, {0, 10}, {0, 18}, {1, 2}, {1, 7}, {1, 11}, {1, 12}, {2, 4}, {2, 5}, {2, 13}, {2, 16}, {3, 6}, {3, 8}, {4, 9}, {5, 17}, {7, 20}, {7, 22}, {8, 10}, {10, 19}, {11, 15}, {13, 14}, {14, 21}, {15, 23}, {19, 24}, {20, 22}}
	fmt.Println(findWhetherExistsPath(25, graph, 0, 12))
}

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	routes := make(map[int]map[int]bool)
	// 转换为邻接表
	for _, path := range graph {
		start := path[0]
		end := path[1]
		route, existed := routes[start]
		if !existed {
			route = make(map[int]bool)
		}
		route[end] = true
		routes[start] = route
	}

	cachedNodes := map[int]bool{start: true}

	// 遍历邻接表，注意环
	ends := routes[start]
	for len(ends) > 0 {
		newEnds := make(map[int]bool)
		for end := range ends {
			cachedNodes[end] = true

			// 检测终点
			if end == target {
				return true
			}
			// 不是终点，就把节点放到遍历目标上去
			// 防止环
			for e := range routes[end] {
				if cachedNodes[e] || newEnds[e] {
					continue
				}
				newEnds[e] = true
			}
		}
		ends = newEnds
	}
	return false
}
