package error

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
