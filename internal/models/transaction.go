package models

type Transaction struct {
	ID      int64
	BlockID int64
	To      [20]byte
	Value   int64

	LogicalPresent bool
}
