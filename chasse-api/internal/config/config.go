package config

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Type struct {
	AppName string     `json:"name" validate:"required"`
	Version string     `json:"version" validate:"required"`
	Port    string     `json:"port" validate:"required"`
	Store   Connection `json:"store" validate:"required"`
}

type Connection struct {
	User     string `json:"user"`
	Password string `json:"password" validate:"required"`
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
}
