package main

import (
	"fmt"
	"math/rand"
)

// 二分查找
func BinarySearch(slice []int, t int) int {
	l, r := 0, len(slice)-1
	for l <= r {
		m := (l + r) / 2
		if t > slice[m] {
			l = m + 1
		} else if t < slice[m] {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

// 找最长回文子串 O(N^2)
// 从中心向两边找
func LongestPalindrome(s string) string {
	var lps string
	for i := 0; i < len(s); i++ {
		for n := 0; i-n >= 0 && i+n < len(s); n++ {
			// 判断奇数位的回文串
			ps := s[i-n : i+n+1]
			if isPalindrome(ps) && len(ps) > len(lps) {
				lps = ps
			}
			// 判断偶数位的回文串
			if i+1+n < len(s) {
				ps = s[i-n : i+1+n+1]
				if isPalindrome(ps) && len(ps) > len(lps) {
					lps = ps
				}
			}
		}
	}
	return lps
}

// 判断是否是回文串
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 找第K大的数
// 利用快排算法的思想找
func FindMaxK(slice []int, k int) int {
	if k > len(slice) || k <= 0 {
		panic(-1)
	}
	p, l, r := 0, 0, len(slice)-1
	for l < r {
		for slice[r] >= slice[p] && l < r {
			r--
		}
		for slice[l] <= slice[p] && l < r {
			l++
		}
		if l >= r {
			slice[p], slice[l] = slice[l], slice[p]
			break
		}
		slice[l], slice[r] = slice[r], slice[l]
	}
	// 左右哨兵相遇位置就是基准数排好序的位置，计算相遇位置是第n大的数
	n := len(slice) - l
	if k > n {
		return FindMaxK(slice[:l], k-n)
	} else if k < n {
		return FindMaxK(slice[l+1:], k)
	} else {
		return slice[len(slice)-k]
	}
}

func main() {
	fmt.Println("BinarySearch=>", BinarySearch(sortedInts(10), 100))
	fmt.Println("LongestPalindrome(abbabba)=>", LongestPalindrome("xxabbabbacc"))
	fmt.Println("LongestPalindrome(aabbaa)=>", LongestPalindrome("xaabbaa"))
	fmt.Println("LongestPalindrome(aba)=>", LongestPalindrome("abax"))
	fmt.Println("LongestPalindrome(aba)=>", LongestPalindrome("abax"))
	fmt.Println("LongestPalindrome(aab)=>", LongestPalindrome("aab"))

	ints := randomInts(10)
	fmt.Println("FindMaxK(k=6)=>", ints, FindMaxK(ints, 6))
	fmt.Println("FindMaxK(k=1)=>", ints, FindMaxK(ints, 1))
	fmt.Println("FindMaxK(k=10)=>", ints, FindMaxK(ints, 10))
}

func sortedInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = i
	}
	return res
}

func randomInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(n)
	}
	return res
}
