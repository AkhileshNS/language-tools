/*
 ** Queue
 ** New - newQueue(value)
 ** Enqueue - .enqueue(value)
 ** Dequeue - .dequeue()
 ** Peek - .peek()
 */

// =============================================
package main

import "fmt"

type node struct {
	value string
	next  *node
}

type queue struct {
	front  *node
	back   *node
	length int
}

func newQueue(value string) queue {
	n := node{value: value, next: nil}
	return queue{front: &n, back: &n, length: 1}
}

func (q *queue) isEmpty() bool {
	if q.front == nil {
		fmt.Println("Queue is empty!")
		return true
	}
	return false
}

func (q *queue) enqueue(value string) {
	if q.isEmpty() {
		*q = newQueue(value)
	}

	n := node{value: value, next: nil}
	q.back.next = &n
	q.back = &n
	q.length++
}

func (q *queue) dequeue() *node {
	if q.isEmpty() {
		return nil
	}
	n := q.front
	q.front = q.front.next
	n.next = nil
	q.length--
	return n
}

func (q *queue) peek() string {
	if !q.isEmpty() {
		return q.front.value
	}
	return ""
}

func (q *queue) print() {
	if !q.isEmpty() {
		curr := q.front
		str := q.front.value
		curr = curr.next
		for curr != nil {
			str += " <-- " + curr.value
			curr = curr.next
		}
		fmt.Println(str)
	}
}

func main() {
	q := newQueue("Akhilesh")
	q.print()
	q.enqueue("Anirudh")
	q.enqueue("Anirban")
	q.print()
	fmt.Println(q.peek())
	fmt.Println(q.dequeue())
	q.print()
}
