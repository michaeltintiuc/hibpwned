package pwd

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	domain       = "https://api.pwnedpasswords.com/"
	askForInput  = "Input password (hidden): "
	emptyPassErr = "\nCan't check an empty password"
)

// Pwd represents a container for the password in question
type Pwd struct {
	plain string
	hash  string
	pwned bool
	count int
}

// ValidatePass checks the provided password,
// if it is an empty string - user is prompted for input.
func (p *Pwd) ValidatePass() {
	if p.plain != "" {
		return
	}

	fmt.Print(askForInput)
	stdin, err := terminal.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	p.plain = strings.TrimSpace(string(stdin))

	if p.plain == "" {
		fmt.Fprintln(os.Stderr, emptyPassErr)
		os.Exit(2)
	}

	fmt.Println()
}

// Hash creates a SHA-1 hash of the password
func (p *Pwd) Hash() {
	hash := sha1.Sum([]byte(p.plain))
	p.hash = strings.ToUpper(fmt.Sprintf("%x", hash))
}

// Search if first 5 characters of the SHA-1 hash
// are found in the list of comporomised passwords
func (p *Pwd) Search() {
	pwned := FetchPwned(p.hash[:5])
	scanner := bufio.NewScanner(pwned)
	hashPart := p.hash[5:]
	defer pwned.Close()

	for scanner.Scan() {
		if row := scanner.Text(); strings.Contains(row, hashPart) {
			p.pwned = true

			count, err := strconv.ParseFloat(strings.Split(row, ":")[1], 10)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				defer func() { pwned.Close(); os.Exit(2) }()
			}

			p.count = int(count)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// Check if given password was compromised or not.
func Check(pass string) (bool, int) {
	p := Pwd{pass, "", false, 0}
	p.ValidatePass()
	p.Hash()
	p.Search()

	return p.pwned, p.count
}

// FetchPwned sends a request to the HIBPwned API
// and returns list of of compromised passwords
func FetchPwned(hash string) io.ReadCloser {
	res, err := http.Get(domain + "range/" + hash)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	return res.Body
}
