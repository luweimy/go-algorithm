package main

import "fmt"

// 走台阶问题
// 每次走一步或者两步，上到N级有多少种走法
// - n: 台阶数
func stairs(n int) int {
	// 上到N阶，存在两种可能，从n-1阶上或者从n-2阶上，所以上到N阶的情况等于两种情况之和
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return stairs(n-1) + stairs(n-2)
}

// 走台阶问题(DP)
func stairs2(n int) int {
	// 状态为上到第n阶台阶的走法
	// f(n)=f(n-1)+f(n-2)
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println("stairs=>", stairs(10))
	fmt.Println("stairs2=>", stairs2(10))
}
