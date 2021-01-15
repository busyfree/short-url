package main

import "testing"

var n, _ = NewUrlEncoder("", 0)

func TestNewUrlEncoder(t *testing.T) {
	_, err := NewUrlEncoder("s", 0)
	if err != nil {
		t.Error(err)
	}
}

func TestEncodeUrl(t *testing.T) {
	n, _ := NewUrlEncoder("", 0)
	s := n.EncodeUrl(122, 1)
	if s != "4sty6" {
		t.Error("encodeurl error")
	}
}

func TestDecodeUrl(t *testing.T) {
	n, _ := NewUrlEncoder("", 0)
	s := n.DecodeUrl([]rune("jf"))
	if s != 122 {
		t.Error("decodeurl error")
	}
}

func TestEnbase(t *testing.T) {
	n, _ := NewUrlEncoder("", 0)
	s := n.EnBase(122, 1)
	if s != "jf" {
		t.Error("enbase error")
	}
}

func TestDebase(t *testing.T) {
	n, _ := NewUrlEncoder("", 0)
	s := n.DeBase([]rune("jf"))
	if s != 122 {
		t.Error("debase error")
	}
}
