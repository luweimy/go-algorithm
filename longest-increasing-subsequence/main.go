package main

import "fmt"

// 求最长上升子序列长度(递归)，未加缓存，递归用于理解解题思路
func LongestIncreasingSubSequence1(ns []int) int {
	// 找出以各个元素为终点的上升子序列中最长的那个
	var longest = 1
	for i := 1; i < len(ns); i++ {
		longest = maxInt(LongestIncreasingSubSequenceWithLastElement(ns[:i]), longest)
	}
	return longest
}

// 求以ns[len(ns)-1]为终点的最长上升子序列
func LongestIncreasingSubSequenceWithLastElement(ns []int) int {
	if len(ns) == 1 {
		return 1
	}
	if len(ns) == 0 {
		return 0
	}
	// 用于记录以最后一个元素为终点的最长子序列长度
	var longest = 0
	var last = ns[len(ns)-1]
	// 遍历当前序列（刨除最后一个元素）
	for i := 0; i < len(ns)-1; i++ {
		// 只找比当前序列最后一个元素小的，才能构成以最后一个元素为终点的子序列
		if ns[i] < last {
			// 递归计算ns[:i+1]包含第i个元素的最长子序列，+1就代表当前以最后一个元素为终点的最长子序列长度
			longest = maxInt(LongestIncreasingSubSequenceWithLastElement(ns[:i+1])+1, longest)
		}
	}
	return longest
}

func maxInt(a ...int) int {
	var m = a[0]
	for _, n := range a {
		if n > m {
			m = n
		}
	}
	return m
}

// 求最长上升子序列长度(循环)
func LongestIncreasingSubSequence2(ns []int) int {
	// 用于记录ns中对应位置为终点的最长子序列长度
	// 比如：ns[1] => dp[1] 记录的ns[1]为结尾的最长子序列长度
	var dp = make([]int, len(ns))
	// 以dp[0]结尾的最长子序列肯定是1
	dp[0] = 1
	// 逐步推导以ns[i]为终点的最长子序列，并将每一步的结果记录到dp中
	for i := 1; i < len(ns); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			// 只找比当前最后一个元素ns[i]小的，才能构成以ns[i]为终点的子序列
			if ns[j] < ns[i] {
				dp[i] = maxInt(dp[j]+1, dp[i])
			}
		}
	}
	return maxInt(dp...)
}

func main() {
	// 动态规划问题：
	// 求最长上升子序列长度，子序列元素不要求连续，后面的元素大于前面元素
	// 1,7,3,5,9,4,8=>1,3,5,8
	// 子问题：假设序列为a，以a[k]为终点的上升子序列（子序列最后一个元素是a[k]）
	// 推导出以a[1]~a[N]为终端的所有上升子序列，最长的那个即为问题答案
	// f(0) = 1
	// f(k) = max{f(i)+1; 0<=i<k } => 代表以k为终点的最长子序列，是找k之前比k小的元素的最长子序列，其中最长的那个+1就是以k为终点的最长子序列长度
	// 求出所有以每一个元素为终点的最长子序列长度后，找出其中最大的那个就是解。

	fmt.Println("LongestIncreasingSubSequence1=>", LongestIncreasingSubSequence1([]int{
		1, 7, 3, 5, 9, 4, 8, 2,
	}))
	fmt.Println("LongestIncreasingSubSequence2=>", LongestIncreasingSubSequence2([]int{
		1, 7, 3, 5, 9, 4, 8, 2,
	}))
}
