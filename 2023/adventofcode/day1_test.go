package adventofcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDay1(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "eightwothree", want: "83"},
		{input: "abcone2threexyz", want: "13"},
		{input: "xtwone3four", want: "24"},
		{input: "4nineeightseven2", want: "42"},
		{input: "zoneight234", want: "14"},
		{input: "7pqrstsixteen", want: "76"},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			first, last := findFirstAndLastNumber(tc.input)
			assert.Equal(t, tc.want, first+last)
		})
	}
}
