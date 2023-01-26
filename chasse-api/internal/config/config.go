package config

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Type struct {
	AppName    string     `json:"name" validate:"required"`
	Version    string     `json:"version" validate:"required"`
	Host       string     `json:"host" default:"localhost"`
	Port       string     `json:"port" default:"8080"`
	Store      Connection `json:"store" validate:"required"`
	Monitoring Monitoring `json:"monitoring"`
}

type Connection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host" default:"localhost"`
	Port     string `json:"port" validate:"required"`
}

type Monitoring struct {
	Id          int64  `json:"id" default:"0"`
	Key         string `json:"key" default:"undefined"`
	Environment string `json:"environment" default:"development"`
}
