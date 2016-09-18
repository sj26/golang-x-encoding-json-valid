package json_test

import (
	"testing"
	"."
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		in []byte
		want bool
	}{
		{[]byte(`{"foo":"bar"}`), true},
		{[]byte(`"foo"`), true},
		{[]byte(`1`), true},
		{[]byte(`foo`), false},
		{[]byte(``), false},
	}

	for _, c := range cases {
		if json.IsValid(c.in) != c.want {
			t.Errorf("Expected IsValid(%v) to be %v but was %v", c.in, c.want, !c.want)
		}
	}
}
