package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	AppName = "APP_NAME"

	serverPort         = "SERVER_PORT"
	envShutdownTimeout = "SHUTDOWN_TIMEOUT"

	parserShutdownTimeoutError = "config: parse server shutdown timeout error"
)

type AppConf struct {
	AppName string
	Server  Server
	DB      DB
}

type DB struct {
	Driver   string
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

type Server struct {
	Port            string
	ShutdownTimeout time.Duration
}

func NewAppConf() AppConf {
	port := os.Getenv(serverPort)

	return AppConf{
		AppName: os.Getenv(AppName),
		Server: Server{
			Port: port,
		},
		DB: DB{
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}

func (a *AppConf) Init() {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv(envShutdownTimeout))
	if err != nil {
		log.Fatal(parserShutdownTimeoutError)
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second

	a.Server.ShutdownTimeout = shutDownTimeout
}
