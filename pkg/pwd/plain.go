package pwd

import (
	"fmt"
	"strings"
)

// CheckPlain verifies if plain-text password was compromised
// and how many times
func CheckPlain(pass string) (bool, int, error) {
	p := Pwd{pass, "", false, 0}

	if err := p.ValidatePlain(); err != nil {
		return false, 0, err
	}

	p.Hash()

	if err := p.Search(); err != nil {
		return false, 0, err
	}

	return p.pwned, p.count, nil
}

// ValidatePlain checks the provided password
func (p *Pwd) ValidatePlain() error {
	p.plain = strings.TrimSpace(p.plain)
	if p.plain == "" {
		return fmt.Errorf("Provided password is empty")
	}
	return nil
}
