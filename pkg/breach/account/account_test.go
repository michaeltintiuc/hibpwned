package account

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	e = "test@example.com"
	d = "adobe.com"
)

func Test_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch p := r.URL.Query().Get("domain"); p {
		case "retry":
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(429)
		case "fail":
			w.WriteHeader(400)
		}
		fmt.Fprintln(w, "HIBPwned servers should be happy now")
	}))
	defer ts.Close()

	cases := []struct {
		Account
		expectingErr bool
	}{
		{*NewAccount(e, "", false, false), false},
		{*NewAccount(e, d, true, true), false},
		{*NewAccount("", d, false, false), true},
		{*NewAccount(e, "retry", false, false), false},
		{*NewAccount(e, "fail", false, false), true},
	}

	for i, c := range cases {
		fmt.Printf("Running case %d\n", i+1)
		c.url = ts.URL + "?domain=" + c.domain
		_, err := c.Check()

		if c.expectingErr {
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

func Test_FetchBreached(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HIBPwned servers should be happy now")
	}))
	defer ts.Close()

	cases := []struct {
		Account
		URL          string
		expectingErr bool
	}{
		{*NewAccount(e, d, false, false), ts.URL, false},
		{*NewAccount(e, d, false, false), ":", true},
		{*NewAccount("", "foo", false, false), ts.URL, true},
	}

	for i, c := range cases {
		fmt.Printf("Running case %d\n", i+1)
		c.Account.url = c.URL
		res, err := c.FetchBreached()

		if c.expectingErr {
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
		_, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_Format(t *testing.T) {
	a := NewAccount(e, d, false, false)
	cases := []struct {
		data         []byte
		expectingErr bool
	}{
		{[]byte(""), true},
		{[]byte(`[{"Name":"Adobe"}]`), false},
	}

	for i, c := range cases {
		fmt.Printf("Running case %d\n", i+1)
		_, err := a.Format(c.data)

		if c.expectingErr {
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
