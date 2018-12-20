package main

import "fmt"

// 体积
var W = []int{
	11, 3, 8, 9, 7, 8, 40, 3,
}

// 价值
var D = []int{
	9, 4, 20, 9, 4, 10, 40, 1,
}

// 背包问题，递归
// 前k件物品，用m体积的背包盛放
func Backpack(k, m int) int {
	if k == 1 {
		if W[k-1] <= m {
			return D[k-1]
		}
		return 0
	} else if k <= 0 {
		return 0
	}
	if m-W[k-1] >= 0 {
		return maxInt(Backpack(k-1, m), Backpack(k-1, m-W[k-1])+D[k-1])
	}
	return Backpack(k-1, m)
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

// 背包问题，DP
// 前k件物品，用m体积的背包盛放
func Backpack2(k, m int) int {
	var dp = make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= m; j++ {
			if j-W[i-1] >= 0 {
				// 在选和不选当前商品的情况下，选择结果价值更大的那个
				dp[i][j] = maxInt(dp[i-1][j], dp[i-1][j-W[i-1]]+D[i-1])
			} else {
				// 不选当前这件物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[k][m]
}

// 背包问题，DP，使用滚动数组优化内存使用
// 前k件物品，用m体积的背包盛放
func Backpack3(k, m int) int {
	// 使用一行存储
	var dp = make([]int, m+1)
	for i := 1; i <= k; i++ {
		// 由计算特性得知，需要逆序计算才能使用单行dp缓存
		for j := m; j >= 1; j-- {
			if j-W[i-1] >= 0 {
				// 在选和不选当前商品的情况下，选择结果价值更大的那个
				dp[j] = maxInt(dp[j], dp[j-W[i-1]]+D[i-1])
			}
		}
	}
	return dp[m]
}

func main() {
	// 有N件物品，和一个容积为M的背包，每件物品的体积是W[], 每件物品的价值是D[], 每种物品只有一件
	// 求解将物品放入背包使价值最大。
	fmt.Println("Backpack=>", Backpack(4, 20))
	fmt.Println("Backpack=>", Backpack(len(W), 40))
	fmt.Println("Backpack2=>", Backpack2(4, 20))
	fmt.Println("Backpack2=>", Backpack2(len(W), 40))
	fmt.Println("Backpack3=>", Backpack3(4, 20))
	fmt.Println("Backpack3=>", Backpack3(len(W), 40))
}
