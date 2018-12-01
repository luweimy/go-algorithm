package main

import "fmt"

const N = 8               // 棋盘大小，同时也是要摆放的皇后数
var ColumnPreLines [N]int // 每一行肯定只能摆放一个皇后，所以用于记录每行摆放的列位置

// N皇后问题（8皇后问题的扩展）
// - k: 第k行
func NQueen(k int) {
	// 已经摆好了N行的皇后
	if k == N {
		fmt.Println("ColumnPreLines", ColumnPreLines)
		return
	}

	// 遍历当前行(k)的每一列(ci)
	for ci := 0; ci < N; ci++ {
		// 判断当前列是否能摆放
		// 需要遍历上面已经摆放的每一行
		var can = true
		for li := 0; li < k; li++ {
			// 若在在同一列上，或者同一条斜线上，则存在冲突跳过
			if ColumnPreLines[li] == ci || IntAbs(k-li) == IntAbs(ColumnPreLines[li]-ci) {
				can = false
				break
			}
		}
		if can {
			ColumnPreLines[k] = ci
			NQueen(k + 1)
		}
	}
}

func IntAbs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func main() {
	NQueen(0)
}
