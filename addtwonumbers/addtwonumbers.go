package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (L *ListNode) PushFront(dig int) *ListNode {
	if L != nil {
		return (&ListNode{dig, L})
	} else {
		return (&ListNode{dig, nil})
	}
}

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

func getNumber(list *ListNode) int {
	var dig []int
	var sum int
	for ; list != nil; list = list.Next {
		dig = append(dig, list.Val)
	}
	for i := len(dig) - 1; i >= 0; i-- {
		sum = sum*10 + dig[i]
	}
	return (sum)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmpList *ListNode
	first, second := getNumber(l1), getNumber(l2)
	sum := first + second
	for sum != 0 {
		tmp := sum % 10
		sum /= 10
		tmpList = tmpList.PushBack(tmp)
	}
	return (tmpList)
}
