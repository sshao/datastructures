package linked_list

import "testing"
import "github.com/stretchr/testify/assert"
import "reflect"

func assertBothWays(t *testing.T, list *DoubleLinkedList, expected interface{}) {
  slice := reflect.ValueOf(expected)

  cur_node := list.Head
  i := 0

  for cur_node != nil {
    assert.Equal(t, slice.Index(i).Interface(), cur_node.Data)
    cur_node = cur_node.Next
    i++
  }

  cur_node = list.Head
  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }
  i = slice.Len() - 1

  for cur_node != nil {
    assert.Equal(t, slice.Index(i).Interface(), cur_node.Data)
    cur_node = cur_node.Prev
    i--
  }
}

func TestCreateDouble(t *testing.T) {
  list := NewDoubleLinkedList()

  assert.Nil(t, list.Head)
}

func TestInsertDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  list.insert("hello")

  assert.Equal(t, list.Head.Data, "hello")
}

func TestInsertLotsDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  expected := []int{1, 2, 3, 4, 5}

  for _,e := range expected {
    list.insert(e)
  }

  cur_node := list.Head

  i := 0
  for cur_node != nil {
    assert.Equal(t, cur_node.Data, expected[i])
    cur_node = cur_node.Next
    i++
  }

  cur_node = list.Head
  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }

  i = len(expected) - 1

  for cur_node != nil {
    assert.Equal(t, cur_node.Data, expected[i])
    cur_node = cur_node.Prev
    i--
  }
}

func TestDeleteFirstDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  insert := []int{1, 2, 3, 4, 5}
  expected := []int{2, 3, 4, 5}

  for _,e := range insert {
    list.insert(e)
  }

  list.delete(1)

  cur_node := list.Head
  i := 0

  for cur_node.Next != nil {
    assert.Equal(t, cur_node.Data, expected[i])
    cur_node = cur_node.Next
    i++
  }

  cur_node = list.Head
  for cur_node.Next != nil {
    cur_node = cur_node.Next
  }
  i = len(expected) - 1

  for cur_node != nil {
    assert.Equal(t, cur_node.Data, expected[i])
    cur_node = cur_node.Prev
    i--
  }
}

func TestDeleteLastDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  insert := []int{1, 2, 3, 4, 5}
  expected := []int{1, 2, 3, 4}

  for _,e := range insert {
    list.insert(e)
  }

  list.delete(5)

  assertBothWays(t, list, expected)
}

func TestDeleteMiddleDouble(t *testing.T) {
  list := NewDoubleLinkedList()
  insert := []int{1, 2, 3, 4, 5}
  expected := []int{1, 2, 4, 5}

  for _,e := range insert{
    list.insert(e)
  }

  list.delete(3)

  assertBothWays(t, list, expected)
}

func TestDeleteOnly(t *testing.T) {
  list := NewDoubleLinkedList()
  list.insert("hello")

  list.delete("hello")

  assert.Nil(t, list.Head)
}

func TestDeleteNonexistent(t *testing.T) {
  list := NewDoubleLinkedList()
  insert := []int{1, 2, 3, 4, 5}
  expected := insert

  for _,e := range insert{
    list.insert(e)
  }

  list.delete(9)

  assertBothWays(t, list, expected)
}
