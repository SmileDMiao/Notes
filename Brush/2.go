package main

import (
	"fmt"
	"math"
	"strconv"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := 0
	num2 := 0
	result1 := 0
	result2 := 0
	for l1 != nil {
		result1 = result1 + l1.Val*int((math.Pow(10, float64(num1))))
		fmt.Println(result1)

		num1 = num1 + 1
		l1 = l1.Next
	}

	for l2 != nil {
		result2 = result2 + l2.Val*int((math.Pow(10, float64(num2))))

		num2 = num2 + 1
		l2 = l2.Next
	}

	result := result1 + result2

	fmt.Println(result1)
	fmt.Println(result2)

	str := strconv.Itoa(result)

	first, _ := strconv.Atoi(str[0:1])

	l3 := createNode(first)

	if len(str) == 0 {
		return l3
	}

	ttt := len(str)

	fmt.Println(str[1:ttt])
	for _, s := range str[1:ttt] {
		value, _ := strconv.Atoi(string(s))
		tmpNode := createNode(int(value))
		tmpNode.Next = l3
		l3 = tmpNode
	}

	return l3
}

func createNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func main() {
	// node1 := createNode(2)
	// node1.Next = createNode(4)
	// node1.Next.Next = createNode(3)

	// node2 := createNode(5)
	// node2.Next = createNode(6)
	// node2.Next.Next = createNode(4)

	var head ListNode
	var tail ListNode
	a := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	b := []int{5, 6, 4}
	for _, num := range a {
		node := createNode(num)
		if head.Next == nil {
			head.Next = node
		}
	}

	for _, num := range b {
		node := ListNode{Val: num, Next: nil}
		if tail.Next == nil {
			tail.Next = &node
		}
	}

	addTwoNumbers(head.Next, tail.Next)
}
