package main

import "fmt"

type ElementType int
type Callback func(n *Node)

type Node struct {
	value ElementType
	next  *Node
}

type List struct {
	head *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node[val:%d, next:%p]", n.value, n.next)
}

func (l *List) String() string {
	s := "List{{ "

	if !l.Empty() {
		scout := l.head
		for i := 0; scout != nil; i++ {
			if i > 0 {
				s += ", "
			}

			s += fmt.Sprintf("%v", scout)
			scout = scout.next
		}
	}

	return s + " }}"
}

// Constructor of new List
func NewList() *List {
	return &List{nil}
}

// Is this List empty?
func (l *List) Empty() bool {
	if l != nil {
		return l.head == nil
	}
	return true
}

// Add a new node at the end of the List
func (l *List) Append(val ElementType) {
	newNode := &Node{val, nil}

	if l.Empty() {
		l.head = newNode
	} else {
		scout := l.head
		for scout.next != nil {
			scout = scout.next
		}
		scout.next = newNode
	}
}

// Attempt to find a value in a List
func (l *List) Contains(val ElementType) bool {
	if l.Empty() {
		return false
	}

	for scout := l.head; scout != nil; scout = scout.next {
		if scout.value == val {
			return true
		}
	}

	return false
}

// return the index of the element val. If it doesn't exist return -1
func (l *List) Index(val ElementType) int {
	if l.Contains(val) {
		for i, scout := 0, l.head; scout != nil; i, scout = i+1, scout.next {
			if scout.value == val {
				return i
			}
		}
	}
	return -1
}

// return the index of the last instance of element val. If it doesn't exist return -1
func (l *List) LastIndexOf(val ElementType) int {
	if l.Contains(val) {
		index := 0
		for i, scout := 0, l.head; scout != nil; i, scout = i+1, scout.next {
			if scout.value == val {
				index = i
			}
		}
		return index
	}
	return -1
}

// Insert a Node after certain value. Return true if it was added, false otherwise
func (l *List) Insert(val, after ElementType) bool {
	if !l.Contains(after) {
		return false
	}

	newNode := &Node{val, nil}
	for scout := l.head; scout.next != nil; scout = scout.next {
		if scout.value == after {
			newNode.next = scout.next
			scout.next = newNode
			break
		}
	}
	return true
}

// Return the Node with value val
func (l *List) Get(val ElementType) *Node {
	if l.Contains(val) {
		for scout := l.head; scout != nil; scout = scout.next {
			if scout.value == val {
				return scout
			}
		}
	}
	return nil
}

func (l *List) Peek() ElementType {
	if l.Empty() {
		return -1
	}
	return l.head.value
}

func (l *List) Push(val ElementType) {
	n := &Node{val, l.head}
	l.head = n
}

func (l *List) ForEach(f Callback) {
	for scout := l.head; scout != nil; scout = scout.next {
		f(scout)
	}
}

func (l *List) RemoveFirst(val ElementType) {
	// Don't do anything on an empty list
	if l.Empty() {
		return
	}

	// If it's the first node
	if l.head.value == val {
		l.head = l.head.next
		return
	}

	// find the node and then remove it
	for behind, ahead := l.head, l.head.next; ahead != nil; behind, ahead = ahead, ahead.next {
		if ahead.value == val {
			behind.next = ahead.next
			return
		}
	}
}

func (l *List) Remove(val ElementType) {
	// Don't do anything on an empty list
	if l.Empty() {
		return
	}

	// If it's the first node
	if l.head.value == val {
		l.head = l.head.next
		return
	}

	// find the node and then remove it
	for behind, ahead := l.head, l.head.next; ahead != nil; {
		if ahead.value == val {
			behind.next = ahead.next
			ahead = ahead.next
		} else {
			behind, ahead = behind.next, ahead.next
		}
	}
}

func (l *List) Len() int {
	i := 0

	for scout := l.head; scout != nil; scout = scout.next {
		i++
	}

	return i
}

func (l *List) ToSlice() []ElementType {
	var ret []ElementType

	for scout := l.head; scout != nil; scout = scout.next {
		ret = append(ret, scout.value)
	}

	return ret
}

func (l *List) ToArray() []ElementType {
	ret := make([]ElementType, l.Len())
	for i, scout := 0, l.head; scout != nil; i, scout = i+1, scout.next {
		ret[i] = scout.value
	}
	return ret
}

func main() {
	l := NewList()

	// Test Empty
	fmt.Println(l)
	fmt.Println(l.Empty())

	// Test Append once, no longer empty
	l.Append(1)
	fmt.Println(l)
	fmt.Println(l.Empty())

	// Test Appending Multiple
	l.Append(2)
	l.Append(3)
	l.Append(5)
	fmt.Println(l)
	fmt.Println(l.Empty())

	// Test Contains
	fmt.Println("l contains 0?", l.Contains(0))
	fmt.Println("l contains 3?", l.Contains(3))

	// Test Insert
	l.Insert(4, 3)
	fmt.Println(l)

	fmt.Println(l.Get(2))
	fmt.Println(l.Get(42))

	fmt.Println(l.Index(2))

	l.Append(2)
	fmt.Println(l.LastIndexOf(2))

	l.Push(0)
	fmt.Println(l)

	l.RemoveFirst(3)
	fmt.Println(l)

	l.Append(3)
	l.Append(3)
	l.Append(3)
	l.Append(3)
	l.Append(3)
	fmt.Println(l)
	l.Remove(3)
	fmt.Println(l)

	fmt.Println(l.Len())

	fmt.Println(l.ToSlice())
	fmt.Println(l.ToArray())

	l.ForEach(func(n *Node) {
		fmt.Println("Hello I'm a node with value:", n.value)
	})
}
