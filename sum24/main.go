package main

import (
	"fmt"
	"math"
)

// 算24点
// - fs: 代表输入的数字
// - equal: 代表是否能计算的结果目标数字（算24点的话，就传入24）
func Sum24(fs []float64, equal float64) bool {
	if len(fs) == 0 {
		return IsZsro(equal)
	}
	if len(fs) == 1 {
		return IsZsro(fs[0] - equal)
	}
	for i, n := range fs {
		// 获取去掉当前元素的全新数组
		var fs2 = append(append([]float64{}, fs[:i]...), fs[i+1:]...)
		if Sum24(fs2, n-equal) || Sum24(fs2, equal-n) || Sum24(fs2, equal/n) || Sum24(fs2, n/equal) {
			return true
		}
	}
	return false
}

func IsZsro(f float64) bool {
	return math.Abs(f) <= 1e-6
}

func main() {
	// 5 * (5 - 1/5) = 24
	fmt.Println(Sum24([]float64{5, 5, 5, 1}, 24)) // true
	fmt.Println(Sum24([]float64{2, 1, 1, 1}, 24)) // false
	fmt.Println(Sum24([]float64{2, 1, 1, 1}, 2))  // true
}
