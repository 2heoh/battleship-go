package gamecontroller

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"battleship-go/cmd/contracts"
	"testing"
)

func TestCheckIsHitTrue(t *testing.T) {
	ships := InitializeShips()
	counter := 1

	for _, ship := range ships {
		letter := contracts.Letter(counter)

		for i := 0; i < ship.Size; i++ {
			ship.Positions = append(ship.Positions, contracts.Position{Column: letter, Row: i})
		}

		counter++
	}

	result, _ := CheckIsHit(ships, &contracts.Position{Column: contracts.A, Row: 1})

	assert.True(t, result)
}

func TestCheckIsHitFalse(t *testing.T) {
	ships := InitializeShips()
	counter := 1

	for _, ship := range ships {
		letter := contracts.Letter(counter)

		for i := 0; i < ship.Size; i++ {
			ship.Positions = append(ship.Positions, contracts.Position{Column: letter, Row: i})
		}

		counter++
	}

	result, _ := CheckIsHit(ships, &contracts.Position{Column: contracts.H, Row: 1})

	assert.False(t, result)
}

func TestCheckIsHitPositstionIsNull(t *testing.T) {
	_, err := CheckIsHit(InitializeShips(), nil)

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("shot is nil"))
}

func TestCheckIsHitShipIsNull(t *testing.T) {
	_, err := CheckIsHit(nil, &contracts.Position{Column: contracts.H, Row: 1})

	fmt.Println(err)

	assert.Error(t, err)

	assert.Equal(t, err, errors.New("ships is nil"))
}

func TestIsShipValidFalse(t *testing.T) {
	ship := contracts.Ship{Name: "TestShip", Size: 3}
	result := IsShipValid(ship)
	assert.False(t, result)
}

func TestIsShipValidTrue(t *testing.T) {
	positions := []contracts.Position{{Column: contracts.A, Row: 1}, {Column: contracts.A, Row: 2}, {Column: contracts.A, Row: 3}}
	ship := contracts.Ship{Name: "TestShip", Size: 3, Positions: positions}

	result := IsShipValid(ship)

	assert.True(t, result)
}
