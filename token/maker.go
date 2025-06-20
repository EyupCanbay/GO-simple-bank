package token

import (
	"time"
)

// maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the tokens valid or not
	VerifyToken(token string) (*Payload, error)

}
