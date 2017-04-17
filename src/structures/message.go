package structures

type message struct {
	*element
}

func NewMessage(root *element) *message {
	m := &message{}
 	m.element = root
	return m
}



