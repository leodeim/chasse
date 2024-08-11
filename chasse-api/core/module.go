package core

type Module interface {
	Run(func())
	Close()
}
