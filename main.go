package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michaeltintiuc/hibpwned/pkg/breach/account"
	"github.com/michaeltintiuc/hibpwned/pkg/pwd"
)

var (
	pass  string
	hash  string
	email string
)

func init() {
	flag.StringVar(&pass, "p", "", "Password to check")
	flag.StringVar(&hash, "h", "", "SHA-1 hash to check")
	flag.StringVar(&email, "e", "", "Email to check")
}

func main() {
	flag.Parse()
	checkPassword()
	checkAccount()
}

func checkAccount() {
	if email == "" {
		return
	}

	data, err := account.Check(email)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if string(data) == "" {
		fmt.Printf("Account '%s' was not breached, yet...\n", email)
		return
	}

	fmt.Println(string(data))
}

func checkPassword() {
	checks := [2]struct {
		text  string
		value string
		fn    func(value string) (*pwd.Hash, error)
	}{
		{"Checking plain-text password", pass, pwd.CheckPlain},
		{"Checking SHA-1 password hash", hash, pwd.CheckHash},
	}

	for _, c := range checks {
		if c.value != "" {
			fmt.Println(c.text)
			validate(c.fn(c.value))
		}
	}
}

func validate(p *pwd.Hash, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if p.Pwned {
		fmt.Printf("Your password was pwned %d times\n", p.Count)
		return
	}

	fmt.Println("You are secure, for now...")
}
