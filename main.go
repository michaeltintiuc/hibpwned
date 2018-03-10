package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michaeltintiuc/hibpwned/pkg/pwd"
)

var (
	pass string
	hash string
)

func init() {
	flag.StringVar(&pass, "p", "", "Password to check")
	flag.StringVar(&hash, "h", "", "SHA-1 hash to check")
}

func main() {
	flag.Parse()

	if pass != "" {
		checkPlain()
	}

	if hash != "" {
		checkHash()
	}
}

func checkPlain() {
	fmt.Println("Checking plain-text password")

	p, err := pwd.CheckPlain(pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.Pwned {
		fmt.Printf("Your password was pwned %d times\n", p.Count)
		return
	}
	fmt.Println("You are secure, for now...")
}

func checkHash() {
	fmt.Println("Checking SHA-1 password hash")

	p, err := pwd.CheckHash(hash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.Pwned {
		fmt.Printf("Your password was pwned %d times\n", p.Count)
		return
	}
	fmt.Println("You are secure, for now...")
}
