package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	domain = "https://api.pwnedpasswords.com/"
)

var (
	p string
)

func init() {
	flag.StringVar(&p, "p", "", "Password to check")
}

func main() {
	flag.Parse()

	if p == "" {
		flag.PrintDefaults()
		return
	}

	hash := sha1.Sum([]byte(p))
	hashString := strings.ToUpper(fmt.Sprintf("%x", hash))
	res, err := http.Get(domain + "range/" + hashString[:5])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

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
