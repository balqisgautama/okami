package util

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Base64decoder(content string) ([]byte, error) {
	output, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		output, err = base64.RawStdEncoding.DecodeString(content)
		if err != nil {
			inside := content
			inside = inside + "=="
			output, err = base64.StdEncoding.DecodeString(inside)
			if err != nil {
				return nil, errors.New("HASHING_DATA_INVALID")
			}
		}
	}
	return output, err
}

func Base64encoder(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

//created at 06-20-2022
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//created at 06-20-2022
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// created at 08-31-2022
func SHA256(input string) (result string) {
	sum := sha256.Sum256([]byte(input))
	result = hex.EncodeToString(sum[:])
	return
}