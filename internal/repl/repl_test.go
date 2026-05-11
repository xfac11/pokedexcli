package repl

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   foo   bar   baz   ",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    " oneword ",
			expected: []string{"oneword"},
		},
		{
			input:    "   Iol,d  HELLO  ",
			expected: []string{"iol,d", "hello"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			actual := CleanInput(c.input)
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
