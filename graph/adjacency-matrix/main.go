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
	for i, b := range matrix[start] {
		if b == 1 {
			if pathNew, ok := DFS(matrix, i, target, path); ok {
				return pathNew, true
			}
		}
	}
	return path, false
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
	matrix[1][2] = 1
	matrix[1][3] = 1
	matrix[1][4] = 1
	matrix[4][8] = 1
	matrix[3][9] = 1
	matrix[2][8] = 1

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
