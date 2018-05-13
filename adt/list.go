package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func NewList(n int) *ListNode {
	var node *ListNode
	for i := n; i > 0; i-- {
		if node == nil {
			node = &ListNode{Value: i}
		} else {
			node = &ListNode{Value: i, Next: node}
		}
	}
	return node
}

// 链表逆序
func ListReverse(head *ListNode) *ListNode {
	var (
		pre *ListNode = nil
		cur *ListNode = head
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 链表插入节点
func ListInsert(head, node *ListNode, i int) *ListNode {
	var (
		ni  = 0
		pre *ListNode
	)
	for cur := head; cur != nil; cur = cur.Next {
		if ni == i {
			// 新插入节点为头节点
			if pre == nil {
				node.Next = cur
				return node
			}
			node.Next = cur
			pre.Next = node
			return head
		}
		pre = cur
		ni++
	}
	// 在尾部插入
	if i == ni {
		pre.Next = node
	}
	return head
}

func ListDelete(head *ListNode, i int) *ListNode {
	var (
		ni  = 0
		pre *ListNode
	)
	for cur := head; cur != nil; cur = cur.Next {
		if ni == i {
			// 删掉头节点
			if pre == nil {
				return cur.Next
			}
			pre.Next = cur.Next
			return head
		}
		pre = cur
		ni++
	}
	return head
}

// 链表检测是否有环，空间复杂度O(1)
// 双指针一快一慢遍历，若指针相遇，则代表有环
func ListContainsRing(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 返回环的入口节点，空间复杂度O(1)
// 双指针一快一慢遍历相遇时，重置一个指针到起始点，然后一次走一个节点，再次相遇即为环的入口
// 实际快指针比慢指针多走了一个环的节点个数才会相遇
func ListRingJoinNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast != slow {
		return nil
	}
	// 慢指针重置到起始点，然后快慢指针都每次只走一步，再次交汇处即为环的入口。
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 返回环的入口节点，空间复杂度O(1)
// 先求环的周长，然后两个指针再从头相隔周长个节点的距离出发，相遇时就是入口
func ListRingJoinNodeV2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow != fast {
		return nil
	}
	// 求环的周长，快指针不动，慢指针走一圈计数
	count, slow := 1, slow.Next
	for slow != fast {
		count++
		slow = slow.Next
	}
	// 知道环的周长，然后两个指针再次从头出发，并且相隔周长个节点，再次相遇必然是，正好快指针绕弯一圈，慢指针刚到入口
	slow, fast = head, head
	for ; count > 0; count-- {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	head := NewList(10)
	listDump("new", head)
	head = ListReverse(head)
	listDump("reverse", head)

	head = ListInsert(head, &ListNode{Value: 100}, 10)
	listDump("insert(value=100, pos=10)", head)
	head = ListInsert(head, &ListNode{Value: 100}, 0)
	listDump("insert(value=100, pos=0)", head)
	head = ListInsert(head, &ListNode{Value: 100}, 5)
	listDump("insert(value=100, pos=5)", head)

	head = ListDelete(head, 5)
	listDump("delete(value=100, pos=5)", head)
	head = ListDelete(head, 0)
	listDump("delete(value=100, pos=0)", head)
	head = ListDelete(head, 10)
	listDump("delete(value=100, pos=10)", head)

}

func listDump(prefix string, head *ListNode) {
	fmt.Println(strings.Repeat("-", 20))
	for node := head; node != nil; node = node.Next {
		fmt.Println(prefix, "=>", node.Value)
	}
}
