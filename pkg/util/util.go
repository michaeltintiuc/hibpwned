package util

import "fmt"

// LogErr of passed functions, most often deferred and the error is of no real interest
func LogErr(f ...func() error) {
	for i := len(f) - 1; i >= 0; i-- {
		if err := f[i](); err != nil {
			fmt.Println("Encountered error:", err)
		}
	}
}
