package config

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Type struct {
	AppName    string     `json:"name" validate:"required"`
	Version    string     `json:"version" validate:"required"`
	Prefork    bool       `json:"prefork" default:"false"`
	Host       string     `json:"host" default:"localhost"`
	Port       string     `json:"port" default:"8080"`
	Store      Storage    `json:"store" validate:"required"`
	Monitoring Monitoring `json:"monitoring"`
}

type Storage struct {
	Type         string `json:"type" default:"badger"`        // [badger, redis]
	InMemory     bool   `json:"inMemory" default:"true"`      // for badger
	FileLocation string `json:"fileLocation" default:"./db/"` // for badger
	Password     string `json:"password"`                     // for redis
	Host         string `json:"host" default:"localhost"`     // for redis
	Port         string `json:"port" default:"6379"`          // for redis
	Expiration   int    `json:"expiration" default:"24"`      // for redis
}

type Monitoring struct {
	Id          int64  `json:"id" default:"0"`
	Key         string `json:"key" default:"undefined"`
	Environment string `json:"environment" default:"development"`
}
