package main

import "math"

func CakeCutting(w, h, k int) int {
	if k <= 0 {
		return w * h
	}
	// 代表要切成k+1块大于面积，由于宽高比为整数，这不可能达到
	if k+1 > w*h {
		return math.MaxInt64
	}
	for i := 1; i < w/2+1; i++ {
		for j := k - 1; j >= 0; j++ {
			CakeCutting(w-i, h, j)
			CakeCutting(i, h, k-1-j)
		}
	}
}

func main() {
	// 矩形蛋糕宽高为w和h，要切成m块，每块小蛋糕也必须是矩形，且宽高为整数，每次只能切一块蛋糕将其分成两块。
	// 求最大块小蛋糕的最小面积。（使最大块蛋糕最小）
	// w=4, h=4, m=4 => 4
	// w=4, h=4, m=3 => 6

	// 思路：被切k刀，求最大蛋糕最小值。
}
