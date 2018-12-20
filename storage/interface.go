package storage

const AUTH_TYPE_USERNAME_PASSWORD = "username_password"

// Storage authentication storage struct
type Storage struct {
	UsernamePassword UsernamePasswordStorage
	Token            TokenStorage
}

// UsernamePasswordStorage username/password storage interface
type UsernamePasswordStorage interface {
	FindUser(username string, hashedPassword string, domain string) (User, error)
}

// TokenStorage token storage interface
type TokenStorage interface {
	FindOrCreate(user User, authType string) (AuthToken, error)
}

// Hasher dedicated hashing interface
type Hasher interface {
	HashPassword(username string, password string, domain string) string
}

// User definition of a user
type User struct {
	Id       string
	Username string
	Domain   string
}

type AuthToken struct {
	Token          string
	RefreshToken   string
	ExpirationDate uint32
}
