package core_test

import (
	"testing"
	"jesao++/core"
)

// assume the position doesnt matter
var parser core.Parser = core.Parser{}

func TestEscapeDollar(t *testing.T) {
	in := "$franzl$"
	want := "\\$franzl\\$"
	got := parser.Escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestEscapeSlash(t * testing.T) {
	in := "\\franzl"
	want := "\\\\franzl"
	got := parser.Escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestOpenBracket(t *testing.T) {
	in := "[franzl"
	want := "\\[franzl"
	got := parser.Escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}
}
func TestOpenBrackets(t *testing.T) {
	in := "[]franzl"
	want := "\\[\\]franzl"
	got := parser.Escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}	
}
func TestOpenBracketsDollar(t *testing.T) {
	in := "[$]franzl"
	want := "\\[\\$\\]franzl"
	got := parser.Escape(in)
	if got != want {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
	}	
}
func TestParseQuerySeparate(t *testing.T) {
	want_key := "$format"
	query := "$" + parser.Escape(want_key) + "[0]"

	key, idx_char_map := parser.ParseSortQuery(query)
	got_idx, digits, err := parser.ParseIdxOperator(query, idx_char_map)
	key = parser.Unescape(key[:len(key) - digits])

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
func TestParseQuery(t *testing.T) {
	want_key := "$[]format"
	want_idx := 0
	query := "$" + parser.Escape(want_key) + "[0]"
	key, idx, err := parser.ParseKeyAndIdx(query)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	if idx != want_idx {
		t.Errorf("\ngot:\t%d\nwanted:\t%d", idx, want_idx)
	}
	if key != want_key {
		t.Errorf("\ngot:\t%s\nwanted:\t%s", key, want_key)
	}

}



