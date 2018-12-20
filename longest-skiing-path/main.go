package main

import (
	"fmt"
)

// 最长滑雪路径，递归
func LongestSkiingPathCount1(matrix [][]int) int {
	var longest = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			var t = skiing(matrix, i, j)
			if t > longest {
				longest = t
			}
		}
	}
	return longest
}

func skiing(matrix [][]int, i, j int) int {
	var cur = height(matrix, i, j)
	var l = height(matrix, i-1, j) // 左侧
	var r = height(matrix, i+1, j) // 右侧
	var t = height(matrix, i, j+1) // 上方
	var b = height(matrix, i, j-1) // 下方
	if cur < 0 {
		return 0
	}
	var count = 0
	// 左侧点高度小于当前点，递归
	if l >= 0 && cur > l {
		count = maxInt(count, skiing(matrix, i-1, j))
	}
	if r >= 0 && cur > r {
		count = maxInt(count, skiing(matrix, i+1, j))
	}
	if t >= 0 && cur > t {
		count = maxInt(count, skiing(matrix, i, j+1))
	}
	if b >= 0 && cur > b {
		count = maxInt(count, skiing(matrix, i, j-1))
	}
	// 当前点可达，记录滑雪路径+1
	return count + 1
}

// 小于0代表点不可达
func height(matrix [][]int, i, j int) int {
	if i < 0 || i >= len(matrix) {
		return -1
	}
	if j < 0 || j >= len(matrix[i]) {
		return -1
	}
	return matrix[i][j]
}

func main() {
	// 最长滑雪路径
	// 一个人可以从一个点，滑雪到上下左右相邻的四个点，但必须相邻的点高度小于当前点才能滑过去。
	// 高度分部矩阵如下matrix
	// 其中每个位置的数字代表高度
	// 求出最长的滑雪路径长度。
	// 例如此题最长路径为：25->24->23->22->21->20->19->18->17->160>15->14->13->12->11->10->9->8->7->6->5->4->3->2->1
	var matrix = [][]int{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}
	fmt.Println(LongestSkiingPathCount1(matrix))
}

func maxInt(a ...int) int {
	var m = a[0]
	for _, n := range a {
		if n > m {
			m = n
		}
	}
	return m
}
