package lists

import (
	"errors"
)

type singlyLinkedListElement struct {
	value int // interface {} ?
	next  *singlyLinkedListElement
}

type SinglyLinkedList struct {
	head *singlyLinkedListElement
}

func (sll *SinglyLinkedList) Push(value int) {
	if sll.head == nil {
		sll.head = &singlyLinkedListElement{value: value}
	} else {
		sll.head = &singlyLinkedListElement{value: value, next: sll.head}
	}
}

func (sll *SinglyLinkedList) Pop() (int, error) {
	if sll.head == nil {
		return 0, errors.New("empty singly linked list")
	}

	value := sll.head.value
	sll.head = sll.head.next

	return value, nil
}
