package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	ListPushBack(link, "three")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "1")

// 	fmt.Println(ListLast(link))
// 	fmt.Println(ListLast(link2))
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head.Next != nil {
		l.Head = l.Head.Next
	}
	return l.Head.Data
}

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = n
// 		l.Tail = n
// 	} else {
// 		l.Tail.Next = n
// 		l.Tail = n
// 	}
// }
