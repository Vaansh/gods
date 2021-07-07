package graph

import "fmt"

type Vertex struct {
	adjacent []*Vertex
	key      int
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if HasVertex(g.vertices, k) {
		fmt.Println("Vertex %v not added since it is an existing key", k)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) AddEdge(a, b int) {
	begin := g.GetVertex(a)
	end := g.GetVertex(b)
	if begin == nil || end == nil {
		fmt.Printf("Invalid edge (%v-->%v)\n", a, b)
		return
	}
	if HasVertex(begin.adjacent, b) {
		fmt.Printf("Existing edge (%v-->%v)\n", a, b)
		return
	}
	begin.adjacent = append(begin.adjacent, end)
}

func (g *Graph) GetVertex(k int) *Vertex {
	for a, b := range g.vertices {
		if b.key == k {
			return g.vertices[a]
		}
	}
	return nil
}

func HasVertex(v []*Vertex, k int) bool {
	for _, i := range v {
		if k == i.key {
			return true
		}
	}
	return false
}

func (g *Graph) Display() {
	for _, i := range g.vertices {
		fmt.Printf("\nVertex %v : ", i.key)
		for _, i := range i.adjacent {
			fmt.Printf("%v", i.key)
		}
	}
	fmt.Println()
}
