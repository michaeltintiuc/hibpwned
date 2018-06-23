package pwd

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// NewPlain creates a new Hash instance
// from a plain-text password
func NewPlain(pwd string) (*Hash, error) {
	if err := ValidatePlain(pwd); err != nil {
		return nil, err
	}

	p := &Hash{BaseURL, SHA1(pwd), false, 0}

	return p, nil
}

// ValidatePlain checks the provided plain-text password
func ValidatePlain(plain string) error {
	if plain == "" {
		return fmt.Errorf("Provided password is empty")
	}
	return nil
}

// SHA1 creates a SHA-1 hash of the provided plain-text password
func SHA1(plain string) string {
	hash := sha1.Sum([]byte(plain))
	return strings.ToUpper(fmt.Sprintf("%x", hash))
}
