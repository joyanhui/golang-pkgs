package exePath

import (
	"log"
	"os"
)

func GetExePath(isDev bool) string {
	if isDev {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("cannot get current working directory: %v", err)
		}
		return wd
	} else {
		executablePath, err := os.Executable()
		if err != nil {
			log.Fatalf("cannot get executable path: %v", err)
		}
		return executablePath
	}
}
