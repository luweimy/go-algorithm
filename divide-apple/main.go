package main

import "fmt"

// 分苹果问题
// 把m个同样的苹果放在n个同样的盘子里，允许有的盘子空着不放，问有多少种不同的分法？
func DivApple(m, n int) int {
	if m == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	// 如果苹果数量少于盘子
	if m < n {
		// 则多余盘子不影响分法，扔掉(注意是问有几种分法，每个盘子都一样)
		return DivApple(m, m)
	}

	// 苹果数量多于盘子，分两种情况处理，一种是摆满所有盘子，一种是不摆满盘子
	// 不摆满盘时，仅去掉一个盘子，然后变为子问题m个苹果分到n-1个盘子（因为子问题还会有不摆满，而去掉一个盘子的情况，所以这里就不用再去掉第二个盘子，否则就重复了）
	return DivApple(m, n-1) + DivApple(m-n, n)
}

func main() {
	fmt.Println("DivApple=>", DivApple(7, 3)) // 8
}
