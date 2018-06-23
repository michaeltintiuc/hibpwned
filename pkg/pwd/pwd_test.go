package pwd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Pwd(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "B1B3773A05C0ED0176787A4F1574FF0075F7521E:123\r\nB1B3773A05C0ED0176787A4F1574FF0075F7521A:")
	}))
	defer ts.Close()

	pwds := []struct {
		str          string
		URL          string
		isPlain      bool
		expectingErr bool
	}{
		{"qwerty", ts.URL, true, false},
		{"пароль", ts.URL, true, false},
		{"", ts.URL, true, true},
		{"B1B3773A05C0ED0176787A4F1574FF0075F7521E", ts.URL, false, false},
		{"B1B3773A05C0ED0176787A4F1574FF0075F7521A", ts.URL, false, true},
		{"B1B3773A05C0ED0176787A4F1574FF0075F7521A", "foo", false, true},
		{"", ts.URL, false, true},
	}

	var p *Hash

	for i, pwd := range pwds {
		fmt.Printf("Running case %d\n", i+1)

		var err error
		if pwd.isPlain {
			p, err = NewPlain(pwd.str)
		} else {
			p, err = NewHash(pwd.str)
		}

		if err != nil {
			if pwd.expectingErr {
				fmt.Println(err)
				continue
			} else {
				t.Error(err)
			}
		}

		p.url = pwd.URL + "/"
		err = p.Search()
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
	h, _ := NewHash("B1B3773A05C0ED0176787A4F1574FF0075F7521E")

	if err := h.ScanRow("hash"); err == nil {
		t.Error("Expected malformed data")
	}

	if err := h.ScanRow("hash:"); err == nil {
		t.Error("Expected malformed data")
	}
}
