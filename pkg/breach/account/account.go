package account

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/michaeltintiuc/hibpwned/pkg/breach"
)

// Account to verify for breaches
type Account struct {
	email      string
	domain     string
	truncated  bool
	unverified bool
}

// NewAccount creates a new Account instance
func NewAccount(email, domain string, truncated, unverified bool) *Account {
	return &Account{email, domain, truncated, unverified}
}

// BuildURL to send request to
func (a Account) BuildURL() string {
	url := "breachedaccount/" + a.email
	params := []string{}

	if a.domain != "" {
		params = append(params, "domain="+a.domain)
	}
	if a.truncated == true {
		params = append(params, "truncateResponse=true")
	}
	if a.unverified == true {
		params = append(params, "includeUnverified=true")
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))
}

// Check if said account was breached
func Check(email, domain string, truncated, unverified bool) ([]byte, error) {
	a := NewAccount(email, domain, truncated, unverified)
	res, err := breach.Get(a.BuildURL())
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
