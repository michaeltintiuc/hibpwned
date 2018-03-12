package pwd

import "testing"

func Test_CheckPlain(t *testing.T) {
	pwds := [3]struct {
		plainPass    string
		expectingErr bool
	}{
		{"qwerty", false},
		{"пароль", false},
		{"", true},
	}

	for _, pwd := range pwds {
		_, err := CheckPlain(pwd.plainPass)

		if pwd.expectingErr {
			if err == nil {
				t.Error("Expected an error")
			}
			return
		}

		if err != nil {
			t.Error(err)
		}
	}
}

func Test_CheckHash(t *testing.T) {
	pwds := [2]struct {
		hash         string
		expectingErr bool
	}{
		{"B1B3773A05C0ED0176787A4F1574FF0075F7521E", false},
		{"", true},
	}

	for _, pwd := range pwds {
		_, err := CheckHash(pwd.hash)

		if pwd.expectingErr {
			if err == nil {
				t.Error("Expected an error")
			}
			return
		}

		if err != nil {
			t.Error(err)
		}
	}
}

func Test_ScanRow(t *testing.T) {
	h := NewHash("B1B3773A05C0ED0176787A4F1574FF0075F7521E")

	if err := h.ScanRow("hash"); err == nil {
		t.Errorf("Expected malformed data")
	}

	if err := h.ScanRow("hash:"); err == nil {
		t.Errorf("Expected malformed data")
	}
}
