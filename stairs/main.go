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

func main() {
	fmt.Println("stairs=>", stairs(10))
}
