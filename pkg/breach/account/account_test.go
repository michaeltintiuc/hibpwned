package account

import (
	"fmt"
	"testing"
	"time"
)

var (
	e = "test@example.com"
	d = "adobe.com"
)

func Test_Check(t *testing.T) {
	cases := []struct {
		Account
		sleep        bool
		expectingErr bool
	}{
		{Account{" ", d, false, false}, false, true},
		{Account{e, "", false, false}, true, false},
		{Account{e, d, false, false}, false, false},
		{Account{e, d, true, true}, true, false},
		{Account{"", d, false, false}, false, true},
		{Account{"", "foo", false, false}, true, true},
	}

	for i, c := range cases {
		fmt.Printf("Running case %d\n", i+1)
		_, err := Check(c.email, c.domain, c.truncated, c.unverified)

		if c.expectingErr == true {
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

		if c.sleep {
			time.Sleep(4 * time.Second)
		}
	}
}

func Test_FetchBreached(t *testing.T) {
	cases := []struct {
		Account
		expectingErr bool
	}{
		{Account{"", "foo", false, false}, true},
	}

	for i, c := range cases {
		fmt.Printf("Running case %d\n", i+1)
		_, err := c.FetchBreached()

		if c.expectingErr == true {
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
