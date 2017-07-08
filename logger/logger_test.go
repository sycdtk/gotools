package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	l1 := NewLogger()
	l2 := NewLogger()

	l1.Debug("1 - aaaa")
	l1.Debug("1 - bbbb")
	l1.Debug("1 - cccc")

	l2.Debug("2 - aaaa")

	l1.Info("1 info - dddd")
	l1.Debug("1 - eeee")
	l2.Info("2 info bbbb")

	l2.SetLevel(Linfo)

	l1.Info("11 info - ffff")
	l1.Debug("1 - gggg")
	l2.Info("2 info cccc")
}
