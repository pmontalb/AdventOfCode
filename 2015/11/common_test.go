package main

import "testing"

func TestIsValid(t *testing.T) {
	if IsValid("hijklmmn") {
		t.Errorf("hijklmmn is not valid")
	}
	if IsValid("abbceffg") {
		t.Errorf("abbceffg is not valid")
	}
	if IsValid("abbcegjk") {
		t.Errorf("abbcegjk is not valid")
	}
	if !IsValid("abcdffaa") {
		t.Errorf("abbcegjk is valid")
	}
	if !IsValid("ghjaabcc") {
		t.Errorf("ghjaabcc is valid")
	}
	if IsValid("abcdeggg") {
		t.Errorf("ghjaabcc is not valid")
	}
}

func TestNextPassword(t *testing.T) {
	t1 := FindNextPassword("abcdefgh")
	if t1 != "abcdffaa" {
		t.Errorf("abcdffaa: next is incorrect, got %s, want %s", t1, "abcdffaa")
	}
	t2 := FindNextPassword("ghijklmn")
	if t2 != "ghjaabcc" {
		t.Errorf("abcdffaa: next is incorrect, got %s, want %s", t2, "ghjaabcc")
	}
}
