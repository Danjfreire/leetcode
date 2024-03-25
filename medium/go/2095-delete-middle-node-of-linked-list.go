package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    length := 1
	cur := head

	for cur.Next != nil {
		cur = cur.Next
		length++
	}

	middle := length/2

	if middle == 0 {
		return nil
	}

	cur = head
	idx := 0

	for idx != middle-1 {
		cur = cur.Next
		idx++
	}

	removed := cur.Next
	cur.Next = removed.Next
	removed.Next = nil

	return head
}