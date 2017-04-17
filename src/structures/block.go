package structures

type block struct {
	elements []*element
}


func Block(els []*element) *block {
	b := &block{}
	b.elements = els
	return b
}
