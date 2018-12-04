package main

import "fmt"

// 计算数字三角路径最大和 O(2^n)（简单递归）
//     7
//    3 8
//   8 1 0
//  2 7 4 4
// 4 5 2 6 5
// 找一条数字三角形顶部到底部的路径，使得路径上数字之和最大。每次只能左下和右下方向走
// - ns: 用于存储三角形数据
// - i,j: 是当前处于行，列
func TriangleMaxSum1(ns [][]int, i, j int) int {
	if len(ns) == i {
		return 0
	}
	return ns[i][j] + maxInt(TriangleMaxSum1(ns, i+1, j), TriangleMaxSum1(ns, i+1, j+1))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 带记忆的递归计算（动态规划）
//     7
//    3 8
//   8 1 0
//  2 7 4 4
// 4 5 2 6 5
// 简单的递归运算，存在大量的重复计算，所以性能低下
// 采用缓存避免重复计算，可以极大的提升性能
// - dp: 用于存储每个位置的最大路径和，默认为-1
func TriangleMaxSum2(ns, dp [][]int, i, j int) int {
	if len(ns) == i {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	r := ns[i][j] + maxInt(TriangleMaxSum2(ns, dp, i+1, j), TriangleMaxSum2(ns, dp, i+1, j+1))
	dp[i][j] = r
	return r
}

// 采用循环递推（动态规划）
//     7
//    3 8
//   8 1 0
//  2 7 4 4
// 4 5 2 6 5
// 摒弃递归，采用从最底层往上递推的方式计算
func TriangleMaxSum3(ns [][]int) int {
	// 用ns的最后一行作为缓存，最后一层上的数字就是最后一层各个位置的最大路径和
	// 仅缓存一行即可，然后用此行向上递推，然后再缓存上面的那一行。
	var dp = ns[len(ns)-1]
	for i := len(ns) - 2; i >= 0; i-- {
		// 计算当前行每一个位置的最大路径和
		for j := 0; j < len(ns[i]); j++ {
			// 当前行的元素ns[i][j]与下面那一行的左右两个位置中最大的那个相加就是当前最大路径和
			// dp[j], dp[j+1] 就是缓存的下面那行左右的两个位置的最大路径和
			// 下面那行的dp[j]会被当前行位置j的最大路径和覆盖，但是不影响
			dp[j] = ns[i][j] + maxInt(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

func main() {
	//     7
	//    3 8
	//   8 1 0
	//  2 7 4 4
	// 4 5 2 6 5
	ns := [][]int{
		{7},
		{3, 8},
		{8, 1, 0},
		{2, 7, 4, 4},
		{4, 5, 2, 6, 5},
	}
	fmt.Println("TriangleMaxSum1=>", TriangleMaxSum1(ns, 0, 0))
	dp := makeEmptyCache(ns)
	fmt.Println("TriangleMaxSum2=>", TriangleMaxSum2(ns, dp, 0, 0))
	fmt.Println("TriangleMaxSum2=> dp", dp)
	fmt.Println("TriangleMaxSum3=>", TriangleMaxSum3(ns))
	ns2 := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	fmt.Println("TriangleMaxSum1=>", TriangleMaxSum1(ns2, 0, 0))
}

func makeEmptyCache(ns [][]int) [][]int {
	var dp = make([][]int, len(ns))
	for i, n := range ns {
		dp[i] = makeEmptySlice(n)
	}
	return dp
}

func makeEmptySlice(s []int) []int {
	var c = make([]int, len(s))
	for i := range c {
		c[i] = -1
	}
	return c
}
