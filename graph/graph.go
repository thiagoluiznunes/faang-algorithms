package graph

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Value   string
	Edges   []*Vertex
	Visited bool
}

type Graph struct {
	Vertexs []*Vertex
}

func (g *Graph) InsertVertex(vertex *Vertex) (err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to insert vertex")
		}
	}()

	g.Vertexs = append(g.Vertexs, vertex)

	return nil
}

func (g *Graph) PrintGraph() (err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to print graph")
		}
	}()

	for _, val := range g.Vertexs {
		var vertxsValues string
		for _, vertx := range val.Edges {
			vertxsValues = vertxsValues + fmt.Sprintf("%s, ", vertx.Value)
		}
		fmt.Println("Profile: ", val.Value)
		fmt.Println("Friends: ", vertxsValues)
	}

	return nil
}
