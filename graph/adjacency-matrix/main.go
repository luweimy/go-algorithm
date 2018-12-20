package main

import "fmt"

const N = 10

// 判断是否start和target是否联通，并且返回路径
func DFS(matrix [][]int, start, target int, path []int) ([]int, bool) {
	path = append(path, start)
	if start == target {
		return path, true
	}
	// 遍历起始点start的所有邻接点，为1才代表有路径
	for node, b := range matrix[start] {
		if b == 1 {
			// 避免死循环
			if contains(path, node) {
				continue
			}
			if pathNew, ok := DFS(matrix, node, target, path); ok {
				return pathNew, true
			}
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
	var matrix = make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	// 无向图，节点0~9
	// 9
	// |
	// 3-1-2
	//   | |
	//   4-8
	matrix[1][2] = 1 // 为1则代表从1->2是联通的
	matrix[1][3] = 1
	matrix[1][4] = 1
	matrix[4][8] = 1
	matrix[3][9] = 1
	matrix[2][8] = 1

	// 上面只给正向设置了联通，下面循环给反向也设置上联通
	for i, l := range matrix {
		for j, e := range l {
			matrix[j][i] = e
		}
	}

	dump(matrix)
	fmt.Println(DFS(matrix, 1, 8, nil))
	fmt.Println(DFS(matrix, 1, 9, nil))
	fmt.Println(DFS(matrix, 1, 7, nil))
}

func dump(matrix [][]int) {
	for _, l := range matrix {
		fmt.Println(l)
	}
}
