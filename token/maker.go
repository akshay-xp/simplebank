package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is volid or not
	VerifyToken(token string) (*Payload, error)
}
