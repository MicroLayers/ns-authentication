package service

import (
	"crypto/md5"
	"fmt"
	"ns-auth/storage"

	"golang.org/x/crypto/bcrypt"
)

type StdHasher struct {
	salt HasherSalt
}

type HasherSalt string

func NewStdHasher(salt HasherSalt) storage.Hasher {
	return &StdHasher{salt: salt}
}

func buildSecret(username string, domain string, password string, salt HasherSalt) [16]byte {
	data := []byte(fmt.Sprintf("%s$%s~%s?%s", username, domain, password, string(salt)))
	hash := md5.Sum(data)

	return hash
}

func (h *StdHasher) HashPassword(username string, domain string, password string) string {
	salt := buildSecret(username, domain, password, h.salt)
	bytes, _ := bcrypt.GenerateFromPassword(salt[:], 14)

	return string(bytes)
}

func (h *StdHasher) CheckPassword(
	username string,
	domain string,
	password string,
	storedPassword string,
) bool {
	salt := buildSecret(username, domain, password, h.salt)

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), salt[:])

	return err == nil
}
