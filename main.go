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
	checkPlain()
	checkHash()
}

func checkPlain() {
	if pass == "" {
		return
	}

	fmt.Println("Checking plain-text password")

	pwned, count, err := pwd.CheckPlain(pass)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if pwned {
		fmt.Printf("Your password was pwned %d times\n", count)
		return
	}

	fmt.Println("You are secure, for now...")
}

func checkHash() {
	if hash == "" {
		return
	}

	fmt.Println("Checking SHA-1 password hash")
	pwned, count, err := pwd.CheckHash(hash)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if pwned {
		fmt.Printf("Your password was pwned %d times\n", count)
		return
	}

	fmt.Println("You are secure, for now...")
}
