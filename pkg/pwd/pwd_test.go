package pwd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Pwd(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HIBPwned servers should be happy now")
	}))
	defer ts.Close()

	pwds := [5]struct {
		Hash
		expectingErr bool
		fn           func(value string) (*Hash, error)
	}{
		{*NewHash("qwerty"), false, CheckPlain},
		{*NewHash("пароль"), false, CheckPlain},
		{*NewHash(""), true, CheckPlain},
		{*NewHash("B1B3773A05C0ED0176787A4F1574FF0075F7521E"), false, CheckHash},
		{*NewHash(""), true, CheckHash},
	}

	for i, pwd := range pwds {
		fmt.Printf("Running case %d\n", i+1)

		pwd.url = ts.URL
		_, err := pwd.fn(pwd.Hashed)
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
