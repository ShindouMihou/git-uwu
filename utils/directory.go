package utils

import (
	"log"
	"os"
	"path/filepath"
)

var CurrentDirectory string

func GetCurrentDirectory() string {
	if CurrentDirectory == "" {
		executable, err := os.Executable()
		if err != nil {
			log.Fatal("failed to get current directory : ", err)
		}
		CurrentDirectory = filepath.Dir(executable)
	}
	return CurrentDirectory
}
