package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	Push(s, "aaa")
	Push(s, "b")
	Push(s, "ww")

	if v := Pop(s); v == "ww" {
		t.Log(v, "ww")
	} else {
		t.Fatal(v, "ww")
	}
	Push(s, "qq")
	Pop(s)
	Pop(s)
	if v := Pop(s); v == "aaa" {
		t.Log(v, "aaa")
	} else {
		t.Fatal(v, "aaa")
	}
}
