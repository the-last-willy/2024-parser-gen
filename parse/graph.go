package parse

import "parsium/graph"

type Graph struct {
	Topology *graph.Topology

	Start       graph.V
	Productions map[graph.V]string
	Transitions map[graph.V]map[string]graph.V
}

func NewGraph() *Graph {
	g := &Graph{
		Topology:    graph.NewTopology(),
		Productions: make(map[graph.V]string),
		Transitions: make(map[graph.V]map[string]graph.V),
	}
	g.Topology, g.Start = g.Topology.AddV()
	return g
}

func (g *Graph) AddProduction(v graph.V, typ string) {
	g.Productions[v] = typ
}

func (g *Graph) AddTransition(from graph.V, transition string) (to graph.V) {
	if fromTrs, ok := g.Transitions[from]; ok {
		v2, has := fromTrs[transition]
		if !has {
			g.Topology, v2 = g.Topology.AddV()
			g.Transitions[from][transition] = v2
		}
		return v2
	}
	var v2 graph.V
	g.Topology, v2 = g.Topology.AddV()
	g.Transitions[from] = map[string]graph.V{transition: v2}
	return v2
}
