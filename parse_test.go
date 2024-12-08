package main

import (
	//"strings"
	"testing"
)

// assume the position doesnt matter

func TestEscapeDollar(t *testing.T) {
	in := "$franzl$"
	want := "\\$franzl\\$"
	got := escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestEscapeSlash(t * testing.T) {
	in := "\\franzl"
	want := "\\\\franzl"
	got := escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestOpenBracket(t *testing.T) {
	in := "[franzl"
	want := "\\[franzl"
	got := escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestOpenBrackets(t *testing.T) {
	in := "[]franzl"
	want := "\\[\\]franzl"
	got := escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}	
}
func TestOpenBracketsDollar(t *testing.T) {
	in := "[$]franzl"
	want := "\\[\\$\\]franzl"
	got := escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}	
}
func TestParseQueryEscaped(t *testing.T) {
	want_key := "$format"
	query := "$" + escape(want_key) + "[0]"

	key, idx_char_map := parse_sort_query(query)
	got_idx, digits, err := parse_idx_operator(query, idx_char_map)
	key = unescape(key[:len(key) - digits])

	want_idx := 0
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if got_idx != want_idx {
		t.Errorf("\ngot:\t%d\nwanted:\t%d", got_idx, want_idx)
	}	
	got_dollar := idx_char_map[0]
	want_dollar := "$"
	if got_dollar != want_dollar {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got_dollar, want_dollar)
	}	
	if key != want_key {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", key, want_key)
	}
}



