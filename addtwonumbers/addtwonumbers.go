package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// func (L *ListNode) PushFront(dig int) *ListNode {
// 	if L != nil {
// 		return (&ListNode{dig, L})
// 	} else {
// 		return (&ListNode{dig, nil})
// 	}
// }

func (L *ListNode) PushBack(dig int) *ListNode {
	if tmp := L; tmp != nil {
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = &ListNode{dig, nil}
	} else {
		L = &ListNode{dig, nil}
	}
	return (L)
}

// func getNumber(list *ListNode) int {
// 	var dig []int
// 	var sum int
// 	for ; list != nil; list = list.Next {
// 		dig = append(dig, list.Val)
// 	}
// 	for i := len(dig) - 1; i >= 0; i-- {
// 		sum = sum*10 + dig[i]
// 	}
// 	return (sum)
// }

// func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var tmpList *ListNode
// 	var digL1, digL2 []int
// 	var sumL1, sumL2 int64
// 	for tmpL1, tmpL2 := l1, l2; tmpL1 != nil || tmpL2 != nil; {
// 		if tmpL1 != nil {
// 			digL1 = append(digL1, tmpL1.Val)
// 			tmpL1 = tmpL1.Next
// 		}
// 		if tmpL2 != nil {
// 			digL2 = append(digL2, tmpL2.Val)
// 			tmpL2 = tmpL2.Next
// 		}
// 	}
// 	for i, j := len(digL1)-1, len(digL2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
// 		if i >= 0 {
// 			sumL1 = sumL1*10 + int64(digL1[i])
// 		}
// 		if j >= 0 {
// 			sumL2 = sumL2*10 + int64(digL2[j])
// 		}
// 	}
// 	sum := sumL1 + sumL2
// 	if sum == 0 {
// 		return (&ListNode{0, nil})
// 	}
// 	for sum != 0 {
// 		dig := sum % 10
// 		sum /= 10
// 		if tmp := tmpList; tmp != nil {
// 			for tmp.Next != nil {
// 				tmp = tmp.Next
// 			}
// 			tmp.Next = &ListNode{int(dig), nil}
// 		} else {
// 			tmpList = &ListNode{int(dig), nil}
// 		}
// 	}
// 	return (tmpList)
// }

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	x := &ListNode{}
	root := x
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		x.Next = &ListNode{Val: sum % 10}
		sum = sum / 10
		x = x.Next
	}
	return root.Next
}
