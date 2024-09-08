package graph

import "slices"

type Topology struct {
	edges  []e
	vcount int
}

func NewTopology() *Topology {
	return &Topology{}
}

func (t *Topology) containsV(vert V) bool {
	idx := int(vert.Impl.(v))
	return idx >= 0 && idx < t.vcount
}

// Builder interface

func (t *Topology) AddE(v1, v2 V) (*Topology, E) {
	t2 := *t
	if !t.containsV(v1) {
		panic("Topology.AddE: does not contain v1")
	}
	if !t.containsV(v2) {
		panic("Topology.AddE: does not contain v2")
	}
	e := NewE(v1, v2).(e)
	if !slices.Contains(t2.edges, e) {
		t2.edges = append(t2.edges, e)
	}
	return &t2, e
}

func (t *Topology) AddV() (*Topology, V) {
	t2 := *t
	v := V{Impl: v(t2.vcount)}
	t2.vcount += 1
	return &t2, v
}

// Topology interface

func (t *Topology) Es() <-chan E {
	c := make(chan E)
	go func() {
		for _, e := range t.edges {
			c <- e
		}
		close(c)
	}()
	return c
}

func (t *Topology) Vs() <-chan V {
	c := make(chan V)
	go func() {
		for i := 0; i < t.vcount; i++ {
			c <- V{Impl: v(i)}
		}
		close(c)
	}()
	return c
}
