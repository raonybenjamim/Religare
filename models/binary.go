package models

type Binary int

const (
	Zero Binary = 0
	One  Binary = 1
)

func (b Binary) String() string {
	if b == Zero {
		return "0"
	}
	return "1"
}
