package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Параметры валидации

var SigninKeys = []string {
	"email",
	"password",
}

var SignupKeys = []string {
	"first_name",
	"last_name",
	"email",
	"password",
	"password_confirmation",
}

var MaxFieldLength int = 20

// Общая конфигурация

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Timeout struct {
			Server	time.Duration `ytml:"server"`
			Write	time.Duration `yaml:"write"`
			Read	time.Duration `yaml:"read"`
			Idle	time.Duration `yaml:"idle"`
		}	`yaml:"timeout"`
	}	`yaml:"server"`
	Database struct {
		User		string `yaml:"user"`
		Password	string `yaml:"password"`
		Host		string `yaml:"host"`
		Port		int `yaml:"port"`
		Name		string `yaml:"dbname"`
	}	`yaml:"database"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}