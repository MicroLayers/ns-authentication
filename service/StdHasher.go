package service

import (
	"crypto/md5"
	"fmt"
	"ns-auth/storage"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// StdHasher standard hasher
type StdHasher struct {
	salt  HasherSalt
	mutex sync.RWMutex
}

// HasherSalt string type used to discriminate in wire
type HasherSalt string

// NewStdHasher StdHasher's instantiator, used by wire
func NewStdHasher(salt HasherSalt) storage.Hasher {
	return &StdHasher{salt: salt, mutex: sync.RWMutex{}}
}

func buildSecret(username string, domain string, password string, salt HasherSalt) [16]byte {
	data := []byte(fmt.Sprintf("%s$%s~%s?%s", username, domain, password, string(salt)))
	hash := md5.Sum(data)

	return hash
}

// HashPassword hash the password to store it in the database
func (h *StdHasher) HashPassword(username string, domain string, password string) string {
	h.mutex.Lock()
	salt := buildSecret(username, domain, password, h.salt)
	h.mutex.Unlock()
	bytes, _ := bcrypt.GenerateFromPassword(salt[:], 14)

	return string(bytes)
}

// CheckPassword verify the credentials against the given stored password
func (h *StdHasher) CheckPassword(
	username string,
	domain string,
	password string,
	storedPassword string,
) bool {
	h.mutex.Lock()
	salt := buildSecret(username, domain, password, h.salt)
	h.mutex.Unlock()

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), salt[:])

	return err == nil
}
