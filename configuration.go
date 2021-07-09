package main

import (
	"os"
	"strings"

	"github.com/siddontang/go/log"
)

type Configuration struct {
	// Authentication settings
	Homeserver   string
	Username     string
	PasswordFile string
}

func (c *Configuration) GetPassword() (string, error) {
	log.Debug("Reading password from ", c.PasswordFile)
	buf, err := os.ReadFile(c.PasswordFile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buf)), nil
}
