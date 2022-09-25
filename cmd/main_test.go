package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"battleship-go/cmd/controller"
	"battleship-go/cmd/letter"
)

func TestParsePosition(t *testing.T) {
	actual := parsePosition("A1")
	expected := &controller.Position{Column: letter.A, Row: 1}

	assert.Equal(t, expected, actual)
}

func TestParsePosition2(t *testing.T) {
	//given
	var expected *controller.Position = controller.NewPosition(letter.B, 1)
	//when
	var actual *controller.Position = parsePosition("B1")
	//then
	assert.Equal(t, expected, actual)
}
