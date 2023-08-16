package main

import (
	"testing"
	"unicode/utf8"
)

/* Previous test function for reverse, without using Fuzz
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
*/

/* The unit test above has limitations, namely that each input must be added to the test by the developer.
One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test
cases you came up with didn’t reach.
In this section you will convert the unit test to a fuzz test so that you can generate more inputs with less work
*/

func FuzzReverse(f *testing.F) {
	/*
		Fuzzing has a few limitations as well. In your unit test, you could predict the expected output of the Reverse
		function, and verify that the actual output met those expectations.
		For example, in the test case Reverse("Hello, world") the unit test specifies the return as "dlrow ,olleH".

		When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.

		However, there are a few properties of the Reverse function that you can verify in a fuzz test. The two
		properties being checked in this fuzz test are:

		- Reversing a string twice preserves the original value
		- The reversed string preserves its state as valid UTF-8.
		- testcases := []string{"Hello, word", " ", "!12345"}
	*/
	testcases := []string{"Hello, world", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	// takes a fuzz target function whose parameters are *testing.T and the types to be fuzzed. T
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		// reverse the string rune-by-rune
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
