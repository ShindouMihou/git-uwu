package utils

import (
	"log"
	"os"
)

var CurrentDirectory string

func GetCurrentDirectory() string {
	if CurrentDirectory == "" {
		executable, err := os.Getwd()
		if err != nil {
			log.Fatal("failed to get current directory : ", err)
		}
		CurrentDirectory = executable
	}
	return CurrentDirectory
}
