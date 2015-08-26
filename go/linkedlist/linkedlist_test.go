package linkedlist

import "testing"

func TestCreate(t *testing.T) {
  list := NewLinkedList()

  if list.Head != nil {
    t.Errorf("did not create an empty LinkedList")
  }
}

func TestInsertEmpty(t *testing.T) {
  list := NewLinkedList()
  list.insert("goodbye")

  if list.Head.Data != "goodbye" {
    t.Errorf("list.Head.Next.Data == %v, want %v", list.Head.Next.Data, "goodbye")
  }
}

func TestInsertLots(t *testing.T) {
  a := [5]interface{}{9, "four", 8, "three", 2}
  list := NewLinkedList()

  for _,e := range a {
    list.insert(e)
  }

  cur_node := list.Head

  for i,e := range a {
    if cur_node.Data != e {
      t.Errorf("data at index %v == %v, want %v", i, cur_node.Data, e)
    }
    cur_node = cur_node.Next
  }
}

// delete from ll with one node
func TestDeletionSingle(t *testing.T) {
  list := NewLinkedList()
  list.insert("one")
  list.delete("one")

  if list.Head != nil {
    t.Errorf("head node is not nil")
  }
}

// delete last node from ll with more than one node
func TestDeletionLast(t *testing.T) {
  list := NewLinkedList()
  list.insert("one")
  list.insert("two")
  list.delete("two")

  if list.Head.Next != nil {
    t.Errorf("list.head.Next is not nil")
  }
}

// delete first node from ll with more than one node
func TestDeletionFirst(t *testing.T) {
  list := NewLinkedList()
  list.insert("one")
  list.insert("two")
  list.delete("one")

  if list.Head.Data != "two" {
    t.Errorf("list.Head.Data == %v, want %v", list.Head.Data, "two")
  }
}

// delete node from middle of ll
func TestDeletionMiddle(t *testing.T) {
  list := NewLinkedList()
  list.insert("one")
  list.insert("two")
  list.insert("three")
  list.insert("four")
  list.insert("five")
  list.delete("three")

  expected := [4]string{"one", "two", "four", "five"}

  cur_node := list.Head

  for i,e := range expected {
    if cur_node.Data != e {
      t.Errorf("data at index %v == %v, want %v", i, cur_node.Data, e)
    }
    cur_node = cur_node.Next
  }
}

func TestDeletionFromEmpty(t *testing.T) {
  list := NewLinkedList()
  list.delete("nothing")

  if list.Head != nil {
    t.Errorf("list.Head not nil")
  }
}

func TestDeletionOfMissingElement(t *testing.T) {
  list := NewLinkedList()

  expected := [3]string{"one", "two", "three"}

  for _,e := range expected {
    list.insert(e)
  }

  list.delete("four")

  cur_node := list.Head

  for i,e := range expected {
    if cur_node.Data != e {
      t.Errorf("data at index %v == %v, want %v", i, cur_node.Data, e)
    }
    cur_node = cur_node.Next
  }
}
