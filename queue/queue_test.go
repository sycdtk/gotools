package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	Put(q, "aaa")
	Put(q, "bbb")
	Put(q, "ccc")

	if v := Take(q); v == "aaa" {
		t.Log(v, "aaa")
	} else {
		t.Fatal(v, "aaa")
	}
	Put(q, "ddd")
	Take(q)
	Take(q)
	if v := Take(q); v == "ddd" {
		t.Log(v, "ddd")
	} else {
		t.Fatal(v, "ddd")
	}
}
