package csv

import (
	"testing"
)

func TestNewCSV(t *testing.T) {
	c := NewCSV("test.csv")

	data := [][]string{
		{"天津", "1", "aa", "你好！"},
		{"北京", "2", "bb", "不好！"},
		{"乌鲁木齐", "3", "cc", "好吧！"},
	}

	c.Writer(data, true)

	t.Log("success!")
}
