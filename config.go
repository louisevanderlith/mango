package mango

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetFileContent(configPath string) []byte {
	dat, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Print("GetFileContent: ", err)
	}

	return dat
}

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
