package pwd

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
