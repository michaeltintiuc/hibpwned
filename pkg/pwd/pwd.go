package pwd

import (
	"io"
	"net/http"
)

// Pwd is an interface representing the ability to search compromised passwords
type Pwd interface {
	Search() error
	ValidateHash() error
}

// CheckPlain verifies if a plain-text password was compromised and how many times
func CheckPlain(pass string) (Plain, error) {
	p := Plain{Hash{"", false, 0}, pass}

	if err := p.ValidatePlain(); err != nil {
		return p, err
	}

	p.SHA1()

	return p, p.Search()
}

// CheckHash verifies if SHA-1 hash of a password was compromised and how many times
func CheckHash(hash string) (Hash, error) {
	p := Hash{hash, false, 0}
	return p, p.Search()
}

// FetchPwned passwords from the HIBPwned API
func FetchPwned(hash string) (io.ReadCloser, error) {
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + hash)
	return res.Body, err
}
