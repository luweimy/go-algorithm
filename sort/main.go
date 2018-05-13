package main

import (
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

// BucketSort 桶排序 O(N+M) 稳定
// 将待排序数据分别放入对应的桶中，放入桶中时进行插入排序，最后将桶中数据依次放入待排序数组。
func BucketSort(slice []int) {
	buckets := make(map[int][]int, len(slice)/10)
	for _, n := range slice {
		bi := bucketAt(n)
		bucket := buckets[bi]
		// bucket是已排序的桶，此处内嵌插入排序
		i := 0
		for ; i < len(bucket); i++ {
			if n < bucket[i] {
				bucket = insertAt(bucket, i, n)
				break
			}
		}
		if i == len(bucket) {
			bucket = append(bucket, n)
		}
		buckets[bi] = bucket
	}
	var i = 0
	for _, k := range mapKeys(buckets) {
		bucket := buckets[k]
		copy(slice[i:i+len(bucket)], bucket)
		i += len(bucket)
	}
}

// BucketSortV2 桶排序 O(N+M) 稳定
// 将待排序数据分别放入对应的桶中，最后将桶中数据分别进行插入排序，然后依次放入待排序数组。
func BucketSortV2(slice []int) {
	buckets := make(map[int][]int, len(slice)/10)
	for _, n := range slice {
		bi := bucketAt(n)
		bucket := buckets[bi]
		bucket = append(bucket, n)
		buckets[bi] = bucket
	}
	var i = 0
	for _, k := range mapKeys(buckets) {
		bucket := buckets[k]
		InsertionSort(bucket)
		copy(slice[i:i+len(bucket)], bucket)
		i += len(bucket)
	}
}

func bucketAt(n int) int {
	return n / 10
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

func merge(s, s1, s2 []int) {
	if len(s) != len(s1)+len(s2) {
		panic(-1)
	}
	for i, i1, i2 := 0, 0, 0; i < len(s); i++ {
		// i1>=len(s1)&&i2>=len(s2)的情况必然i>=len(s)，所以不会运行到此处
		if i1 >= len(s1) {
			copy(s[i:], s2[i2:])
			break
		}
		if i2 >= len(s2) {
			copy(s[i:], s1[i1:])
			break
		}
		if s1[i1] <= s2[i2] {
			s[i] = s1[i1]
			i1++
		} else {
			s[i] = s2[i2]
			i2++
		}
	}
}

// QuickSort 快排 O(NlogN) 不稳定
// 通过一趟排序将要排序的数据分割成两部分，其中一部分的所有数据都比另外一部分的所有数据小，然后再按此方法对这两部分数据分别递归进行，以此达到整个数据有序。
func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	// p(pivot)是基准值位置，l(left)是左侧哨兵位置，r(right)是右侧哨兵位置
	// 此处直接用最左边的值作为基准值，更合理的方法是最左，最右，中间三个数，取中间的值作为基准值。
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

func main() {
	unsorted := randomInts(9)

	withSort("InsertionSort=>", InsertionSort, copySlice(unsorted))
	withSort("BubbleSort=>", BubbleSort, copySlice(unsorted))
	withSort("SelectionSort=>", SelectionSort, copySlice(unsorted))
	withSort("BucketSort=>", BucketSort, copySlice(unsorted))
	withSort("BucketSortV2=>", BucketSortV2, copySlice(unsorted))
	withSort("MergeSort=>", MergeSort, copySlice(unsorted))
	withSort("QuickSort=>", QuickSort, copySlice(unsorted))
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

func withSort(prompt string, sortfn func([]int), slice []int) time.Duration {
	t := time.Now()
	sortfn(slice)
	d := time.Now().Sub(t)
	log.Println(prompt, sort.IntsAreSorted(slice), d)
	if !sort.IntsAreSorted(slice) {
		log.Println(prompt, slice)
	}
	return d
}

func insertAt(bucket []int, i, n int) []int {
	bucket = append(bucket, n)
	copy(bucket[i+1:], bucket[i:])
	bucket[i] = n
	return bucket
}

func mapKeys(buckets map[int][]int) []int {
	keys := make([]int, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
