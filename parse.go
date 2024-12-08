package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)


func get_escape_chars() []string {
	return []string{"$", "\\", "[", "]"}
}

func escape(in string) string {
	out := in
	chars := strings.Split(in, "")
	esc_count := 0
	for i := 0; i < len(chars); i++ {
		if slices.Contains(get_escape_chars(), chars[i]) {
			out = out[:i+esc_count] + "\\" + out[i+esc_count:]
			esc_count++
		}
	}
	return out
}

func parse_sort_query(query string) map[int]string {
	idx_char_map := make(map[int]string)
	// ignore all characters with \\ before the char
	spl := strings.Split(query, "")
	last := spl[0]
	if slices.Contains(get_escape_chars(), last) {
		idx_char_map[0] = last
	}
	for i := 1; i < len(spl); i++ {
		if slices.Contains(get_escape_chars(), spl[i]) && last != "\\" && spl[i] != "\\" {
			// encountered escaped char
			idx_char_map[i] = spl[i]
		}
		last = spl[i]
	}
	return idx_char_map
}
func parse_idx_operator(query string, idx_char_map map[int]string) (int, error) {
	open := -1
	close := -1
	for k,v := range idx_char_map {
		if v == "[" {
			open = k
		}
		if v == "]" {
			close = k
		}
	}
	if open == -1 || close == -1 {
		return -1, errors.New(`either "[" or "]" not found`)
	}
	if open > close {
		return -1, errors.New(`open bracket after close bracket`)
	}
	idx_str := query[open+1:close]
	idx, err := strconv.Atoi(idx_str)
	return idx, err
}
