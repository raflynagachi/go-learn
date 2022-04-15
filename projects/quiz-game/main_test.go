package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var problemTestCase = []problem{
	{
		q: "1+2",
		a: "3",
	},
	{
		q: "11+5",
		a: "16",
	},
	{
		q: "4+3",
		a: "7",
	},
}

func TestCalculation(t *testing.T) {
	var count int
	for _, p := range problemTestCase {
		checkProblem(p, p.a, &count)
	}
	assert.Equal(t, count, 3, "must be equal to 3")
}

func TestParseLines(t *testing.T) {
	var lines = [][]string{
		{
			"1+2", "3",
		},
		{
			"11+5", "16",
		},
		{
			"4+3", "7",
		},
	}
	assert.Equal(t, parseLines(lines), problemTestCase)
}
