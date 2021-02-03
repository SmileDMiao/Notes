/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var new *ListNode

	for head != nil {
		tmp := head.Next
		cur.next = new
		new = head
		current = tmp
	}
	return new
}
