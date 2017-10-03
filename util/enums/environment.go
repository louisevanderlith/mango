package enums

import (
	"strings"
)

// Environment provides indicates in which environment a system is
type Environment int

const (
	LIVE Environment = iota
	UAT
	DEV
)

var environments = [...]string{
	"LIVE",
	"UAT",
	"DEV"}

func (e Environment) String() string {
	return environments[e]
}

func GetEnvironment(name string) Environment {
	var result Environment

	for k, v := range environments {
		if strings.ToUpper(name) == v {
			result = Environment(k)
			break
		}
	}

	return result
}
