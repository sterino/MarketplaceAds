package email

import (
	"crypto/rand"
	"fmt"
)

func GenerateCode() string {
	var code [6]byte
	rand.Read(code[:])
	return fmt.Sprintf("%06d", code[0:6])
}
