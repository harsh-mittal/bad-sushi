package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

const (
	tcpFileName = "tcp.toml"
)

type tcpConfig struct {
	Host           string `toml: "host"`
	Port           int    `toml: "port"`
	MaxConnections int    `toml: "max-connections"`
	LogLevel       string `toml: "log-level"`
}

func (tc *tcpConfig) Serialize() string {
	return fmt.Sprintf("%s:%d", tc.Host, tc.Port)
}

func LoadTcpConfig() (*tcpConfig, error) {
	config := &tcpConfig{}
	filePath := fmt.Sprintf("%s", tcpFileName)

	err := LoadToml(filePath, config)
	if err != nil {
		return nil, err
	}
	if config.LogLevel != "debug" {
		return nil, errors.New("invalid log level")
	}

	return config, nil
}

func LoadToml(filePath string, config *tcpConfig) error {
	err := FileExists(filePath)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		return errors.New("file access error")
	}
	if info.IsDir() {
		return errors.New("not a file, it is a directory")
	}

	return nil
}
