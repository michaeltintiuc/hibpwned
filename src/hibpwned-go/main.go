package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	domain = "https://api.pwnedpasswords.com/"
)

func main() {
	hash := sha1.Sum([]byte("P@ssw0rd"))
	hashString := strings.ToUpper(fmt.Sprintf("%x", hash))

	res, err := http.Get(domain + "range/" + hashString[:5])
	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(string(data), hashString[5:]) {
		fmt.Println("You've been PWNED!")
	} else {
		fmt.Println("You're secure, for now...")
	}
}
