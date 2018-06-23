package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michaeltintiuc/hibpwned/pkg/breach/account"
	"github.com/michaeltintiuc/hibpwned/pkg/pwd"
)

var (
	pass       string
	hash       string
	email      string
	domain     string
	truncated  bool
	unverified bool
	formatJSON bool
)

func init() {
	flag.StringVar(&pass, "p", "", "Password to check")
	flag.StringVar(&hash, "h", "", "SHA-1 hash to check")
	flag.StringVar(&email, "e", "", "Email or username to check")
	flag.StringVar(&domain, "d", "", "Domain to check email against")
	flag.BoolVar(&truncated, "t", false, "Display less detailed email breach info")
	flag.BoolVar(&unverified, "u", false, "Include unverified email breaches")
	flag.BoolVar(&formatJSON, "f", false, "Format JSON output")
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

	a := account.NewAccount(email, domain, truncated, unverified)
	data, err := a.Check()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Printf("Account '%s' was not breached, yet...\n", email)
		return
	}

	if !formatJSON {
		fmt.Println(string(data))
		return
	}

	dataFormatted, err := a.Format(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, c := range dataFormatted {
		if i > 0 {
			fmt.Println("---------------------")
		}

		fmt.Println("Title:", c.Title)
		fmt.Println("Domain:", c.Domain)
		fmt.Println("Date:", c.BreachDate)
		fmt.Println("Count:", c.PwnCount)
	}
}

func checkPassword() {
	checks := [2]struct {
		text  string
		value string
		fn    func(value string) (*pwd.Hash, error)
	}{
		{"Checking plain-text password", pass, pwd.NewPlain},
		{"Checking SHA-1 password hash", hash, pwd.NewHash},
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

	if err = p.Search(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if p.Pwned {
		fmt.Printf("Your password was pwned %d times\n", p.Count)
		return
	}

	fmt.Println("You are secure, for now...")
}
