package main

import (
	"fmt"
)

type Node struct {
	Value string
	Next  *Node
}

func addNode(value string, node *Node) error {

	if node == nil {
		return nil
	} else if node.Next != nil {
		addNode(value, node.Next)
	} else {
		node.Next = &Node{
			Value: value,
			Next:  nil,
		}
	}

	return nil
}

func printList(node *Node) error {

	if node == nil {
		return nil
	} else if node.Next != nil {
		printList(node.Next)
	}
	fmt.Println(node.Value)

	return nil
}

func (n *Node) removeNode() error {
	return nil
}

func main() {

	rootNode := Node{}
	addNode("Thiago", &rootNode)
	addNode("Luiz", &rootNode)
	addNode("Pereira", &rootNode)
	addNode("Nunes", &rootNode)
	printList(&rootNode)
}
