package main

import "testing"

func toRune(s string) rune {
	return []rune(s)[0]
}

var getIndexTests = []struct{
	in string
	expected int
}{
	{"a", 0},
	{"z", 25},
	{"A", 0},
	{"Z", 25},
}

func TestGetIndex(t *testing.T) {
	for _, c := range getIndexTests {
		if out := getIndex(toRune(c.in)); out != c.expected {
			t.Errorf("getIndex(%v) = %v, expected %v", c.in, out, c.expected)
		}
	}
}

var convertTests = []struct{
	input string
	chartype string
	expected string
}{
	{"helLo", "smallcaps", "ʜᴇʟʟᴏ"},
	{"helLo", "spacedcase", "ｈｅｌｌｏ"},
	{"hElLo", "lower", "hello"},
	{"hElLo", "upper", "HELLO"},
}

func TestConvert(t *testing.T) {
	for _, c := range convertTests {
		if out := convert(c.chartype, c.input); out != c.expected {
			t.Errorf("convert(%v, %v) = %v, expected %v", c.chartype, c.input, out, c.expected)
		}
	}
}
