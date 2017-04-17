package structures

import (
	"fmt"
	"strconv"
)

type field struct {
	len       int32
	val       string
	formatter string
	padder    string
}

func (f field) Equals(f2 field) bool {
	if f.len == f2.len &&
		f.val == f2.val &&
		f.formatter == f2.formatter &&
		f.padder == f2.padder {
		return true
	}
	return false
}

type currencyField struct {
	*field
}

type dateField struct {
	*field
}

func (f field) Format() string {
	return fmt.Sprintf(f.formatter, f.val)
}

func (f field) Pad() string {
	return fmt.Sprintf(f.padder, f.val)
}

func (f field) FormatMq() string {
	return f.Pad()[:f.len]
}

func newField(len int32, val string) *field {
	f := &field{}
	f.len = len
	f.val = val
	return f
}

func CurrencyField(len int32, val string) *currencyField {
	f := &currencyField{newField(len, val)}
	f.formatter = "$%s"
	f.padder = "%0" + strconv.FormatInt(int64(f.len), 10) + "s"
	return f
}

func DateField(len int32, pos int32, val string) *currencyField {
	f := &currencyField{newField(len, val)}
	f.formatter = "%s"
	return f
}
