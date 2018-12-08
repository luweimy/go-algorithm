package main

import "fmt"

func maxInt(a ...int) int {
	var m = a[0]
	for _, n := range a {
		if n > m {
			m = n
		}
	}
	return m
}

// 最长公共子序列长度（递归）
func LongestCommonSubSequence1(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	// 两个串的最后一个字符相同
	if s1[len(s1)-1] == s2[len(s2)-1] {
		// 两个串都刨去最后一个字符，然后求其最长公共子序列
		return LongestCommonSubSequence1(s1[:len(s1)-1], s2[:len(s2)-1]) + 1
	} else {
		// 若最后一个字符不等，则求其他所有情况的最长公共子序列
		// 计算两个串分别去掉一个字符的各种情况，从而减小问题，求子问题的解
		return maxInt(
			LongestCommonSubSequence1(s1, s2[:len(s2)-1]),
			LongestCommonSubSequence1(s1[:len(s1)-1], s2),
			//LongestCommonSubSequence1(s1[:len(s1)-1], s2[:len(s2)-1]), // 此情况肯定不会比上面两种情况长，所以不用计算
		)
	}
}

// 最长公共子序列长度（DP）
func LongestCommonSubSequence2(s1, s2 string) int {
	var dp = make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i][j] = take(dp, i-1, j-1) + 1
			} else {
				dp[i][j] = maxInt(
					take(dp, i-1, j), take(dp, i, j-1), // dp[i-1][j-1], 同理，dp[i-1][j-1]肯定不会比另外两个大
				)
			}
		}
	}
	return dp[len(s1)-1][len(s2)-1]
}

func take(dp [][]int, i, j int) int {
	// 因为s1或s2为0时，其最长公共子序列肯定为0
	if i < 0 || j < 0 {
		return 0
	}
	return dp[i][j]
}

func main() {
	// 动态规划问题：
	// 求最长公共子序列长度，子序列元素不要求连续，但要求顺序
	// abcfbc, abfcab => 4(abcb)
	// programming, contest => 2(on)
	// abcd, mnp => 0
	// 子问题：见递归解法
	fmt.Println("LongestCommonSubSequence1=>", LongestCommonSubSequence1("abcfbc", "abfcab"))
	fmt.Println("LongestCommonSubSequence1=>", LongestCommonSubSequence1("programming", "contest"))
	fmt.Println("LongestCommonSubSequence1=>", LongestCommonSubSequence1("abcd", "mnp"))
	fmt.Println("LongestCommonSubSequence2=>", LongestCommonSubSequence2("abcfbc", "abfcab"))
	fmt.Println("LongestCommonSubSequence2=>", LongestCommonSubSequence2("programming", "contest"))
	fmt.Println("LongestCommonSubSequence2=>", LongestCommonSubSequence2("abcd", "mnp"))
}
