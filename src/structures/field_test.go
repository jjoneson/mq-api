package structures

import "testing"

func TestFormat(t *testing.T) {
	currency := NewCurrencyField(0, "123")
	date := NewDateField(0, "123")
	cases := []struct {
		got  string
		want string
	}{
		{currency.Format(), "$123"},
		{date.Format(), "123"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("Format == %q, want %q", c.got, c.want)
		}
	}
}
