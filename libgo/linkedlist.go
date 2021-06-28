package libgo

type List struct {
	head *ListNode
}

type ListNode struct {
	data interface{}
	next *ListNode
}

func (L *List) PushBack(data interface{}) {
	var new ListNode
	var tmp *ListNode
	if tmp != nil {
		tmp := L.head
		new.next = nil
		new.data = data
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &new
	} else {
		L.head = &ListNode{data, nil}
	}
}

// func (L *List) GetData(idx int) (interface{}, error) {
// var tmp ListNode = L.first
// for (i := 0; i < )
// }
