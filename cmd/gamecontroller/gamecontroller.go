package gamecontroller

import (
	"errors"
	"battleship-go/cmd/contracts"
	"math/rand"
)

func CheckIsHit(ships []*contracts.Ship, shot *contracts.Position) (bool, error) {

	if ships == nil {
		return false, errors.New("ships is nil")
	}

	if shot == nil {
		return false, errors.New("shot is nil")
	}

	for _, ship := range ships {
		for _, position := range ship.Positions {
			if shot.Row == position.Row && shot.Column == position.Column {
				return true, nil
			}
		}
	}

	return false, nil
}

func InitializeShips() []*contracts.Ship {
	return []*contracts.Ship{
		NewShip("Aircraft Carrier", 5, contracts.CADET_BLUE),
		NewShip("Battleship", 4, contracts.RED),
		NewShip("Submarine", 3, contracts.CHARTREUSE),
		NewShip("Destroyer", 3, contracts.YELLOW),
		NewShip("Patrol Boat", 2, contracts.ORANGE),
	}
}

func NewShip(name string, size int, color contracts.Color) *contracts.Ship {
	return &contracts.Ship{
		Name:  name,
		Size:  size,
		Color: color,
	}
}

func IsShipValid(ship contracts.Ship) bool {
	return len(ship.Positions) == ship.Size
}

func GetRandomPosition(size int) *contracts.Position {
	letter := contracts.Letter(rand.Intn(size-1) + 1)
	number := rand.Intn(size-1) + 1
	return &contracts.Position{Column: letter, Row: number}
}
