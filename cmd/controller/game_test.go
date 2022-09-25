package controller

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"battleship-go/cmd/letter"
	"testing"
)

func TestCheckIsHitTrue(t *testing.T) {
	ships := InitializeShips()
	counter := 1

	for _, ship := range ships {
		letter := letter.Letter(counter)

		for i := 0; i < ship.Size; i++ {
			ship.positions = append(ship.positions, Position{Column: letter, Row: i})
		}

		counter++
	}

	result, _ := CheckIsHit(ships, &Position{letter.A, 1})

	assert.True(t, result)
}

func TestCheckIsHitFalse(t *testing.T) {
	ships := InitializeShips()
	counter := 1

	for _, ship := range ships {
		letter := letter.Letter(counter)

		for i := 0; i < ship.Size; i++ {
			ship.positions = append(ship.positions, Position{Column: letter, Row: i})
		}

		counter++
	}

	result, _ := CheckIsHit(ships, &Position{letter.H, 1})

	assert.False(t, result)
}

func TestCheckIsHitPositstionIsNull(t *testing.T) {
	_, err := CheckIsHit(InitializeShips(), nil)

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("shot is nil"))
}

func TestCheckIsHitShipIsNull(t *testing.T) {
	_, err := CheckIsHit(nil, &Position{letter.H, 1})

	fmt.Println(err)

	assert.Error(t, err)

	assert.Equal(t, err, errors.New("ships is nil"))
}

func TestIsShipValidFalse(t *testing.T) {
	ship := Ship{Name: "TestShip", Size: 3}
	result := IsShipValid(ship)
	assert.False(t, result)
}

func TestIsShipValidTrue(t *testing.T) {
	positions := []Position{{letter.A, 1}, {letter.A, 1}, {letter.A, 1}}
	ship := Ship{Name: "TestShip", Size: 3, positions: positions}

	result := IsShipValid(ship)

	assert.True(t, result)
}
