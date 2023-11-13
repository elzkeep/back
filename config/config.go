package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type _Database struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	Name             string `yaml:"name"`
	Owner            string `yaml:"owner"`
	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	Type             string `yaml:"type"`
	ConnectionString string `yaml:"connectionString"`
}

type _Mode struct {
	Port         string    `yaml:"port"`
	UploadPath   string    `yaml:"path"`
	DocumentRoot string    `yaml:"documentRoot"`
	Mail         _Mail     `yaml:"mail"`
	Sms          _Sms      `yaml:"sms"`
	Cors         []string  `yaml:"cors"`
	Database     _Database `yaml:"database"`
}

type _Mail struct {
	Sender string `yaml:"sender"`
}

type _Sms struct {
	User   string `yaml:"user"`
	Key    string `yaml:"key"`
	Sender string `yaml:"sender"`
}

type Config struct {
	Version    string `yaml:"version"`
	Develop    _Mode  `yaml:"develop"`
	Production _Mode  `yaml:"production"`
}

var Mail _Mail
var Database _Database
var Sms _Sms
var UploadPath string
var DocumentRoot string
var Version string
var Mode string
var Port string
var Cors []string

func init() {
	buf, err := os.ReadFile(".env.yml")
	if err != nil {
		return
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}

	Mode = os.Getenv("APP_MODE")

	if len(os.Args) == 3 {
		if os.Args[1] == "--mode" {
			Mode = os.Args[2]
		}
	}

	if Mode != "production" {
		Mode = "develop"
	}

	if Mode == "production" {
		Mail = config.Production.Mail
		Sms = config.Production.Sms
		UploadPath = config.Production.UploadPath
		DocumentRoot = config.Production.DocumentRoot
		Port = config.Production.Port
		Database = config.Production.Database
		Cors = config.Production.Cors
	} else {
		Mail = config.Develop.Mail
		Sms = config.Develop.Sms
		UploadPath = config.Develop.UploadPath
		DocumentRoot = config.Develop.DocumentRoot
		Port = config.Develop.Port
		Database = config.Develop.Database
		Cors = config.Develop.Cors
	}

	if Port == "" {
		Port = "80"
	}

	if UploadPath == "" {
		UploadPath = "webdata"
	}

	if Database.Type == "" {
		Database.Type = "mysql"
	}

	if Database.Port == "" {
		Database.Port = "3306"
	}

	if Database.ConnectionString == "" {
		Database.ConnectionString = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", Database.User, Database.Password, Database.Host, Database.Port, Database.Name)
	}

	Version = config.Version

	log.Println("MODE :", Mode)
	log.Println("DATABASE :", Database.ConnectionString)
}
