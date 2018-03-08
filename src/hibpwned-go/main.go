package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	domain = "https://api.pwnedpasswords.com/"
)

var (
	pass string
)

func init() {
	flag.StringVar(&pass, "p", "", "Password to check")
}

func main() {
	flag.Parse()

	if pass == "" {
		fmt.Print("Input password (hidden): ")
		p, err := terminal.ReadPassword(int(os.Stdin.Fd()))

		if err != nil {
			fmt.Println(err)
			return
		}

		pass = strings.TrimSpace(string(p))

		if pass == "" {
			fmt.Println("\nCan't check an empty password")
			return
		}

		println()
	}

	hash := sha1.Sum([]byte(pass))
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
