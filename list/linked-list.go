package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func (n *Node) addNode() error {
	return nil
}

func (n *Node) removeNode() error {
	return nil
}

func main() {
	fmt.Println("Teste go mod")
}
