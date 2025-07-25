package exePath

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetExeDir(isDev bool) string {
	if isDev {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("[GetExeDir] cannot get current working directory: %v \n", err)
		}
		return wd
	} else {
		executablePath, err := os.Executable()
		if err != nil {
			fmt.Printf("[GetExeDir] cannot get executable path: %v \n", err)
		}
		return filepath.Dir(executablePath)
	}
}
