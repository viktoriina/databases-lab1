package models

type Block struct {
	ID        int64
	Timestamp int64
	Nonce     int64

	LogicalPresent bool
}
