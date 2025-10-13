/*
题意：删除链表中等于给定值 val 的所有节点。

示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]

示例 2： 输入：head = [], val = 1 输出：[]
*/
package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	pre := head
	index := head
	for index != nil {
		if index.Val == val {
			pre.Next = index.Next
		} else {
			pre = index
		}
		index = index.Next
	}

	return head
}
