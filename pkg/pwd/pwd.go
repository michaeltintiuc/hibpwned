package pwd

import (
	"bufio"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Pwd represents a container for the password in question
type Pwd struct {
	plain string
	hash  string
	pwned bool
	count int
}

// Search if first 5 characters of the SHA-1 hash
// are found in the list of comporomised passwords
func (p *Pwd) Search() error {
	if err := p.ValidateHash(); err != nil {
		return err
	}

	pwned, err := FetchPwned(p.hash[:5])
	if err != nil {
		return err
	}
	defer pwned.Close()

	scanner := bufio.NewScanner(pwned)
	hashPart := p.hash[5:]

	for scanner.Scan() {
		if row := scanner.Text(); strings.Contains(row, hashPart) {
			p.pwned = true

			count, err := strconv.ParseFloat(strings.Split(row, ":")[1], 10)
			if err != nil {
				return err
			}

			p.count = int(count)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// FetchPwned sends a request to the HIBPwned API
// and returns list of of compromised passwords
func FetchPwned(hash string) (io.ReadCloser, error) {
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + hash)
	return res.Body, err
}
