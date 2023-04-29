package main

import (
	"bytes"
	"testing"
)

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2\n word3\n word4")

	exp := 4

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")

	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
