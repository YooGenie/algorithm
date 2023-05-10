package data_structure

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func Linked() {
	const node1 = NewNode{LinkedList{nil, nil, 0}, Value: 3}
}
