package exePath

import (
	"testing" // [test] 导入 testing 包
)

func TestGetExePath(t *testing.T) {
	t.Logf("Executable path (with dev): %s", GetExeDir(true))
	t.Logf("Executable path (without dev): %s", GetExeDir(false))
}
