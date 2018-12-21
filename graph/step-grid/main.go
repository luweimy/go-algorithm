package main

import "fmt"

// i, j 代表当前走的位置，用于记录是否访问过
// n 代表剩余步数
func Ways(visited [][]bool, i, j, n int) int {
	if n <= 0 {
		return 1
	}
	if visited[i][j] {
		return 0
	}
	// 标记位置(i,j)已走过
	visited[i][j] = true
	// 当前走法就等于分别向三个方法的走法之和
	var r = Ways(visited, i+1, j, n-1)
	r += Ways(visited, i-1, j, n-1)
	r += Ways(visited, i, j+1, n-1)
	// 计算出位置(i,j)已走过的情况，下面要把位置(i,j)标记为未走过，这样才不影响后续其他走法
	visited[i][j] = false
	return r
}

func main() {
	// 踩方格问题
	// 有个方格矩阵足够大
	// a. 一步只能走一个格子
	// b. 格子只能走一次
	// c. 只能向北、东、西三个方向走
	// 请问：如果允许走N(N<=20)步，共有多少种走的方案。
	var visited = make([][]bool, 40)
	for i := range visited {
		visited[i] = make([]bool, 20)
	}
	fmt.Println("Ways", Ways(visited, 20, 0, 20))
}
