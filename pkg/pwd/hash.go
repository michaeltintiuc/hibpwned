package pwd

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Hash represents a SHA-1 hash of a password
type Hash struct {
	Hashed string
	Pwned  bool
	Count  int
}

// Search if the first 5 characters of the SHA-1 hash
// are found in the list of comporomised passwords
func (p *Hash) Search() error {
	if err := p.ValidateHash(); err != nil {
		return err
	}

	pwned, err := FetchPwned(p.Hashed[:5])
	if err != nil {
		return err
	}
	defer pwned.Close()

	scanner := bufio.NewScanner(pwned)
	hashPart := p.Hashed[5:]

	for scanner.Scan() {
		if row := scanner.Text(); strings.Contains(row, hashPart) {
			p.Pwned = true

			count, err := strconv.ParseFloat(strings.Split(row, ":")[1], 10)
			if err != nil {
				return err
			}

			p.Count = int(count)
			break
		}
	}

	return scanner.Err()
}

// ValidateHash is a proper SHA-1 hash
func (p *Hash) ValidateHash() error {
	re := regexp.MustCompile("^[a-fA-F0-9]{40}$")
	if re.MatchString(p.Hashed) {
		p.Hashed = strings.ToUpper(p.Hashed)
		return nil
	}
	return fmt.Errorf("'%s' is not a valid SHA-1 hash", p.Hashed)
}
