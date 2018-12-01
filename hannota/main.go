package main

import "fmt"

// 汉诺塔
// - n: 代表有n层
// - org, tmp, dst: 分别代表原始位置，中转位置，目标位置
func Hannota(n int, org, tmp, dst string) {
	// 移动N个可以分解为：
	// 1.org移动N-1个到tmp，以dst为中转位置
	// 2.把org剩下的最底下的一个移到dst
	// 3.把tmp的N-1个以org为中转位置，按同样的方法移到dst
	if n == 1 {
		fmt.Println("#", org, "->", dst)
		return
	}
	Hannota(n-1, org, dst, tmp)
	Hannota(1, org, tmp, dst)
	Hannota(n-1, tmp, org, dst)
}

func main() {
	Hannota(3, "A", "B", "C")
}
