package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var valList []int
	for head != nil {
		valList = append(valList, head.Val)
		head = head.Next
	}
	maxVal, n := 0, len(valList)
	for i := 0; i < n/2; i++ {
		maxVal = max(maxVal, valList[i]+valList[n-i-1])
	}
	return maxVal
}

func pairSum2(head *ListNode) int {
	// 这里的fast必须是head.Next
	fast, slow := head.Next, head
	// get middle node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil
	fast = reverse(fast)
	res := head.Val + fast.Val
	head = head.Next
	fast = fast.Next
	for head != nil && fast != nil {
		tmp := head.Val + fast.Val
		res = max(res, tmp)
		head = head.Next
		fast = fast.Next
	}
	return res

}

func reverse(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		p := head
		head = head.Next
		p.Next = res
		res = p
	}
	return res
}

func main() {

	valList := []int{4, 2, 2, 3, 6, 7, 9, 120}
	head := &ListNode{Val: valList[0]}
	tail := head
	for i := 1; i < len(valList); i++ {
		tail.Next = &ListNode{Val: valList[i]}
		tail = tail.Next
	}
	test := pairSum(head)
	fmt.Println(test)
}
