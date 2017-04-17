package structures

import (
	"fmt"
	"unicode/utf8"
)

type element struct {
	elements []element
	*field
}

func (el element) Equals(el2 element) bool {
	if el.elements == nil && el2.elements != nil {
		return false
	}

	if el.elements != nil && el2.elements == nil {
		return false
	}

	if len(el.elements) != len(el2.elements) {
		return false
	}

	for i := 0; i < len(el.elements); i++ {
		if !el.elements[i].Equals(el2.elements[i]) {
			return false
		}
	}

	if !el.field.Equals(*el2.field) {
		return false
	}

	return true
}

func NewElement(f *field, els []element) element {
	el := element{}
	el.field = f
	el.elements = els
	return el
}

func (el element) ToString() string {
	var s string

	if el.len > 0 {
		s += el.FormatMq()
	}

	if el.elements == nil || len(el.elements) == 0 {
		return s
	}

	for _, e := range el.elements {
		s += e.ToString()
	}

	return s
}

func (el *element) Parse(s string) int32 {
	var pos int32
	var strLen = int32(utf8.RuneCountInString(s))

	if strLen < pos+el.len {
		fmt.Printf("String length %d less than %d", strLen, pos+el.len+1)
		return pos
	}

	if el.len > 0 {
		el.val = s[pos:el.len]
		pos += el.len
	}

	if el.elements == nil || len(el.elements) == 0 {
		return pos
	}

	for _, e := range el.elements {
		pos += e.Parse(s[pos : strLen])
	}

	return pos
}
