package pwd

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// Plain represents a plain-text password
// extending the pwd.Hash struct
type Plain struct {
	Hash
	Plain string
}

// ValidatePlain checks the provided plain-text password
func (p *Plain) ValidatePlain() error {
	p.Plain = strings.TrimSpace(p.Plain)
	if p.Plain == "" {
		return fmt.Errorf("Provided password is empty")
	}
	return nil
}

// SHA1 creates a SHA-1 hash of the provided plain-text password
func (p *Plain) SHA1() {
	hash := sha1.Sum([]byte(p.Plain))
	p.Hashed = strings.ToUpper(fmt.Sprintf("%x", hash))
}
