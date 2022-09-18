package controller

import (
	"errors"
	"go-battleship/cmd/console"
	"go-battleship/cmd/letter"
	"math/rand"
)

func CheckIsHit(ships []*Ship, shot *Position) (bool, error) {

	if ships == nil {
		return false, errors.New("ships is nil")
	}

	if shot == nil {
		return false, errors.New("shot is nil")
	}

	for _, ship := range ships {
		for _, position := range ship.positions {
			if shot.Row == position.Row && shot.Column == position.Column {
				return true, nil
			}
		}
	}

	return false, nil
}

func InitializeShips() []*Ship {
	return []*Ship{
		NewShip("Aircraft Carrier", 5, console.CADET_BLUE),
		NewShip("Battleship", 4, console.RED),
		NewShip("Submarine", 3, console.CHARTREUSE),
		NewShip("Destroyer", 3, console.YELLOW),
		NewShip("Patrol Boat", 2, console.ORANGE),
	}
}

func NewShip(name string, size int, color console.Color) *Ship {
	return &Ship{
		Name:  name,
		Size:  size,
		Color: color,
	}
}

func IsShipValid(ship Ship) bool {
	return len(ship.positions) == ship.Size
}

func GetRandomPosition(size int) *Position {
	letter := letter.Letter(rand.Intn(size-1) + 1)
	number := rand.Intn(size-1) + 1
	return &Position{Column: letter, Row: number}
}
