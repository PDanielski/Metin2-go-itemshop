package account

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

//Password is a type representing an account password
type Password string

//NewPassword is used to create hashed passwords.
//If the password given in the input is already hashed, set isHashed to true
func NewPassword(psw string, isHashed bool) Password {
	if isHashed {
		return Password(psw)
	}

	h := sha1.New()

	h.Write([]byte(psw))
	psw1 := h.Sum(nil)
	h.Reset()

	h.Write(psw1)
	psw2 := h.Sum(nil)

	psw3 := "*" + strings.ToUpper(hex.EncodeToString(psw2))
	return Password(psw3)
}
