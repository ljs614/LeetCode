package medium

//Definition for singly-linked list.

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers ...
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 처음 답
	// tempNode := &ListNode{}
	// result := tempNode
	// temp := 0
	// for {
	//     if l1 == nil {
	//         l1 = &ListNode{}
	//     }
	//     if l2 == nil {
	//         l2 = &ListNode{}
	//     }
	//     tempNode.Val = l1.Val + l2.Val + temp
	//     if tempNode.Val > 9 {
	//         tempNode.Val -= 10
	//         temp = 1
	//     } else {
	//         temp = 0
	//     }

	//     if l1.Next != nil || l2.Next != nil {
	//         tempNode.Next = &ListNode{}
	//         tempNode, l1, l2 = tempNode.Next, l1.Next, l2.Next
	//     } else if  temp == 1{
	//         tempNode.Next = &ListNode{}
	//         tempNode.Next.Val = 1
	//         break
	//     } else {
	//         break
	//     }
	// }
	// return result

	//참고 후 답
	tempNode := &ListNode{}
	result := tempNode
	sum := 0
	for l1 != nil || l2 != nil || sum == 1 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tempNode.Next = &ListNode{Val: sum % 10}
		sum /= 10
		tempNode = tempNode.Next
	}
	return result.Next
}
