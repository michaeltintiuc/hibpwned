package pwd

import (
	"crypto/sha1"
	"fmt"
	"regexp"
	"strings"
)

// CheckHash verifies if SHA-1 hash of a password was compromised
// and how many times
func CheckHash(hash string) (bool, int, error) {
	p := Pwd{"", hash, false, 0}

	if err := p.Search(); err != nil {
		return false, 0, err
	}

	return p.pwned, p.count, nil
}

// Hash creates a SHA-1 hash of the password
func (p *Pwd) Hash() {
	hash := sha1.Sum([]byte(p.plain))
	p.hash = strings.ToUpper(fmt.Sprintf("%x", hash))
}

// ValidateHash checks if the provided value is a valid SHA-1 hash
func (p *Pwd) ValidateHash() error {
	re := regexp.MustCompile("^[a-fA-F0-9]{40}$")
	if re.MatchString(p.hash) {
		p.hash = strings.ToUpper(p.hash)
		return nil
	}
	return fmt.Errorf("'%s' is not a valid SHA-1 hash", p.hash)
}
