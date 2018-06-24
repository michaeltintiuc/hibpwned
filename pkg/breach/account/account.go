package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/michaeltintiuc/hibpwned/pkg/breach"
	"github.com/michaeltintiuc/hibpwned/pkg/util"
)

// Account to verify for breaches
type Account struct {
	URL        string
	Email      string
	Domain     string
	Truncated  bool
	Unverified bool
}

// JSON structure of breached account
type JSON struct {
	Title        string
	Name         string
	Domain       string
	BreachDate   string
	AddedDate    string
	ModifiedDate string
	Description  string
	LogoType     string
	DataClasses  []string
	PwnCount     int
	IsVerified   bool
	IsFabricated bool
	IsActive     bool
	IsRetired    bool
	IsSpamList   bool
}

// NewAccount creates a new Account instance
func NewAccount(email, domain string, truncated, unverified bool) *Account {
	a := &Account{"", email, domain, truncated, unverified}
	a.URL = a.BuildURL()
	return a
}

// BuildURL to send request to
func (a Account) BuildURL() string {
	url := breach.BaseURL + "breachedaccount/" + a.Email
	params := []string{}

	if a.Domain != "" {
		params = append(params, "domain="+a.Domain)
	}
	if a.Truncated {
		params = append(params, "truncateResponse=true")
	}
	if a.Unverified {
		params = append(params, "includeUnverified=true")
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))
}

// FetchBreached account info
func (a Account) FetchBreached() (*http.Response, error) {
	if a.Email == "" {
		return nil, fmt.Errorf("Cannot fetch empty account breaches")
	}
	return breach.Get(a.URL)
}

// Check if said account was breached
func (a Account) Check() ([]byte, error) {
	retries := 0
RETRY:
	breached, err := a.FetchBreached()
	if err != nil {
		return []byte{}, err
	}

	retry, err := breach.VerifyAndRetry(breached)
	if err != nil {
		return []byte{}, err
	}
	if retry && retries < breach.MaxRetries {
		retries++
		goto RETRY
	}

	defer util.LogErr(breached.Body.Close)
	return ioutil.ReadAll(breached.Body)
}

// Format the details of an account breach
func (a Account) Format(breachData []byte) ([]JSON, error) {
	var accountJSON []JSON
	if len(breachData) == 0 {
		return accountJSON, fmt.Errorf("Cannot format empty string")
	}
	err := json.Unmarshal(breachData, &accountJSON)
	return accountJSON, err
}
