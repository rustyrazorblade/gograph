/**
 * Created with IntelliJ IDEA.
 * User: jhaddad
 * Date: 7/12/13
 * Time: 2:53 PM
 * To change this template use File | Settings | File Templates.
 */
package graph

type Vertex struct {
	id string
	out_edges []*Edge
	in_edges []*Edge
}

func (v *Vertex) AddEdge(to_vertex *Vertex, score int) *Edge {
	e := &Edge{in:v, out:to_vertex, score:score}
	v.out_edges = append(v.out_edges, e)
	to_vertex.in_edges = append(to_vertex.out_edges, e)
	return e
}
func (v *Vertex) OutV() []*Vertex {
	total := len(v.out_edges)
	result := make([]*Vertex, total)
	for i := 0; i < total; i++ {
		result[i] = v.out_edges[i].out
	}
	return result
}

type Edge struct {
	in *Vertex
	out *Vertex
	score int
}

type Graph struct {
	ids map[string] *Vertex;
}


func (g *Graph) AddVertex(id string) *Vertex {
	v := &Vertex{id:id}
	return v
}
