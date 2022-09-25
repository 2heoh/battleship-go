package contracts

type Letter int

const (
	A Letter = iota + 1
	B
	C
	D
	E
	F
	G
	H
)

func (l Letter) String() string {
	return [...]string{"A", "B", "C", "D", "E", "F", "G", "H"}[l-1]
}

func FromString(input string) Letter {
	switch input {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	case "E":
		return E
	case "F":
		return F
	case "G":
		return G
	case "H":
		return H
	default:
		panic("out of range")
	}
}
