package core

import (
	"chasse-api/internal/logger"
	"fmt"
)

var log = logger.New("ERRORS")

type NotFound struct {
	Message string
}

func (e NotFound) Error() string { return e.Message }

type Internal struct {
	Message string
}

func (e Internal) Error() string { return e.Message }

type BadRequest struct {
	Message string
}

func (e BadRequest) Error() string { return e.Message }

type Info struct {
	Message string
}

func (e Info) Error() string { return e.Message }

type Errors int

const (
	NOT_FOUND Errors = iota
	INTERNAL
	BAD_REQUEST
	INFO
)

func BuildError(t Errors, msg string) error {
	log.Debug(msg)

	switch t {
	case NOT_FOUND:
		return NotFound{Message: msg}
	case INTERNAL:
		return Internal{Message: msg}
	case BAD_REQUEST:
		return BadRequest{Message: msg}
	case INFO:
		return Info{Message: msg}
	default:
		return Internal{Message: msg}
	}
}

func BuildErrorf(t Errors, msg string, vars ...any) error {
	log.Debugf(msg, vars...)

	switch t {
	case NOT_FOUND:
		return NotFound{Message: fmt.Sprintf(msg, vars...)}
	case INTERNAL:
		return Internal{Message: fmt.Sprintf(msg, vars...)}
	case BAD_REQUEST:
		return BadRequest{Message: fmt.Sprintf(msg, vars...)}
	case INFO:
		return Info{Message: fmt.Sprintf(msg, vars...)}
	default:
		return Internal{Message: fmt.Sprintf(msg, vars...)}
	}
}
