package data_structure

type LinkedList struct {
	Header *Node
	Tail   *Node
	Count  int
}

type Node struct {
	Next  *Node
	Value interface{}
}
