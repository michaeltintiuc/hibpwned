package pwd

import (
	"fmt"
	"testing"
)

func Test_Pwd(t *testing.T) {
	pwds := [5]struct {
		value        string
		expectingErr bool
		fn           func(value string) (*Hash, error)
	}{
		{"qwerty", false, CheckPlain},
		{"пароль", false, CheckPlain},
		{"", true, CheckPlain},
		{"B1B3773A05C0ED0176787A4F1574FF0075F7521E", false, CheckHash},
		{"", true, CheckHash},
	}

	for i, pwd := range pwds {
		fmt.Printf("Running case %d\n", i+1)

		_, err := pwd.fn(pwd.value)
		if pwd.expectingErr {
			if err == nil {
				t.Errorf("Expecting an error in case %d\n", i+1)
			} else {
				fmt.Println(err)
			}
			continue
		}

		if err != nil {
			t.Error(err)
		}
	}
}

func Test_ScanRow(t *testing.T) {
	h := NewHash("B1B3773A05C0ED0176787A4F1574FF0075F7521E")

	if err := h.ScanRow("hash"); err == nil {
		t.Error("Expected malformed data")
	}

	if err := h.ScanRow("hash:"); err == nil {
		t.Error("Expected malformed data")
	}
}
