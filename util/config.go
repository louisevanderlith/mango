package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Host           string
	MaxConnections int
	MaxWaiting     int
	Discovery      string
	Key            string
	Environment    string
}

func (c *Config) LoadConfig(configPath string) {
	content := GetFileContent(configPath)
	contentLines := strings.Split(string(content), "\r\n")

	for _, val := range contentLines {
		parts := strings.Split(val, "=")
		confKey := parts[0]
		confVal := strings.Trim(parts[1], " ")

		switch strings.Trim(strings.ToLower(confKey), " ") {
		case "httpport":
			port := fmt.Sprintf("localhost:%v", confVal)
			c.Host = port
		case "maxconnections":
			c.MaxConnections = parse(confVal)
		case "maxwaiting":
			c.MaxWaiting = parse(confVal)
		case "discovery":
			c.Discovery = confVal
		case "environment":
			c.Environment = confVal
		}
	}
}

func parse(value string) int {
	i, err := strconv.Atoi(value)

	if err != nil {
		log.Printf("util.parse: ", err)
	}

	return i
}

func GetFileContent(configPath string) []byte {
	dat, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Printf("GetFileContent: ", err)
	}

	return dat
}

func FindFilePath(fileName, targetFolder string) string {
	var result string
	wp := getWorkingPath() + "/" + targetFolder //"/conf"

	result = filepath.Join(wp, filepath.FromSlash(path.Clean("/"+fileName)))

	return result
}

func getWorkingPath() string {
	ex, err := os.Getwd()

	if err != nil {
		log.Printf("getWorkingPath: ", err)
	}

	return ex
}
