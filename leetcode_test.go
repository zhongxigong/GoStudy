package main

import (
	"fmt"
	"testing"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func TestArraySum(t *testing.T) {
	nums := []int{1, 5, 6, 2, 7, 11, 15}

	target := 26

	for index := 0; index < len(nums); index++ {
		fmt.Println(index)
		for innerIndex := index + 1; innerIndex < len(nums); innerIndex++ {
			if nums[index]+nums[innerIndex] == target {
				fmt.Println([]int{index, innerIndex})
				break
			}
		}
	}
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type Node struct {
	value int
	next  *Node
}

type LinkList struct {
	head  *Node
	last  *Node
	count int
}

func New() *LinkList {
	return &LinkList{
		head:  nil,
		last:  nil,
		count: 0,
	}
}

func (ll *LinkList) Append(value int) {
	node := &Node{value: value, next: nil}
	if ll.head == nil {
		ll.head = node
		ll.last = node
		ll.count += 1
	} else {
		ll.last.next = node
		ll.count += 1
		ll.last = node
	}
}

func (ll *LinkList) Print() {
	node := ll.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.value)
		if ll.last != node {
			fmt.Print("->")
		}
		node = node.next
	}
	fmt.Println()
}

func TestAddTwoNumbers(t *testing.T) {

	fmt.Scan()
	firstList := New()
	firstList.Append(2)
	firstList.Append(4)
	firstList.Append(3)

	secondList := New()
	secondList.Append(5)
	secondList.Append(6)
	secondList.Append(4)

	fNode := firstList.head
	sNode := secondList.head

	result := New()
	tmp := false
	for {
		if fNode == nil || sNode == nil {
			break
		}
		sum := fNode.value + sNode.value
		if tmp {
			sum = sum + 1
		}
		item := sum % 10
		result.Append(item)

		if sum >= 10 {
			tmp = true
		} else {
			tmp = false
		}
		fmt.Println(item)
		if fNode != nil || sNode != nil {
			fNode = fNode.next
			sNode = sNode.next
		}
	}
	result.Print()
}

func TestHelloScan(t *testing.T) {

}
