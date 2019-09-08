package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

// InsertionSort 直接插入排序 O(N^2) 稳定
// 将数组分为两部分，待排序部分，已排序部分。
// 从已排序部分依次取出，然后插入到已排序部分适当位置，使已排序部分仍然有序。
func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		// 以此为分界线，数组前面为已排序部分，后面为待排序部分
		// 后面slice[i]会被覆盖，所以先取出
		e := slice[i]
		// 在待排序部分找到合适的位置插入
		j := i - 1
		for ; j >= 0 && slice[j] > e; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = e
	}
}

// BubbleSort 冒泡排序 O(N^2) 稳定
// 依次比较相邻的两个数，将小数放在前面，大数放在后面。
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// SelectionSort 选择排序 O(N^2) 不稳定
// 将数组分为两部分，待排序部分，已排序部分。
// 每次从待排序部分选一个最小的，依次放入已排序部分。
func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		// i前面是已排序部分，m用来记录要查找最小值的index
		m := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[m] {
				m = j
			}
		}
		slice[i], slice[m] = slice[m], slice[i]
	}
}

func SelectionSortV2(slice []int) {
	for i := 0; i < len(slice); i++ {
		// slice[i]位置保存临时最大值
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// BucketSortV2 桶排序 O(N) 稳定
// 将待排序数据分别放入对应的桶中，最后将桶中数据分别进行插入排序，然后依次放入待排序数组。
func BucketSortV2(slice []int) {
	const N = 10
	buckets := make(map[int][]int, len(slice)/N)
	for _, n := range slice {
		buckets[n/N] = append(buckets[n/N], n)
	}
	var i = 0
	for _, k := range sortMapKeys(buckets) {
		bucket := buckets[k]
		InsertionSort(bucket)
		copy(slice[i:i+len(bucket)], bucket)
		i += len(bucket)
	}
}

func BucketSortV3(slice []int) {
	const N = 10
	var buckets = make(map[int][]int, len(slice)/N)
	for _, n := range slice {
		buckets[n/N] = append(buckets[n/N], n)
	}
	var i = 0
	// 注意要顺序遍历buckets
	for _, k := range sortMapKeys(buckets) {
		bu := buckets[k]
		BucketSortV2(bu)
		copy(slice[i:i+len(bu)], bu)
		i += len(bu)
	}
}

// MergeSort 归并排序 O(NlogN) 稳定
// 把数组分成两个子数组,对每一个子数组递归进行归并排序，把子数组与子数组合并,合并后仍然有序,直到全部合并完,形成有序的数组。
func MergeSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	s1 := slice[:len(slice)/2]
	s2 := slice[len(slice)/2:]
	MergeSort(s1)
	MergeSort(s2)
	merge(slice, s1, s2)
}

func merge(to, s1, s2 []int) {
	if len(to) != len(s1)+len(s2) {
		panic(-1)
	}
	i, i1, i2 := 0, 0, 0
	for ; i1 < len(s1) && i2 < len(s2); i++ {
		if s1[i1] <= s2[i2] {
			to[i] = s1[i1]
			i1++
		} else {
			to[i] = s2[i2]
			i2++
		}
	}
	if i1 < len(s1) {
		copy(to[i:], s1[i1:])
	} else if i2 < len(s2) {
		copy(to[i:], s2[i2:])
	}
}

// QuickSort 快排 O(NlogN) 不稳定
// 通过一趟排序将要排序的数据分割成两部分，其中一部分的所有数据都比另外一部分的所有数据小，然后再按此方法对这两部分数据分别递归进行，以此达到整个数据有序。
func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	// p(pivot)是基准值位置，l(left)是左侧哨兵位置，r(right)是右侧哨兵位置
	// 此处直接用最左边的值作为基准值，更合理的方法是最左，最右，中间三个数，取中间的值作为基准值。(见QuickSortV2)
	p, l, r := 0, 0, len(slice)-1
	for l < r {
		// 从后面向前找小于基准数的位置
		// 必须是最右侧的哨兵先前移，这关系到哨兵交汇位置的值是否小于等于基准值
		for ; slice[r] >= slice[p] && l < r; r-- {
		}
		// 从前面向后找大于基准数的位置
		for ; slice[l] <= slice[p] && l < r; l++ {
		}
		if l >= r {
			// 哨兵交汇位置与基准值交换
			slice[p], slice[l] = slice[l], slice[p]
			break
		}
		slice[r], slice[l] = slice[l], slice[r]
	}
	QuickSort(slice[:l])
	QuickSort(slice[l+1:])
}

