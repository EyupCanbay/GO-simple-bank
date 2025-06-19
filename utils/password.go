package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashed password retrun the bcrypt hash of the password
func HasPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}

	return string(hashedPassword), nil
}

// CheckPassword checks if the provided passsord is correct or not
func CheckPassword(password string, hasshedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasshedPassword), []byte((password)))
}
