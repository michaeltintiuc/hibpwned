package breach

import (
	"fmt"
	"net/http"
)

// Breach is the interface that provides the shared Breach methods
type Breach interface {
	BuildURL() string
}

// Get a HIBPwned API endpoint
func Get(endpoint string) (*http.Response, error) {
	c := &http.Client{}
	url := "https://haveibeenpwned.com/api/v2/" + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "HIBPwned-Golang")
	return c.Do(req)
}

// VerifyResponse status code
func VerifyResponse(status int) (bool, error) {
	switch status {
	case 429:
		return true, nil
	case 200:
		return false, nil
	default:
		return false, fmt.Errorf("Received %d response", status)
	}
}
