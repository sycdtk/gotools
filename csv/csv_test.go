package csv

import (
	"testing"
)

//func TestWriter(t *testing.T) {
//	c := NewCSV("test.csv")

//	data := [][]string{
//		{"天津", "1", "aa", "你好！"},
//		{"北京", "2", "bb", "不好！"},
//		{"乌鲁木齐", "3", "cc", "好吧！"},
//	}

//	c.Writer(data, true)

//	t.Log("success!")
//}

func TestReader(t *testing.T) {

	c := NewCSV("test.csv")
	c.Reader()

	for _, d := range c.Datas {
		for _, v := range d {
			t.Log(v)
		}
		t.Log("\n")
	}
}
