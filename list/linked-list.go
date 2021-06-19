package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func insertFirst(value string, head *Node) {

	if head.Value == "" && head.Next == nil {
		head.Value = value
	} else {
		next := *head
		*head = Node{
			Value: value,
			Next:  &next,
		}
	}
}

func insertLast(value string, head *Node) {

	if head.Value == "" && head.Next == nil {
		head.Value = value
	} else {
		node := head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = &Node{
			Value: value,
			Next:  nil,
		}
	}
}

func printList(head *Node) error {

	if head == nil {
		return nil
	}
	currentNode := head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}

	return nil
}

func removeFirst(head *Node) error {

	if head == nil {
		return nil
	} else if head.Next != nil {
		*head = *head.Next
	} else {
		*head = Node{}
	}

	return nil
}

func main() {
	var head Node
	insertLast("Thiago", &head)
	insertLast("Luiz", &head)
	insertLast("Pereira", &head)
	insertLast("Nunes", &head)
	printList(&head)

	removeFirst(&head)
	removeFirst(&head)
	removeFirst(&head)
	removeFirst(&head)
	printList(&head)
}
