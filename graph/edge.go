package graph

type E interface {
	V1() V
	V2() V
}

type e struct {
	v1 v
	v2 v
}

func NewE(v1, v2 V) E {
	return e{v1: v1.Impl.(v), v2: v2.Impl.(v)}
}

func (e e) V1() V {
	return V{e.v1}
}

func (e e) V2() V {
	return V{e.v2}
}
