package csv

import (
	"testing"
)

func TestNewCSV(t *testing.T) {
	c := NewCSV("a.csv")

	data := [][]string{
		{"1", "aa", "你好！"},
		{"2", "bb", "不好！"},
		{"3", "cc", "好吧！"},
	}

	c.Writer(data)

	t.Log("success!")
}
