package pwd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

// NewHash creates a Hash instance
func NewHash(hash string) *Hash {
	return &Hash{hash, false, 0}
}

// Search the SHA-1 hash in in the list of compromised passwords
func (p *Hash) Search() error {
	pwned, err := p.FetchPwned()
	if err != nil {
		return err
	}
	defer pwned.Close()

	scanner := bufio.NewScanner(pwned)
	hashPart := p.Hashed[5:]

	for scanner.Scan() {
		if row := scanner.Text(); strings.Contains(row, hashPart) {
			p.ScanRow(row)
			break
		}
	}

	return scanner.Err()
}

// ScanRow for password data in the format of "hash:count"
func (p *Hash) ScanRow(row string) error {
	slice := strings.Split(row, ":")
	if len(slice) <= 1 {
		return fmt.Errorf("Malformed password data")
	}

	count, err := strconv.ParseFloat(slice[1], 10)
	if err != nil {
		return err
	}

	p.Pwned = true
	p.Count = int(count)

	return nil
}

// ValidateHash as a proper SHA-1 hash
func (p *Hash) ValidateHash() error {
	re := regexp.MustCompile("^[a-fA-F0-9]{40}$")
	if re.MatchString(p.Hashed) {
		p.Hashed = strings.ToUpper(p.Hashed)
		return nil
	}
	return fmt.Errorf("'%s' is not a valid SHA-1 hash", p.Hashed)
}

// FetchPwned passwords from the HIBPwned API
// using the first 5 characters of the SHA-1 hash
func (p Hash) FetchPwned() (io.ReadCloser, error) {
	if err := p.ValidateHash(); err != nil {
		return nil, err
	}
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + p.Hashed[:5])
	return res.Body, err
}
