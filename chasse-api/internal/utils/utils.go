package utils

import (
	e "chasse-api/internal/error"
	"encoding/json"
)

func Encode[T any](data T) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, e.BuildErrorf(e.INTERNAL, "failed while marshal object: %s", err.Error())
	}
	return result, nil
}

func Decode[T any](s []byte, t T) error {
	if err := json.Unmarshal(s, t); err != nil {
		return e.BuildErrorf(e.INTERNAL, "failed while unmarshal object: %s", err.Error())
	}
	return nil
}

func MapToStruct[T any](data any) (*T, error) {
	out := new(T)
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(j, out)

	return out, nil
}
