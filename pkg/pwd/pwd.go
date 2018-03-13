package pwd

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// CheckPlain verifies if a plain-text password was compromised and how many times
func CheckPlain(pass string) (*Hash, error) {
	p := NewHash("")

	if err := ValidatePlain(pass); err != nil {
		return p, err
	}

	p.Hashed = SHA1(pass)
	return p, p.Search()
}

// CheckHash verifies if SHA-1 hash of a password was compromised and how many times
func CheckHash(hash string) (*Hash, error) {
	p := NewHash(hash)
	return p, p.Search()
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
