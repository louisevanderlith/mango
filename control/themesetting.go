package control

import (
	"os"

	"github.com/louisevanderlith/husk"
)

type ThemeSetting struct {
	LogoKey    husk.Key
	Name       string
	Host       string
	InstanceID string
}

func NewThemeSetting(name string, logoKey husk.Key, instanceID string) ThemeSetting {
	return ThemeSetting{
		Name:       name,
		LogoKey:    logoKey,
		Host:       os.Getenv("HOST"),
		InstanceID: instanceID,
	}
}
