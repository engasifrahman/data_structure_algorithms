package main

import "fmt"

type Node struct {
	data  int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert at the beginning of the doubly linked list
func (list *DoublyLinkedList) insertAtBeginning(data int) {
	newNode := &Node {
		data: data, 
		prev: nil, 
		next: list.head,
	}

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode

	if list.tail == nil {
		list.tail = newNode
	}
}

// Insert at the end of the doubly linked list
func (list *DoublyLinkedList) insertAtEnd(data int) {
	newNode := &Node {
		data: data, 
		prev: list.tail, 
		next: nil,
	}

	if list.tail != nil {
		list.tail.next = newNode
	}

	list.tail = newNode

	if list.head == nil {
		list.head = newNode
	}
}

// Insert after a given node in the doubly linked list
func (list *DoublyLinkedList) insertAfter(node *Node, data int) {
	if node == nil {
		fmt.Println("Invalid node provided")
		return
	}

	newNode := &Node {
		data: data, 
		prev: node, 
		next: node.next,
	}

	if node.next != nil {
		node.next.prev = newNode
	}

	node.next = newNode

	if list.tail == node {
		list.tail = newNode
	}
}

// Insert before a given node in the doubly linked list
func (l *DoublyLinkedList) insertBefore(node *Node, data int) {
	if node == nil {
		fmt.Println("Invalid node provided!")
		return
	}

	newNode := &Node{
		data: data,
		prev: node.prev,
		next: node,
	}

	if node.prev != nil {
		node.prev.next = newNode
	}

	node.prev = newNode

	if l.head == node {
		l.head = newNode
	}
}

// Delete a node from the doubly linked list
func (list *DoublyLinkedList) deleteNode(node *Node) {
	if node == nil {
		fmt.Println("Invalid node provided")
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
}

// Delete a node by data from the doubly linked list
func (list *DoublyLinkedList) deleteNodeByData(data int) {
	current := list.head

	for current != nil {
		if current.data == data {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				list.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				list.tail = current.prev
			}

			// Free up the memory (optional in Go due to garbage collection)
			// current.prev = nil
			// current.next = nil
			return
		}

		current = current.next
	}
}

// Traverse forward the doubly linked list in forward order
func (list *DoublyLinkedList) traverseForward() {
	fmt.Print("Forward traversal: ")

	current := list.head
	for current != nil {
		fmt.Print(" →", current.data)
		current = current.next
	}
	fmt.Println()
}

// Traserve backward the doubly linked list in reverse order
func (list *DoublyLinkedList) traverseBackward() {
	fmt.Print("Backward traversal: ")

	current := list.tail
	for current != nil {
		fmt.Print(" →", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	list := DoublyLinkedList{}

	list.insertAtEnd(70)
	list.insertAtBeginning(60)
	list.insertAtBeginning(30)
	list.insertAtEnd(80)
	list.insertBefore(list.head.next, 40)
	list.insertAfter(list.head.next, 50)

	list.traverseForward()
	list.traverseBackward()

	list.deleteNode(list.head.next)
	list.deleteNodeByData(70)

	fmt.Println("\n------- After delete operation -------")
	list.traverseForward()
	list.traverseBackward()
}
