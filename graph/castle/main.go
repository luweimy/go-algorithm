package main

import "fmt"

const (
	WallLeft   = 0x01 // 左
	WallTop    = 0x02 // 上
	WallRight  = 0x04 // 右
	WallBottom = 0x08 // 下
)

// 深度优先遍历，所有房间，并且给不同的房间染色，同一个颜色代表同一个房间
// 最后统计有几种颜色，最多的单个颜色是哪个是多少即可
func CastleRoom(walls [][]int) [][]int {
	var colors = make([][]int, len(walls))
	for i := range colors {
		colors[i] = make([]int, len(walls[i]))
	}
	var color = 1
	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors[i]); j++ {
			// 因为不同的房间不连通，所以要遍历所有节点找到每个没有被遍历的点去遍历
			if colors[i][j] == 0 {
				// 遍历单个联通子图，并且给染色为color，当返回时，此联通子图每个节点都已被遍历过
				var area = DFS(walls, colors, i, j, color, 0)
				fmt.Println("color=>", color, "area=>", area)
				// 换种颜色，给其他房间染色
				color++
			}
		}
	}
	return colors
}

// 遍历单个联通的子图，返回房间的面积
func DFS(walls, colors [][]int, i, j, color, area int) int {
	if i < 0 || j < 0 || i >= len(walls) || j >= len(walls[i]) {
		return area
	}
	// 已经走过的地方
	if colors[i][j] > 0 {
		return area
	}
	// 当前位置之前没有走过，记录面积
	area++
	// 给当前位置染色
	colors[i][j] = color
	// 判断每个方向是不是没有墙
	var w = walls[i][j]
	if w&WallLeft == 0 {
		area = DFS(walls, colors, i, j-1, color, area)
	}
	if w&WallTop == 0 {
		area = DFS(walls, colors, i-1, j, color, area)
	}
	if w&WallRight == 0 {
		area = DFS(walls, colors, i, j+1, color, area)
	}
	if w&WallBottom == 0 {
		area = DFS(walls, colors, i+1, j, color, area)
	}
	return area
}

func main() {
	// 城堡问题
	// 每个位置记录其四个方向是否有墙
	// 每个联通的子图算一个房间，求房间数量和最大房间的面积

	//     1   2   3   4   5   6   7
	//   #############################
	// 1 #       #       #           #
	//   #####   #####   #   #####   #
	// 2 #   #       #   #   #   #   #
	//   #   #####   #####   #####   #
	// 3 #           #   #   #   #   #
	//   #   #########   #####   #   #
	// 4 #   #                   #   #
	//   #############################

	// 每一位数字代表其四个方向上是否有墙壁
	var walls = [][]int{
		{11, 6, 11, 6, 3, 10, 6},
		{7, 9, 6, 13, 5, 15, 5},
		{1, 10, 12, 7, 13, 7, 5},
		{13, 11, 10, 8, 10, 12, 13},
	}
	dump(CastleRoom(walls))
}

func dump(matrix [][]int) {
	for _, l := range matrix {
		fmt.Println(l)
	}
}
