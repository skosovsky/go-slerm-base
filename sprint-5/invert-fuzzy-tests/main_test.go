package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzInvertLine(f *testing.F) {
	testCases := []struct {
		result, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		f.Add(tc.result)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := invertLine(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := invertLine(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("got %q, want %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("got %q, want utf8.ValidString", rev)
		}
	})
}
