package main

import (
	"fmt"
	"sort"
)

// N个数（N<100000），其中两个数之和是m，找出这两个数
func FindNumPairs(m int, nums []int) (int, int) {
	// 方案1、直接枚举的复杂度是O(n^2)，抛弃
	// 方案2、先排序复杂度是O(NlogN)，然后从两头往里遍历一遍O(n)，总的复杂度O(NlogN)
	sort.Ints(nums)

	// 从有序数组(升序)的两头往中间走，如果当前指向的两个数之和大于m，则右侧的游标左移，反之右移
	var l, r = 0, len(nums) - 1
	for {
		var rlt = nums[l] + nums[r]
		if rlt == m {
			return nums[l], nums[r]
		} else if rlt > m {
			r--
		} else {
			l++
		}
		if l >= r {
			panic("no result") // 无结果
		}
	}
}

func main() {
	fmt.Println(FindNumPairs(10, []int{1, 5, 3, 4, 3, 6, 9, 7}))
}
