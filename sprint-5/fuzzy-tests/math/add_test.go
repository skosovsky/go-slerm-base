package math_test

import (
	"testing"

	"github.com/skosovsky/go-slerm-base/sprint-5/fuzzy-tests/math"
)

// go test -fuzz=FuzzTestAddWithError ./math
// Указывается явно название функции и папка где функция лежит

// warning: starting with empty corpus
// fuzz: elapsed: 0s, execs: 0 (0/sec), new interesting: 0 (total: 0)
// fuzz: elapsed: 0s, execs: 39 (1184/sec), new interesting: 0 (total: 0)
// --- FAIL: FuzzTestAddWithError (0.04s)
// --- FAIL: FuzzTestAddWithError (0.00s)
// add_test.go:13: Add(-87, 10) = 0, want -77
//
// Failing input written to testdata/fuzz/FuzzTestAddWithError/b9b58ee17bfdd921
// To re-run:
// go test -run=FuzzTestAddWithError/b9b58ee17bfdd921
// FAIL
// exit status 1
// FAIL    github.com/skosovsky/go-slerm-base/sprint-5/fuzzy-tests/math    0.340s

func FuzzTestAddWithError(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		res := math.AddWithError(a, b)
		if res != a+b {
			t.Errorf("Add(%d, %d) = %d, want %d", a, b, res, a+b)
		}
	})
}
