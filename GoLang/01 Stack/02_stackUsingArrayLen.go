package main

import "fmt"

type Stack struct {
    items []int
    capacity int
}

func newStack(capacity int) Stack {
    return Stack{
        items: make([]int, 0, capacity),
        capacity: capacity,
    }
}

// Add an element to the top of a stack
func (s *Stack) push(item int) bool {
    if s.isFull() {
        fmt.Println("Stack is full, unable to push element:", item)
        return false
    }

    s.items = append(s.items, item)

	fmt.Printf("%d has pushed to the stack and updated stack is %v \n", item, s.items)

	return true
}

// Remove an element from the top of a stack
func (s *Stack) pop() (int, bool) {
    item := 0

    if s.isEmpty() {
        fmt.Println("Unable to pop element as the stack is empty!")
        return item, false
    }

	index := len(s.items) - 1
    item, s.items = s.items[index], s.items[:index]

	fmt.Printf("%d has popped from the stack and updated stack is %v \n\n", item, s.items)

    return item, true
}

// Get the value of the top element without removing it
func (s *Stack) peek() (int, bool) {
	if s.isEmpty() {
		fmt.Println("There has no peek element as the stack is empty!")
		return 0, false
	}

	item := s.items[len(s.items) - 1]

	fmt.Printf("%d is the peek item of the stack and the stack is %v \n", item, s.items)

    return item, true
}

// Check if the stack is empty
func (s *Stack) isEmpty() bool {
    return  len(s.items) == 0
}

// Check if the stack is full
func (s *Stack) isFull() bool {
    return  len(s.items) == s.capacity
}

func main() {
    stack := newStack(5)

    stack.push(1)
    stack.push(2)
    stack.push(3)
    stack.push(4)
    stack.push(5)
    stack.push(6) // Trying to push element to a full stack

	fmt.Println("------------------------")

    stack.peek() // 5
    stack.pop() // 5

    stack.peek() // 4
    stack.pop() // 4

    stack.peek() // 3
    stack.pop() // 3

    stack.peek() // 2
    stack.pop() // 2

    stack.peek() // 1
    stack.pop() // 1

    stack.peek() // Trying to get peek element from an empty stack
    stack.pop() // Trying to pop element from an empty stack
    
	fmt.Println("------------------------")

    stack.push(1)
    stack.push(2)
    stack.push(3)
    stack.push(4)
    stack.push(5)
    stack.push(6) // Trying to push element to a full stack
}