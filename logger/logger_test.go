package logger

import (
	"testing"
)

func TestPrint(t *testing.T) {
	SetLevel(L_DEBUG | L_ERROR)

	Info("Info……")
	Debug("Debug……")
	Error("Error……")
}
