package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	tail := head

	for tail != nil && tail.Next != nil {
		middle = middle.Next
		tail = tail.Next.Next
	}

	return middle
}
