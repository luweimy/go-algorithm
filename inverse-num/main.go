package main

import (
	"fmt"
)

// 求逆数O(NlogN)
func InverseNum(ns []int) int {
	// 思路：采用分治的思想
	// 将数列分为两组，分别求每组的逆数
	// 然后两组分别求完逆数之后，再求两组之间的逆数（这一步为了达到O(n)，需要两组分别有序，所以此问题是在归并排序基础上做的）
	if len(ns) <= 1 {
		return 0
	}
	m := len(ns) / 2
	n1, n2 := ns[:m], ns[m:]
	// 对两组分别查找逆数
	r := InverseNum(n1) + InverseNum(n2)
	// 对两组组合查找逆数
	r += count(n1, n2)
	// 归并（升序）
	nsr := merge(n1, n2)
	copy(ns, nsr) // 将结果拷贝会ns数列
	return r
}

// 找出所有左侧与右侧组合的逆数O(n)
func count(n1, n2 []int) int {
	// 注意：两侧数组均为升序
	var rlt = 0
	// 重点：两侧均从最大值开始遍历
	l, r := len(n1)-1, len(n2)-1
	for l >= 0 && r >= 0 {
		// 构成逆数
		if n1[l] > n2[r] {
			rlt += r + 1 // n2中比n2[r]小的也都能构成逆数
			l--
		} else {
			r--
		}

	}
	return rlt
}

func merge(n1, n2 []int) []int {
	var tmp = make([]int, len(n1)+len(n2))
	var i, i1, i2 = 0, 0, 0
	for ; i1 < len(n1) && i2 < len(n2); i++ {
		if n1[i1] < n2[i2] {
			tmp[i] = n1[i1]
			i1++
		} else {
			tmp[i] = n2[i2]
			i2++
		}
	}
	if i1 < len(n1) {
		copy(tmp[i:], n1[i1:])
	} else if i2 < len(n2) {
		copy(tmp[i:], n2[i2:])
	}
	return tmp
}

func main() {
	// 2,6,3,4,5,1 => (2,1) (6,3) (6,4) (6,5) (6,1) (3,1) (4,1) (5,1)
	// 结果是8组逆数(逆数即为：前面比后面大的一对儿数)
	var ns = []int{
		2, 6, 3, 4, 5, 1,
	}
	fmt.Println("InverseNum=>", InverseNum(ns))
	fmt.Println("ns=>", ns)
}
