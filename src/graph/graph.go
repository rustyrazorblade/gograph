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
	in_edges []Edge
	properties map[string] string
}

func (*Vertex) addEdge(inV *Vertex, outV *Vertex) {

}
func (*Vertex) outV() []Vertex {
	tmp := make([]Vertex, 1)
	return tmp
}

type Edge struct {
	in *Vertex
	out *Vertex
	properties map[string] string
}

type Graph struct {
	ids map[string] *Vertex;
}

func (g *Graph) addVertex(id string, properties map[string]string) {
	v := new(Vertex)
	v.properties = properties
	g.ids[id] = v
}
