package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		count++
		l = l.Next
	}
	return nil
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
