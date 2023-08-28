package uuid

import (
	"crypto/rand"
	"fmt"
)

func NewV4() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		randBytes[0:4], randBytes[4:6], randBytes[6:8],
		randBytes[8:10], randBytes[10:16])

	return uuid
}
