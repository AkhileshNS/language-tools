/*
 Note. T denotes any type

 ** Arrays
 ** New = []T{}
 ** Insert (at start) - = insert(arr, 0, items)
 ** Insert (at end) - = append(arr, n)
 ** Insert (at middle) - = insert(arr, index, items)
 ** Lookup - [n]
 ** Remove (at start) - = arr[1:]
 ** Remove (at end) - = arr[:len(arr)-1]
 ** Remove (at middle) - = delete(arr, index) [OR] = deleteNoOrder(arr, index)
 ** Slice - arr[start:end] # start is inclusive and end is exclusive
*/

package main

import (
	"fmt"
	"time"
)

func insert(arr []int, index int, items ...int) []int {
	if n := len(arr) + len(items); n <= cap(arr) {
		newArr := arr[:n]
		copy(newArr[index+len(items):], arr[index:])
		copy(newArr[index:], items)
		return newArr
	}
	newArr := make([]int, len(arr)+len(items))
	copy(newArr, arr[:index])
	copy(newArr[index:], items)
	copy(newArr[index+len(items):], arr[index:])
	return newArr
}

func delete(arr []int, index int) []int {
	return arr[:index+copy(arr[index:], arr[index+1:])]
}

func deleteNoOrder(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

// ========================================================

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}

func generateArray(n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	return arr
}

func main() {
	size := 10

	arr := generateArray(size)
	arr = insert(arr, 1, -1)
	fmt.Println(arr)
	arr = delete(arr, 1)
	fmt.Println(arr)
	arr = insert(arr, 1, -1)
	arr = deleteNoOrder(arr, 1)
	fmt.Println(arr)
}
