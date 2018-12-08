package main

import (
	"fmt"
	"math"
)

// 最佳加法表达式，递归解法，DP见BestAdditionExpr2
func BestAdditionExpr1(num []int, m int) int {
	if m+1 > len(num) {
		panic("'+' too many")
	}
	if m <= 0 {
		return number(num)
	}
	// 分解为剩余m-1个加号插入到数字中
	var rlts = make([]int, 0, len(num)-m)
	for i := m - 1; i < len(num)-1; i++ {
		rlts = append(rlts, BestAdditionExpr1(num[:i+1], m-1)+number(num[i+1:]))
	}
	return minInt(rlts...)
}

func number(num []int) int {
	var r = 0
	for i := 0; i < len(num); i++ {
		r += num[len(num)-i-1] * int(math.Pow10(i))
	}
	return r
}

func minInt(a ...int) int {
	if len(a) <= 0 {
		return 0
	}
	var m = a[0]
	for _, n := range a {
		if n < m {
			m = n
		}
	}
	return m
}

//  最佳加法表达式，DP
func BestAdditionExpr2(num []int, m int) int {
	// dp[i][j]  代表有i个加号(i=0代表没有加号)，代表索引j前面(包含j)的数的最小和(j从0开始)
	var dp = make([][]int, m+1)

	// dp[0]代表没有加号，则直接返回对应的num
	dp[0] = make([]int, len(num))
	for j := 0; j < len(num); j++ {
		dp[0][j] = number(num[:j+1])
	}
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, len(num))
		// 代表包含num[j]的子序列的子问题
		for j := i; j < len(num); j++ {
			t := num[:j+1]
			for k := i - 1; k < len(t)-1; k++ {
				v := dp[i-1][k] + number(num[k+1:j+1])
				if dp[i][j] <= 0 {
					dp[i][j] = v
				} else {
					dp[i][j] = minInt(dp[i][j], v)
				}
			}
		}
	}
	return dp[m][len(num)-1]
}

func main() {
	// 最佳加法表达式
	// 有一个由1~9组成的数字串，如果将m个加号插入到数字串中，值最小的那个表达式值是多少。
	// 如 2345 可以插入m=2个加号变为：2+34+5，23+4+5，2+3+45等
	fmt.Println("BestAdditionExpr1", BestAdditionExpr1([]int{2, 3, 3, 1}, 1)) // => 54
	fmt.Println("BestAdditionExpr1", BestAdditionExpr1([]int{2, 3, 3, 1}, 2)) // => 27
	fmt.Println("BestAdditionExpr1", BestAdditionExpr1([]int{2, 3, 3, 1}, 3)) // => 9

	fmt.Println("BestAdditionExpr2", BestAdditionExpr2([]int{2, 3, 3, 1}, 1)) // => 54
	fmt.Println("BestAdditionExpr2", BestAdditionExpr2([]int{2, 3, 3, 1}, 2)) // => 27
	fmt.Println("BestAdditionExpr2", BestAdditionExpr2([]int{2, 3, 3, 1}, 3)) // => 9
}
