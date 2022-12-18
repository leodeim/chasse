package config

type Subscriber int

const (
	STORE Subscriber = iota
	NUMBER_OF_SUBS
)

type Type struct {
	AppName string     `json:"name"`
	Version string     `json:"version"`
	Port    string     `json:"port"`
	Store   Connection `json:"store"`
}

type Connection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
