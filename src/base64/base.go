package base64

import (
	"encoding/base64"
)

func Encode(word string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(word))
	return encoded
}
func Decode(word string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(word)
	return decoded, err
}
