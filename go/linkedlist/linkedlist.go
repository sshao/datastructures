package linkedlist

type Node struct {
  Data interface{}
  Next *Node
}

type LinkedList struct {
  Head *Node
}

func NewLinkedList() *LinkedList {
  return &LinkedList{Head: nil}
}

func (l *LinkedList) insert(data interface{}) *LinkedList {
  if l.Head == nil {
    l.Head = &Node{data, nil}
    return l
  }

  new_node := &Node{data, nil}
  cur_node := l.Head

  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }

  cur_node.Next = new_node
  return l
}

func (l *LinkedList) delete(data interface{}) *LinkedList {
  if l.Head.Data == data {
    cur_node := l.Head
    l.Head = cur_node.Next
    return l
  }

  prev_node := l.Head
  cur_node := l.Head

  for cur_node.Next != nil && cur_node.Data != data {
    prev_node = cur_node
    cur_node = cur_node.Next
  }

  prev_node.Next = cur_node.Next

  return l
}
