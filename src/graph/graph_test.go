/**
 * Created with IntelliJ IDEA.
 * User: jhaddad
 * Date: 7/13/13
 * Time: 3:58 PM
 * To change this template use File | Settings | File Templates.
 */
package graph

import "testing"

func TestSimpleGraph(t *testing.T) {
	g := new(Graph)

	v1 := g.AddVertex("someid")
	v2 := g.AddVertex("someid2")

	v1.AddEdge(v2, 10)

	vertices := v1.OutV()

	if(len(vertices) != 1) {
		t.Fail()
	}

	if(vertices[0].id != "someid2") {
		t.Fail()
	}

	vertices2 := v2.InV()
	if(vertices2[0].id != "someid") {
		t.Fail()
	}

	edges := v1.OutE()
	if(edges[0].score != 10) {
		t.Fail()
	}

	edges2 := v2.InE()
	if(edges2[0].score != 10) {
		t.Fail()
	}


}

