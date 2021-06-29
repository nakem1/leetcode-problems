package main

import "github.com/nakem1/leetcode-problems/addtwonumbers"

func main() {
	var list *addtwonumbers.ListNode
	var list1 *addtwonumbers.ListNode

	list = list.PushBack(9)
	list = list.PushBack(9)
	list = list.PushBack(9)
	list = list.PushBack(9)
	list = list.PushBack(9)
	list = list.PushBack(9)
	list = list.PushBack(9)

	list1 = list1.PushBack(9)
	list1 = list1.PushBack(9)
	list1 = list1.PushBack(9)
	list1 = list1.PushBack(9)
	addtwonumbers.AddTwoNumbers(list, list1)
}
