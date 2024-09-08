package graph_test

import (
	"parsium/assert"
	"parsium/graph"
	"testing"
)

func TestDefaultTopology(t *testing.T) {
	t.Run("Es closes immediately", func(t *testing.T) {
		to := graph.NewTopology()
		_, ok := <-to.Es()
		assert.False(t, ok)
	})
	t.Run("Vs closes immediately", func(t *testing.T) {
		to := graph.NewTopology()
		_, ok := <-to.Vs()
		assert.False(t, ok)
	})
}

func TestT(t *testing.T) {
	t.Run("Vs receives one for each AddV", func(t *testing.T) {
		to := &graph.Topology{}

		expected := 3

		// Create expected vertices
		for i := 0; i < expected; i++ {
			to, _ = to.AddV()
		}

		count := 0
		for range to.Vs() {
			count += 1
		}

		assert.Equal(t, count, expected)
	})
	t.Run("Es receives one for each AddE", func(t *testing.T) {
		to := &graph.Topology{}

		to, a := to.AddV()
		to, b := to.AddV()
		to, c := to.AddV()

		to, _ = to.AddE(a, b)
		to, _ = to.AddE(b, c)
		to, _ = to.AddE(c, a)

		count := 0
		for range to.Vs() {
			count += 1
		}

		assert.Equal(t, count, 3)
	})
}

func TestTopology_AddE(t *testing.T) {
	t.Run("returns E containing v1 and v2", func(t *testing.T) {
		to := graph.NewTopology()
		to, v1 := to.AddV()
		to, v2 := to.AddV()
		to, e := to.AddE(v1, v2)
		assert.Equal(t, e.V1(), v1)
		assert.Equal(t, e.V2(), v2)
	})
	t.Run("panic if v1 is not in topology", func(t *testing.T) {
		to := graph.NewTopology()
		to, v1 := to.AddV()
		_, v2 := to.AddV()
		assert.Panic(t, func() { to.AddE(v1, v2) })
	})
	t.Run("panic if v2 is not in topology", func(t *testing.T) {
		to := graph.NewTopology()
		to, v1 := to.AddV()
		_, v2 := to.AddV()
		assert.Panic(t, func() { to.AddE(v1, v2) })
	})
}
