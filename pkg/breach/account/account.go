package account

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

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

// FetchBreached account info
func (a Account) FetchBreached() (*http.Response, error) {
	if a.email == "" {
		return nil, fmt.Errorf("Cannot fetch empty account breaches")
	}
	return breach.Get(a.BuildURL())
}

// RetryRequest after a timeout response
func (a Account) RetryRequest(header string) error {
	delay, err := strconv.ParseFloat(header, 10)
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(delay+1) * time.Second)
	return nil
}

// Check if said account was breached
func Check(email, domain string, truncated, unverified bool) ([]byte, error) {
	a := NewAccount(email, domain, truncated, unverified)
RETRY:
	breached, err := a.FetchBreached()
	if err != nil {
		return []byte{}, err
	}

	retry, err := breach.VerifyResponse(breached.StatusCode)
	if err != nil {
		return []byte{}, err
	}
	if retry {
		goto RETRY
	}

	defer breached.Body.Close()
	return ioutil.ReadAll(breached.Body)
}
