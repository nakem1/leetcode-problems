package main

import (
	"fmt"

	"github.com/nakem1/leetcode-problems/addtwonumbers"
)

func main() {
	tmp := addtwonumbers.AddTwoNumbers(&addtwonumbers.ListNode{5, &addtwonumbers.ListNode{8, nil}}, &addtwonumbers.ListNode{1, &addtwonumbers.ListNode{0, nil}})
	fmt.Printf("%v\n", tmp)
}
