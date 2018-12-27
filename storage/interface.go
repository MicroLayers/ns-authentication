package storage

// AuthTypeUsernamePassword authentication type username/password
const AuthTypeUsernamePassword = "username_password"

// Storage authentication storage struct
type Storage struct {
	UsernamePassword UsernamePasswordStorage
	Token            TokenStorage
}

// UsernamePasswordStorage username/password storage interface
type UsernamePasswordStorage interface {
	AddUser(username string, domain string, password string) (*User, error)
	FindUser(username string, domain string, password string) (*User, error)
}

// TokenStorage token storage interface
type TokenStorage interface {
	FindOrCreateTokenFromUser(user *User, authType string) (*AuthToken, error)
	FindUserFromToken(token string) (*User, error)
}

// Hasher dedicated hashing interface
type Hasher interface {
	HashPassword(username string, domain string, password string) string
	CheckPassword(username string, domain string, password string, storedPassword string) bool
}

// User definition of a user
type User struct {
	ID       string
	Username string
	Domain   string
}

// AuthToken an authorization token
type AuthToken struct {
	Token        string
	RefreshToken string
	ExpiryDate   int64
}
