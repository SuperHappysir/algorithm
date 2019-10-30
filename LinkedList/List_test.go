package LinkedList

import (
    "github.com/davecgh/go-spew/spew"
    "testing"
)

func TestLinkedList_GetNode(t *testing.T) {
    list := LinkedList{}
    list.InsertFirst(1)
    if list.GetNode(1).value != 1 {
        t.Fail()
    }
}

func TestLinkedList_InsertFirst(t *testing.T) {
    list := LinkedList{}
    list.InsertFirst(1)
    list.InsertFirst(2)
    if list.GetNode(2).value != 2 {
        t.Fail()
    }
    if list.head.next.value != 1 {
        t.Fail()
    }
    if list.head.value != 2 {
        t.Fail()
    }
}

func TestLinkedList_InsertLast(t *testing.T) {
    list := LinkedList{}
    list.InsertLast(1)
    list.InsertLast(2)
    if list.GetNode(2).value != 2 {
        t.Fail()
    }
    if list.head.next.value != 2 {
        t.Fail()
    }
    if list.head.value != 1 {
        t.Fail()
    }
}

func TestLinkedList_RemoveByValue(t *testing.T) {
    list := LinkedList{}
    list.InsertFirst(1)
    list.InsertFirst(2)
    list.InsertFirst(3)
    if list.RemoveByValue(1).value != 1 {
        t.Fail()
    }
    if list.GetNode(1) != nil {
        t.Fail()
    }
    if list.head == nil {
        t.Fail()
    }
    if list.head.value != 3 {
        spew.Dump(list)
        t.Fail()
    }
}
