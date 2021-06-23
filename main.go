package main

import (
	challenges "faang-algorithms/coding-challenges"
	"fmt"
)

func main() {

	// List execution test
	// var head list.Node
	// list.InsertFirst("Thiago", &head)
	// list.PrintList(&head)
	// list.RemoveFirst(&head)
	// list.PrintList(&head)

	// var myGraph graph.Graph
	// vertex1 := graph.Vertex{
	// 	Value: "1",
	// 	Edges: nil,
	// }

	// vertex2 := graph.Vertex{
	// 	Value: "2",
	// 	Edges: nil,
	// }
	// vertex2.Edges = append(vertex2.Edges, &vertex1)

	// vertex3 := graph.Vertex{
	// 	Value: "3",
	// 	Edges: nil,
	// }
	// vertex3.Edges = append(vertex3.Edges, &vertex1)
	// vertex3.Edges = append(vertex3.Edges, &vertex2)

	// vertex4 := graph.Vertex{
	// 	Value: "4",
	// 	Edges: nil,
	// }
	// vertex4.Edges = append(vertex4.Edges, &vertex1)

	// vertex5 := graph.Vertex{
	// 	Value: "5",
	// 	Edges: nil,
	// }
	// vertex5.Edges = append(vertex5.Edges, &vertex1)
	// vertex5.Edges = append(vertex5.Edges, &vertex2)
	// vertex5.Edges = append(vertex5.Edges, &vertex3)

	// myGraph.InsertVertex(&vertex1)
	// myGraph.InsertVertex(&vertex2)
	// myGraph.InsertVertex(&vertex3)
	// myGraph.InsertVertex(&vertex4)
	// myGraph.InsertVertex(&vertex5)
	// myGraph.PrintGraph()

	// myGraph.BreadthFirstSearch(&vertex1)

	arr := []int{6, 2, 1, 9, 0}
	swaps, arr := challenges.SelectionSort(arr)
	fmt.Println("Array sorted: ", arr)
	fmt.Println("Swaps: ", swaps)
}
