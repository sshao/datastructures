package linked_list

import "testing"

func TestCreateDouble(t *testing.T) {
  list := NewDoubleLinkedList()

  if list.Head != nil {
    t.Errorf("did not create an empty LinkedList")
  }
}

func TestInsertDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  list.insert("hello")

  if list.Head.Data != "hello" {
    t.Errorf("list.Head.Data == %v, expected %v", list.Head.Data, "hello")
  }
}

func TestInsertLotsDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  expected := []int{1, 2, 3, 4, 5}

  for _,e := range expected {
    list.insert(e)
  }

  cur_node := list.Head

  for i,e := range expected {
    if cur_node.Data != e {
      t.Errorf("cur_node.Data at index %v == %v, expected %v", i, cur_node.Data, e)
    }
    cur_node = cur_node.Next
  }

  cur_node = list.Head
  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }

  for i := len(expected)-1; i >= 0; i-- {
    if cur_node.Data != expected[i] {
      t.Errorf("cur_node.Data at index %v == %v, expected %v", i, cur_node.Data, expected[i])
    }
    cur_node = cur_node.Prev
  }
}
