package main

import "fmt"

// 每个物品的体积
var Volumes = []int{
	20, 20, 20, 10, 10,
}

// 神奇口袋，递归
// 从前num个物品中选择，volume体积的物品
func MagicalPocket(num, volume int) int {
	if volume == 0 {
		return 1
	}
	if num <= 0 {
		return 0
	}
	// 每个物品都有选择或者不选两种情况
	// 当前物品若选择则从volumn减去当前物品的体积
	return MagicalPocket(num-1, volume) + MagicalPocket(num-1, volume-Volumes[num-1])
}

// 神奇口袋，DP
// 从前num个物品中选择，volume体积的物品
func MagicalPocket2(num, volume int) int {
	var dp = make([][]int, num+1)
	for i := range dp {
		dp[i] = make([]int, volume+1)
		// 体积为0则代表有一种情况
		dp[i][0] = 1
	}
	for i := 1; i <= num; i++ {
		for j := 0; j <= volume; j++ {
			// 每个物品都有选择或者不选两种情况
			// 当前物品若选择则从volumn减去当前物品的体积
			var t = j - Volumes[i-1]
			if t < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][t]
			}
		}
	}
	return dp[num][volume]
}

func main() {
	// 有个神奇口袋，总容积是40，用这个口袋可以变出一些物品，总体积必须是40。
	// 有n个物品，每个体积为a1~an。问有多少种不同的选择物品的方式？
	fmt.Println("MagicalPocket=>", MagicalPocket(3, 40))
	fmt.Println("MagicalPocket=>", MagicalPocket(len(Volumes), 40))
	fmt.Println("MagicalPocket2=>", MagicalPocket2(3, 40))
	fmt.Println("MagicalPocket2=>", MagicalPocket2(len(Volumes), 40))
}
