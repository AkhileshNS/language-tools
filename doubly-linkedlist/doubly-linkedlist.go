/*
 Note. All functions are in-order

 ** Doubly Linked List
 ** New - new linkedlist{} and .initialize(value)
 ** Insert (at start) - .prepend(value)
 ** Insert (at end) - .append(value)
 ** Insert (at middle) - .insert(index, value)
 ** Lookup - .get(index)
 ** Remove - .remove(index)
*/

// =============================================

package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (ll *linkedlist) initialize(value int) {
	n := node{value: value, next: nil, prev: nil}
	ll.head = &n
	ll.tail = &n
	ll.length = 1
}

func (ll linkedlist) isEmpty() bool {
	if ll.head == nil {
		fmt.Println("List is empty")
		return true
	}
	return false
}

func (ll *linkedlist) append(value int) {
	if ll.isEmpty() {
		ll.initialize(value)
		return
	}

	n := node{value: value, next: nil, prev: nil}
	ll.tail.next = &n
	n.prev = ll.tail
	ll.tail = &n
	ll.length++
}

func (ll *linkedlist) prepend(value int) {
	if ll.isEmpty() {
		ll.initialize(value)
		return
	}

	n := node{value: value, next: nil}
	n.next = ll.head
	ll.head.prev = &n
	ll.head = &n
	ll.length++
}

func (ll *linkedlist) print() {
	if !ll.isEmpty() {
		curr := ll.head
		nodes := []int{curr.value}
		curr = curr.next
		for curr != nil {
			nodes = append(nodes, (*curr).value)
			curr = curr.next
		}
		fmt.Println(nodes)
	}
}

func (ll *linkedlist) printBackwards() {
	if !ll.isEmpty() {
		curr := ll.tail
		nodes := []int{curr.value}
		curr = curr.prev
		for curr != nil {
			nodes = append(nodes, (*curr).value)
			curr = curr.prev
		}
		fmt.Println(nodes)
	}
}

func (ll *linkedlist) traverse(index int) (*node, *node, *node) {
	if ll.isEmpty() {
		return nil, nil, nil
	}

	curr := ll.head
	if index <= (ll.length / 2) {
		for i := 0; i < index && i < ll.length-1; i++ {
			curr = curr.next
		}
	} else {
		curr = ll.tail
		for i := ll.length - 1; i > index && i > 0; i-- {
			curr = curr.prev
		}
	}

	return curr.prev, curr, curr.next
}

func (ll *linkedlist) get(index int) (int, bool) {
	if !ll.isEmpty() {
		_, curr, _ := ll.traverse(index)
		return curr.value, true
	}

	return 0, false
}

func (ll *linkedlist) insert(index, value int) {
	if ll.isEmpty() {
		ll.initialize(value)
		return
	}
	n := node{value: value, next: nil}

	_, curr, leader := ll.traverse(index)
	curr.next = &n
	n.prev = curr
	n.next = leader
	if leader != nil {
		leader.prev = &n
	} else {
		ll.tail = &n
	}
	ll.length++
}

func (ll *linkedlist) remove(index int) {
	if !ll.isEmpty() {
		follower, curr, leader := ll.traverse(index)
		if follower == nil {
			ll.head = leader
			leader.prev = nil
			curr.next = nil
		} else if leader == nil {
			ll.tail = follower
			follower.next = nil
			curr.prev = nil
		} else {
			follower.next = leader
			leader.prev = follower
			curr.next = nil
			curr.prev = nil
		}
		ll.length--
	}
}

func main() {
	ll := linkedlist{}
	ll.print()
	ll.initialize(5)
	ll.append(16)
	ll.prepend(10)
	ll.print()
	ll.printBackwards()
}
