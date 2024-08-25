package tree

//func TestSimpleBuilder(t *testing.T) {
//	t.Run(".Build for default", func(t *testing.T) {
//		b := NewSimpleBuilder[int]()
//		assert.Nil(t, b.Build())
//	})
//	t.Run(".Build for node", func(t *testing.T) {
//		b := NewSimpleBuilder[int]()
//		b = b.New(1, []Builder[int]{}).(SimpleBuilder[int])
//		tr := b.Build()
//		assert.Equal(t, tr.DataOf(tr.Root()), 1)
//		assert.Equal(t, len(tr.ChildrenOf(tr.Root())), 0)
//	})
//}
