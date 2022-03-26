package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Параметры валидации

var SigninKeys = []string{
	"username",
	"password",
}

var SignupKeys = []string{
	"first_name",
	"last_name",
	"username",
	"password",
	"password_confirmation",
}

var MaxFieldLength int = 20

// Общая конфигурация

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
		Static struct {
			Dir    string `yaml:"dir"`
			Handle string `yaml:"handle"`
		} `yaml:"static"`
		Keys struct {
			CSRFToken string `yaml:"csrf_token"`
			AuthKey   string `yaml:"auth_key"`
			EncKey    string `yaml:"enc_key"`
		} `yaml:"keys"`
	} `yaml:"server"`
	Database struct {
		Type     string `yaml:"type"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"dbname"`
	} `yaml:"database"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	var file *os.File
	file, err := os.Open(configPath)
	if err != nil {
		file, err = os.Open(filepath.Join("../..", configPath))
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func TestConfig() *Config {
	config := &Config{
		Server: struct {
			Port    string "yaml:\"port\""
			Timeout struct {
				Server time.Duration "yaml:\"server\""
				Write  time.Duration "yaml:\"write\""
				Read   time.Duration "yaml:\"read\""
				Idle   time.Duration "yaml:\"idle\""
			} "yaml:\"timeout\""
			Static struct {
				Dir    string "yaml:\"dir\""
				Handle string "yaml:\"handle\""
			} "yaml:\"static\""
			Keys struct {
				CSRFToken string "yaml:\"csrf_token\""
				AuthKey   string "yaml:\"auth_key\""
				EncKey    string "yaml:\"enc_key\""
			} "yaml:\"keys\""
		}{
			Port: "8080",
			Timeout: struct {
				Server time.Duration "yaml:\"server\""
				Write  time.Duration "yaml:\"write\""
				Read   time.Duration "yaml:\"read\""
				Idle   time.Duration "yaml:\"idle\""
			}{
				Server: 10 * time.Second,
				Write:  5 * time.Second,
				Read:   5 * time.Second,
				Idle:   5 * time.Second,
			},
			Static: struct {
				Dir    string "yaml:\"dir\""
				Handle string "yaml:\"handle\""
			}{
				Dir:    "static",
				Handle: "/static",
			},
			Keys: struct{CSRFToken string "yaml:\"csrf_token\""; AuthKey string "yaml:\"auth_key\""; EncKey string "yaml:\"enc_key\""}{
				CSRFToken: "",
				AuthKey: "aPdSgVkYp3s6v9y$B&E)H+MbQeThWmZq",
				EncKey: "cRfUjXn2r5u8x/A?D*G-KaPdSgVkYp3s",
			},
		},
		Database: struct {
			Type     string "yaml:\"type\""
			User     string "yaml:\"user\""
			Password string "yaml:\"password\""
			Host     string "yaml:\"host\""
			Port     int    "yaml:\"port\""
			Name     string "yaml:\"dbname\""
		}{
			Type: "mock",
		},
	}
	return config
}
