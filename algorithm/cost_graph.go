package algorithm

import (
	"github.com/gyuho/goraph"
	"k-anon/model"
	"strconv"
)

func BuildCostGraph(t *model.Table) goraph.Graph {
	g := BuildCoreGraph(t)
	addCosts(g, t)
	return g
}

func addCosts(g goraph.Graph, t *model.Table) {
	for id1 := range g.GetNodes() {
		for id2 := range g.GetNodes() {
			if id1 != id2 {
				i := getIndex(id1)
				j := getIndex(id2)
				v1 := t.Rows[i]
				v2 := t.Rows[j]
				cost := CalculateCost(v1, v2)
				g.AddEdge(id1, id2, cost)
			}
		}
	}
}

func getIndex(id goraph.ID) int {
	i, err := strconv.Atoi(id.String())
	if err != nil { // this shouldn't happen, because we added each node with Itoa()
		panic("node ID cannot be converted back")
	}
	return i
}
