package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	testList := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name:  "1",
			input: "([])",
			exp:   true,
		},
		{
			name:  "2",
			input: "{[(]}",
			exp:   false,
		},
		{
			name:  "3",
			input: "{()[]{[]}}",
			exp:   true,
		},
		{
			name:  "4",
			input: "4*(x+y)",
			exp:   true,
		},
	}
	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			res := Valid(test.input)

			assert.Equal(t, test.exp, res)
		})
	}
}