func QuickSortV2(slice []int) {
	if len(slice) <= 1 {
		return
	}
	doPivot(slice) // 处理基准点的值为，左中右的中间值，并将此值和slice[0]交换位置
	p, l, r := 0, 0, len(slice)-1
	for l < r {
		for ; slice[r] >= slice[p] && l < r; r-- {
		}
		for ; slice[l] <= slice[p] && l < r; l++ {
		}
		if l >= r {
			slice[l], slice[p] = slice[p], slice[l]
			break
		}
		slice[l], slice[r] = slice[r], slice[l]
	}
	QuickSortV2(slice[:l])
	QuickSortV2(slice[l+1:])
}

func doPivot(slice []int) {
	var max = 0
	if slice[max] < slice[len(slice)/2] {
		max = len(slice) / 2
	}
	if slice[max] < slice[len(slice)-1] {
		max = len(slice) - 1
	}
	slice[0], slice[max] = slice[max], slice[0]
}

// ShellSort 希尔排序
// 希尔排序是改进的插入排序，有增量gap，每次间隔gap个元素对比大小并交换位置，直到gap=1，则为插入排序
// 因为gap=1之前，使数据基本有序，此时用插入排序性能会更好
// 讲解：https://www.cnblogs.com/chengxiao/p/6104371.html
func ShellSort(slice []int) {
	for gap := len(slice) / 2; gap >= 1; gap /= 2 {
		for i := gap; i < len(slice); i++ {
			e := slice[i]
			j := i - gap
			for ; j >= 0 && slice[j] > e; j -= gap {
				slice[j+gap] = slice[j]
			}
			slice[j+gap] = e
		}
	}
}

func main() {
	unsorted := randomInts(20000)
	fmt.Println("unsorted", unsorted)

	withSort("InsertionSort=>", InsertionSort, unsorted, false)
	withSort("BubbleSort=>", BubbleSort, unsorted, false)
	withSort("SelectionSort=>", SelectionSort, unsorted, false)
	withSort("SelectionSortV2=>", SelectionSortV2, unsorted, false)
	withSort("BucketSortV2=>", BucketSortV2, unsorted, false)
	withSort("BucketSortV3=>", BucketSortV3, unsorted, false)
	withSort("MergeSort=>", MergeSort, unsorted, false)
	withSort("QuickSort=>", QuickSort, unsorted, false)
	withSort("QuickSortV2=>", QuickSortV2, unsorted, false)
	withSort("ShellSort=>", ShellSort, unsorted, false)
	withSort("Sort.Ints=>", sort.Ints, unsorted, false)
}

func randomInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(n)
	}
	return res
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func copySlice(slice []int) []int {
	res := make([]int, len(slice))
	for i, e := range slice {
		res[i] = e
	}
	return res
}

func withSort(prompt string, sortfn func([]int), slice []int, debug bool) time.Duration {
	copy1 := copySlice(slice)
	t := time.Now()
	sortfn(copy1)
	d := time.Now().Sub(t)
	copy2 := copySlice(copy1)
	sort.Ints(copy2)
	log.Println(prompt, equal(copy1, copy2), d)
	if debug {
		log.Println("sorted", copy1)
	}
	return d
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sortMapKeys(buckets map[int][]int) []int {
	keys := make([]int, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
