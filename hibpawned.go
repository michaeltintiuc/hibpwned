package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michaeltintiuc/hibpwned/pkg/pwd"
)

var (
	pass string
)

func init() {
	flag.StringVar(&pass, "p", "", "Password to check")
}

func main() {
	flag.Parse()
	pwned, count := pwd.Check(pass)

	if pwned {
		fmt.Printf("Your password was pwned %d times\n", count)
		os.Exit(1)
	}

	fmt.Println("You are secure, for now...")
}
