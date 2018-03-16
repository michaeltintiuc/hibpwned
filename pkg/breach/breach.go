package breach

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Breach is the interface that provides the shared Breach methods
type Breach interface {
	BuildURL() string
}

var baseURL = "https://haveibeenpwned.com/api/v2/"

// Get a HIBPwned API endpoint
func Get(endpoint string) (*http.Response, error) {
	c := &http.Client{}
	url := baseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "HIBPwned-Golang")
	return c.Do(req)
}

//VerifyAndRetry an API request
func VerifyAndRetry(res *http.Response) (bool, error) {
	retry, err := VerifyResponse(res.StatusCode)
	if err != nil {
		return false, err
	}
	if retry {
		err = Sleep(res.Header.Get("Retry-After"))
	}
	return retry, err
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

// Sleep after a timeed out response
func Sleep(seconds string) error {
	delay, err := strconv.ParseFloat(seconds, 10)
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(delay+1) * time.Second)
	return nil
}
