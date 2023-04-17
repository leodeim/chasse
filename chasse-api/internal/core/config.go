package core

import (
	"chasse-api/internal/logger"
	"encoding/json"

	"github.com/leonidasdeim/goconfig"
	fh "github.com/leonidasdeim/goconfig/pkg/filehandler"
)

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Configuration struct {
	AppName  string          `json:"name" validate:"required"`
	Version  string          `json:"version" validate:"required"`
	Prefork  bool            `json:"prefork" default:"false"`
	Host     string          `json:"host" default:"localhost"`
	Port     string          `json:"port" default:"8080"`
	Store    StoreConfig     `json:"store" validate:"required"`
	Monitor  MonitorConfig   `json:"monitoring"`
	LogLevel logger.LogLevel `json:"loglevel" default:"INFO"`
}

type StoreConfig struct {
	Type       string `json:"type" default:"badger"`        // [badger, redis]
	InMemory   bool   `json:"inMemory"`                     // for badger
	Location   string `json:"fileLocation" default:"./db/"` // for badger
	Password   string `json:"password"`                     // for redis
	Host       string `json:"host" default:"localhost"`     // for redis
	Port       string `json:"port" default:"6379"`          // for redis
	Expiration int    `json:"expiration" default:"24"`      // for redis
}

type MonitorConfig struct {
	Id  int64  `json:"id" default:"0"`
	Key string `json:"key" default:"undefined"`
	Env string `json:"environment" default:"development"`
}

type Callback func(Configuration) error

type Config struct {
	config    *goconfig.Config[Configuration]
	callbacks map[int]Callback
	quit      chan bool
}

func InitConfig() *Config {
	h, _ := fh.New(fh.WithName("chasse"), fh.WithType(fh.JSON))
	config, err := goconfig.Init[Configuration](h)
	if err != nil {
		log.Errorf("InitConfig error: %s", err.Error())
	} else {
		print(config.GetCfg())
	}

	c := &Config{
		config:    config,
		callbacks: make(map[int]Callback),
		quit:      make(chan bool, 1),
	}

	return c
}

func (c *Config) Get() Configuration {
	return c.config.GetCfg()
}

func (c *Config) GetTimestamp() string {
	return c.config.GetTimestamp()
}

func (c *Config) Close() {
	if len(c.quit) < 1 {
		c.quit <- true
	}
}

func print(a any) {
	data, _ := json.MarshalIndent(a, "", "  ")
	log.Info(string(data))
}
