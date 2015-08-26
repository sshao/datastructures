package linked_list

type DoubleNode struct {
  Data interface{}
  Next *DoubleNode
  Prev *DoubleNode
}

type DoubleLinkedList struct {
  Head *DoubleNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
  return &DoubleLinkedList{Head: nil}
}

func (l *DoubleLinkedList) insert(data interface{}) {
  new_node := &DoubleNode{Data: data, Next: nil, Prev: nil}

  if l.Head == nil {
    l.Head = new_node
    return
  }

  cur_node := l.Head

  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }

  new_node.Prev = cur_node
  cur_node.Next = new_node
}
