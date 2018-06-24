package util

import (
	"errors"
)

func ExampleLogErr() {
	LogErr(GetErr, NoErr)
	// Output: Encountered error: foobar
}

func GetErr() error {
	return errors.New("foobar")
}

func NoErr() error {
	return nil
}
