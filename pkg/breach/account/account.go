package account

import (
	"io/ioutil"

	"github.com/michaeltintiuc/hibpwned/pkg/breach"
)

// Check if said account was breached
func Check(email string) ([]byte, error) {
	res, err := breach.Get("breachedaccount/" + email)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
