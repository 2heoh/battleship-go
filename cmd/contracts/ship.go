package contracts

import (
	"strconv"
	"strings"
)

type Ship struct {
	isPlaced  bool
	Name      string
	Size      int
	Positions []Position
	Color     Color
}

func NewPosition(letter Letter, number int) *Position {
	return &Position{Column: letter, Row: number}
}

func (s *Ship) AddPosition(input string) {
	if s.Positions == nil {
		s.Positions = []Position{}
	}

	letter := FromString(string(strings.ToUpper(input)[0]))
	number, err := strconv.Atoi(string(input[1]))

	if err != nil {
		panic(err)
	}

	s.Positions = append(s.Positions, Position{Column: letter, Row: number})
}

func (s *Ship) GetPositions() []Position {
	return s.Positions
}

func (s *Ship) SetPositions(p *Position) {
	s.Positions = append(s.Positions, *p)
}
