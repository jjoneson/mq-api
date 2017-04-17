package structures

type message struct {
	*element
}

func Message(root *element) *message {
	m := &message{}
 	m.element = root
	return m
}



