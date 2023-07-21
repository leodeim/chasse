package core

import (
	"chasse-api/internal/logger"
	"encoding/json"

	"github.com/leonidasdeim/cog"
	fh "github.com/leonidasdeim/cog/pkg/filehandler"
)

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Configuration struct {
	AppName  string          `json:"name" default:"chasse"`
	Version  string          `json:"version" default:"dev"`
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
	cog       *cog.Config[Configuration]
	callbacks map[int]Callback
	quit      chan bool
}

func InitConfig() *Config {
	h, _ := fh.New(fh.WithName("chasse"), fh.WithType(fh.JSON))
	c, err := cog.Init[Configuration](h)
	if err != nil {
		log.Fatalf("InitConfig error: %s", err.Error())
	}
	print(c.GetCfg())

	return &Config{
		cog:       c,
		callbacks: make(map[int]Callback),
		quit:      make(chan bool, 1),
	}
}

func (c *Config) Get() Configuration {
	return c.cog.GetCfg()
}

func (c *Config) GetTimestamp() string {
	return c.cog.GetTimestamp()
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
