package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newnode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newnode
		l.Tail = l.Head
	} else {
		l.Tail.Next = newnode
		l.Tail = l.Tail.Next
	}
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, "man")
// 	ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }
