package main

import (
	"container/list"
	"fmt"
)

// 二叉树节点
type BinaryTreeNode struct {
	Value interface{}
	L     *BinaryTreeNode
	R     *BinaryTreeNode
}

func NewBinaryTree(values []int) *BinaryTreeNode {
	if len(values) == 0 {
		return nil
	}
	var (
		q    = list.New()
		i    = 0
		root = &BinaryTreeNode{Value: values[i]}
	)

	q.PushBack(root)
	i++

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		node := e.Value.(*BinaryTreeNode)
		if i < len(values) {
			node.L = &BinaryTreeNode{Value: values[i]}
			q.PushBack(node.L)
			i++
		}
		if i < len(values) {
			node.R = &BinaryTreeNode{Value: values[i]}
			q.PushBack(node.R)
			i++
		}
	}
	return root
}

// 前序深度优先递归遍历(Depth First Search)
func WalkPreOrder(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if node == nil {
		return
	}
	visit(node)
	WalkPreOrder(node.L, visit)
	WalkPreOrder(node.R, visit)
}

// 中序深度优先递归遍历(Depth First Search)
func WalkInOrder(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if node == nil {
		return
	}
	WalkInOrder(node.L, visit)
	visit(node)
	WalkInOrder(node.R, visit)
}

// 后序深度优先递归遍历(Depth First Search)
func WalkPostOrder(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if node == nil {
		return
	}
	WalkPostOrder(node.L, visit)
	WalkPostOrder(node.R, visit)
	visit(node)
}

// 前序深度优先非递归遍历(Depth First Search)
func WalkPreOrderV2(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	var stack = list.New()
	stack.PushBack(node)
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*BinaryTreeNode)
		if node.R != nil {
			stack.PushBack(node.R)
		}
		if node.L != nil {
			stack.PushBack(node.L)
		}
		visit(node)
	}
}

// 中序深度优先非递归遍历(Depth First Search)
func WalkInOrderV2(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	var stack = list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			// 把左侧子树全部入栈
			stack.PushBack(node)
			node = node.L
		}
		if stack.Len() > 0 {
			e := stack.Back()
			stack.Remove(e)
			node = e.Value.(*BinaryTreeNode)
			visit(node)
			node = node.R
		}
	}
}

// 后序深度优先非递归遍历(Depth First Search)
func WalkPostOrderV2(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	var (
		stack                 = list.New()
		pre   *BinaryTreeNode = nil
	)
	for node != nil || stack.Len() > 0 {
		for node != nil {
			// 把左侧子树全部入栈
			stack.PushBack(node)
			node = node.L
		}
		e := stack.Back()
		n := e.Value.(*BinaryTreeNode)
		// 无右子树，或者已经遍历过右子树，才能访问当前节点
		if n.R == nil || n.R == pre {
			pre = n
			visit(n)
			node = nil
			stack.Remove(e)
		} else {
			node = n.R
		}
	}
}

// 广度优先遍历(Breadth First Search)
func WalkBFS(node *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	var q = list.New()
	q.PushBack(node)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		node := e.Value.(*BinaryTreeNode)
		if node.L != nil {
			q.PushBack(node.L)
		}
		if node.R != nil {
			q.PushBack(node.R)
		}
		visit(node)
	}
}

// 前序中序推倒后序
func CalcPostOrder(pre, in []int) []int {
	if len(pre) != len(in) {
		panic(-1)
	}
	if len(pre) == 0 {
		return []int{}
	}
	if len(pre) == 1 {
		return []int{pre[0]}
	}
	var post = make([]int, 0, len(pre))
	var root = pre[0]
	var index = sliceIndex(in, root)
	if index < 0 {
		panic(-1)
	}

	var l = CalcPostOrder(pre[1:index+1], in[:index])
	var r = CalcPostOrder(pre[index+1:], in[index+1:])
	post = append(post, l...)
	post = append(post, r...)
	return append(post, root)
}

func sliceIndex(s []int, t int) int {
	for i, e := range s {
		if e == t {
			return i
		}
	}
	return -1
}

// 返回值（是否平衡，以当前为根的最大深度计数）
func checkBalance(node *BinaryTreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	leftBalance, leftDepth := checkBalance(node.L)
	rightBalance, rightDepth := checkBalance(node.R)

	depth := maxInt(leftDepth, rightDepth) + 1
	if !leftBalance || !rightBalance {
		return false, depth
	}
	if absInt(leftDepth-rightDepth) <= 1 {
		return true, depth
	}
	return false, depth
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}

	btree := NewBinaryTree(nums)
	WalkPreOrder(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkPreOrder=>", node.Value)
	})
	WalkPreOrderV2(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkPreOrderV2=>", node.Value)
	})
	WalkInOrder(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkInOrder=>", node.Value)
	})
	WalkInOrderV2(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkInOrderV2=>", node.Value)
	})
	WalkPostOrder(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkPostOrder=>", node.Value)
	})
	WalkPostOrderV2(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkPostOrderV2=>", node.Value)
	})
	WalkBFS(btree, func(node *BinaryTreeNode) {
		fmt.Println("WalkBFS=>", node.Value)
	})
	// 4,3,2,8,9,7,6,5,1
	fmt.Println("CalcPostOrder=>", CalcPostOrder([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{2, 4, 3, 1, 5, 8, 7, 9, 6}))
}
