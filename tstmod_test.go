/**
 * Author: Pascal van Dam
 * File: recursivefunc_test.go
 */

package tstmod

import (
	"github.com/pascal71/tstmod"
	"testing"
)

func TestFac(t *testing.T) {
	var a uint64
	var expected uint64

	a = 4
	expected = 4 * 3 * 2 * 1

	if got := recursivefunc.Fac(a); got != expected {
		t.Errorf("Fac(%d) = %d, didn't return %d", a, got, expected)
	}
}
