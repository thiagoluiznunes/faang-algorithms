package main

import (
	"faang-algorithms/graph"
)

func main() {

	// List execution test
	// var head list.Node
	// list.InsertFirst("Thiago", &head)
	// list.PrintList(&head)
	// list.RemoveFirst(&head)
	// list.PrintList(&head)

	var myGraph graph.Graph
	firstVertex := graph.Vertex{
		Value: "Thiago Luiz",
		Edges: nil,
	}
	myGraph.InsertVertex(&firstVertex)

	secondVertex := graph.Vertex{
		Value: "Debora Santos",
		Edges: nil,
	}
	secondVertex.Edges = append(secondVertex.Edges, &firstVertex)
	myGraph.InsertVertex(&secondVertex)

	thirdVertex := graph.Vertex{
		Value: "Marta Santos",
		Edges: nil,
	}
	thirdVertex.Edges = append(thirdVertex.Edges, &firstVertex)
	thirdVertex.Edges = append(thirdVertex.Edges, &secondVertex)
	myGraph.InsertVertex(&thirdVertex)

	myGraph.PrintGraph()
}
