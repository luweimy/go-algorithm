package main

import (
	"fmt"
	"math"
)

func IsZero(f float64) bool {
	return math.Abs(f) <= 1e-6
}

func Solving(l, r float64, fx func(f float64) float64) float64 {
	// 前提是方程单调递增，用二分法，快速缩小范围，时间复杂度O(logN)
	for {
		mid := (l + r) / 2
		rlt := fx(mid)
		if IsZero(rlt) {
			return mid
		} else if rlt > 0 {
			r = mid
		} else {
			l = mid
		}
	}
	return -1
}

func Fx(x float64) float64 {
	return x*x*x - 5*x*x + 10*x - 80
}

func main() {
	// x^3-5x^2+10x-80=0
	// 注意并非所有方程能被此方式求解，需要单调递增
	// 先缩小范围f(0)<0 && f(100) > 0, 所以在此区间求解
	fmt.Println("x=>", Solving(0, 100, Fx))
	fmt.Println("Fx(5.7)", Fx(5.7))
}
