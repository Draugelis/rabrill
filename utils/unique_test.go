package utils

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestUnique(t *testing.T) {
	s := []string{"unique1", "unique1", "unique2"}
	want := []string{"unique1", "unique2"}
	res := Unique(s)
	if !slices.Equal(res, want) {
		t.Errorf("Failed Unique(%v). Got: %v. Want: %v", s, res, want)
	}
}
