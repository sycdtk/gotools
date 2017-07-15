package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {

	Debug("1 - aaaa")
	Debug("1 - bbbb")
	Debug("1 - cccc")

	Debug("2 - aaaa")

	Info("1 info - dddd")
	Debug("1 - eeee")
	Info("2 info bbbb")

	SetLevel("INFO")

	Info("11 info - ffff")
	Debug("1 - gggg")
	Info("2 info cccc")
}
