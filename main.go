package main

import "faang-algorithms/list"

func main() {
	var head list.Node
	list.InsertFirst("Thiago", &head)
	list.InsertFirst("Luiz", &head)
	list.InsertFirst("Pereira", &head)
	list.InsertFirst("Nunes", &head)
	list.PrintList(&head)

	list.RemoveFirst(&head)
	list.RemoveFirst(&head)
	list.RemoveFirst(&head)
	list.RemoveFirst(&head)
	list.PrintList(&head)
}
