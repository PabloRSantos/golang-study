package infra

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct{}

func (ba *BcryptAdapter) Hash(value string) (string, error) {
	valueBytes := []byte(value)

	hash, err := bcrypt.GenerateFromPassword(valueBytes, bcrypt.DefaultCost)

	return string(hash), err
}

func (ba *BcryptAdapter) Compare(value string, hash string) bool {
	valueBytes := []byte(value)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, valueBytes)
	match := err == nil

	return match
}
