package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"battleship-go/cmd/contracts"
)

func TestParsePosition(t *testing.T) {
	actual := parsePosition("A1")
	expected := &contracts.Position{Column: contracts.A, Row: 1}

	assert.Equal(t, expected, actual)
}

func TestParsePosition2(t *testing.T) {
	//given
	var expected *contracts.Position = &contracts.Position{Column: contracts.B, Row: 1}
	//when
	var actual *contracts.Position = parsePosition("B1")
	//then
	assert.Equal(t, expected, actual)
}
