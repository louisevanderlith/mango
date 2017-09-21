package logic

import "io/ioutil"
import "strings"
import "strconv"
import "fmt"

type Config struct {
	Host           string
	MaxConnections int
	MaxWaiting     int
	Discovery      string
	Key            string
}

func (c *Config) LoadConfig(configPath string) {
	content := getFileContent(configPath)

	for _, val := range content {
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
		panic(err)
	}

	return i
}

func getFileContent(configPath string) []string {
	dat, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\r\n")

	return lines
}
