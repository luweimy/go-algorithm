package main

import "fmt"

const N = 10

// 判断是否start和target是否联通，并且返回路径
func DFS(adjacencyList [][]int, start, target int, path []int) ([]int, bool) {
	path = append(path, start)
	if start == target {
		return path, true
	}
	for _, node := range adjacencyList[start] {
		// 避免死循环
		if contains(path, node) {
			continue
		}
		// 深度优先搜索
		if pathNew, ok := DFS(adjacencyList, node, target, path); ok {
			return pathNew, true
		}
	}
	return path, false
}

func contains(ns []int, t int) bool {
	for _, n := range ns {
		if n == t {
			return true
		}
	}
	return false
}

func main() {
	// 图的存储：邻接矩阵
	var adjacencyList = make([][]int, N)

	// 无向图，节点0~9
	// 9
	// |
	// 3-1-2
	//   | |
	//   4-8
	adjacencyList[1] = []int{2, 3}
	adjacencyList[2] = []int{1, 8}
	adjacencyList[3] = []int{1, 9}
	adjacencyList[4] = []int{1, 8}
	adjacencyList[8] = []int{2, 4}
	adjacencyList[9] = []int{3}

	dump(adjacencyList)
	fmt.Println(DFS(adjacencyList, 1, 8, nil))
	fmt.Println(DFS(adjacencyList, 1, 9, nil))
	fmt.Println(DFS(adjacencyList, 1, 7, nil))
}

func dump(adjacencyList [][]int) {
	for i, l := range adjacencyList {
		fmt.Println(i, "=>", l)
	}
}
