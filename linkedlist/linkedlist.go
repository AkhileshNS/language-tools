/*
 Note. All functions are in-order

 ** Singly Linked List
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
}

type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (ll *linkedlist) initialize(value int) {
	n := node{value: value, next: nil}
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

	n := node{value: value, next: nil}
	ll.tail.next = &n
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

func (ll *linkedlist) traverse(index int) (*node, *node, *node) {
	if ll.isEmpty() {
		return nil, nil, nil
	}

	var follower *node
	curr := ll.head

	for i := 0; i < index && i < ll.length-1; i++ {
		follower = curr
		curr = curr.next
	}

	return follower, curr, curr.next
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
	n.next = leader
	ll.length++
}

func (ll *linkedlist) remove(index int) {
	if !ll.isEmpty() {
		follower, curr, leader := ll.traverse(index)
		if follower == nil {
			ll.head = leader
			curr.next = nil
		} else if leader == nil {
			ll.tail = follower
			follower.next = nil
		} else {
			follower.next = leader
			curr.next = nil
		}
		ll.length--
	}
}

func main() {
	ll := linkedlist{}
	ll.print()
	ll.initialize(5)
	ll.print()
	ll.append(16)
	ll.print()
	ll.prepend(10)
	ll.print()
	ll.remove(-1)
	ll.print()
}
