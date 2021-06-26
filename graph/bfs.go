package graph

import "fmt"

func (g *Graph) BreadthFirstSearch(startVertex *Vertex) {

	for _, value := range startVertex.Edges {
		fmt.Println(value.Value)
		value.Visited = true
	}

}
