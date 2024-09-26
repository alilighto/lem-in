package utils

type Vertex struct {
	Name string
}

type Graph struct {
	Vertices map[string]*Vertex
	Edges    map[string][]*Vertex
	Start    *Vertex
	End      *Vertex
}

func GraphInit() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
		Edges:    make(map[string][]*Vertex),
	}
}

func (g *Graph) AppendVertex(name string) {
	if _, exists := g.Vertices[name]; !exists {
		g.Vertices[name] = &Vertex{Name: name}
	}
}

func (g *Graph) AddEdge(from, to *Vertex) {
	g.Edges[from.Name] = append(g.Edges[from.Name], to)
}

func (g *Graph) SetStart(name string) {
	g.Start = g.Vertices[name]
}

func (g *Graph) SetEnd(name string) {
	g.End = g.Vertices[name]
}

func (g *Graph) GetVertex(name string) *Vertex {
	return g.Vertices[name]
}
