package adv_greetings

import (
	"regexp"
	"testing"
)

func TestHi(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hi("John")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hi("John") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHiEmpty(t *testing.T) {
	msg, err := Hi("")
	if msg != "" || err == nil {
		t.Fatalf(`Hi("") = %q %v want "", error`, msg, err)
	}
}
