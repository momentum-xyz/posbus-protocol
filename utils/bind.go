package utils

import "github.com/google/uuid"

func BinID(x uuid.UUID) []byte {
	binid, _ := x.MarshalBinary()
	return binid
}
