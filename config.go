package mango

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

//Returns the filepath within the current working directory.
func FindFilePath(fileName, targetFolder string) string {
	var result string
	wp := getWorkingPath() + "/" + targetFolder

	result = filepath.Join(wp, filepath.FromSlash(path.Clean("/"+fileName)))

	return result
}

func getWorkingPath() string {
	ex, err := os.Getwd()

	if err != nil {
		log.Print("getWorkingPath: ", err)
	}

	return ex
}
