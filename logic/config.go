package logic

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	Host           string
	MaxConnections int
	MaxWaiting     int
	Discovery      string
	Key            string
}

func (c *Config) LoadConfig(configPath string) {
	content := getFileContent(configPath)
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
		}
	}
}

func parse(value string) int {
	i, err := strconv.Atoi(value)

	if err != nil {
		log.Panic(err)
	}

	return i
}

func getFileContent(configPath string) []byte {
	dat, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Panic(err)
	}

	return dat
}
