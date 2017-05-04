package file

import (
	"testing"
)

func TestExist(t *testing.T) {
	t.Log("file.go exist：", Exist("file.go"))
	t.Log("file1.go exist：", Exist("file1.go"))
}
