package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a Nname, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "rob-o"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("rob-o")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("rob-o") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// The text expects empty message or an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
