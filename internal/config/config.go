package config

import (
	"log"
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
			CSRFTimeout time.Duration `yaml:"csrf_timeout"`
		} `yaml:"timeout"`
		Static struct {
			Dir    string `yaml:"dir"`
			Handle string `yaml:"handle"`
		} `yaml:"static"`
		Keys struct {
			CSRFAuthKey string `yaml:"csrf_authkey"`
			AuthKey     string `yaml:"auth_key"`
			EncKey      string `yaml:"enc_key"`
		} `yaml:"keys"`
		Services struct {
			Auth struct {
				Address	string `yaml:"address"`
				Port string `yaml:"port"`
			} `yaml:"auth"`
			Profile struct {
				Address	string `yaml:"address"`
				Port string `yaml:"port"`
			} `yaml:"profile"`
			MailBox struct {
				Address	string `yaml:"address"`
				Port string `yaml:"port"`
			} `yaml:"mailbox"`
			Database struct {
				Address	string `yaml:"address"`
				Port string `yaml:"port"`
			} `yaml:"database"`
		}
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
	config, err := NewConfig("configs/test.yml")
	if err != nil {
		log.Fatal(err)
	}
	return config
}
