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
	Sender   string `yaml:"sender"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
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
	config := &Config{}
	if buf, err := os.ReadFile(".env.yml"); err == nil {
		if err := yaml.Unmarshal(buf, config); err != nil {
			log.Fatalf("Unmarshal: %v", err)
			return
		}
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

	applyEnv()

	if Database.ConnectionString == "" {
		Database.ConnectionString = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", Database.User, Database.Password, Database.Host, Database.Port, Database.Name)
	}

	Version = config.Version

	log.Println("MODE :", Mode)
	// 접속 문자열엔 비밀번호가 들어 있어 남기지 않는다
	log.Printf("DATABASE : %v@%v:%v/%v", Database.User, Database.Host, Database.Port, Database.Name)
}

// 비밀값은 .env.yml(저장소에 올리지 않음) 대신 환경변수로도 줄 수 있다 — 환경변수가 우선
func applyEnv() {
	set := func(target *string, key string) {
		if v := os.Getenv(key); v != "" {
			*target = v
		}
	}

	set(&Database.Host, "DB_HOST")
	set(&Database.Port, "DB_PORT")
	set(&Database.Name, "DB_NAME")
	set(&Database.User, "DB_USER")
	set(&Database.Password, "DB_PASSWORD")
	set(&Sms.User, "SMS_USER")
	set(&Sms.Key, "SMS_KEY")
	set(&Sms.Sender, "SMS_SENDER")
	set(&Mail.Sender, "MAIL_SENDER")
	set(&Mail.User, "MAIL_USER")
	set(&Mail.Password, "MAIL_PASSWORD")
}
