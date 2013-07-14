/**
 * Created with IntelliJ IDEA.
 * User: jhaddad
 * Date: 7/13/13
 * Time: 3:58 PM
 * To change this template use File | Settings | File Templates.
 */
package graph

import "testing"

func TestCreateGraph(t *testing.T) {
	t.SkipNow()
	g := new(Graph)
	tmp := make(map[string]string)
	tmp["blah"] = "whatever"

	g.addVertex("something", tmp)
}

