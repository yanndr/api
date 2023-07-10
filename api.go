package api

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

// Information represents the information about the API.
type Information struct {
	APIVersion string `json:"api_version"`
}

type Serializable interface {
	Serialize() (string, error)
}

func Serialize(data interface{}) (string, error) {
	writer := md5.New()
	e := gob.NewEncoder(writer)
	err := e.Encode(data)
	if err != nil {
		return "", fmt.Errorf(`failed gob Encode :%w`, err)
	}
	return fmt.Sprintf("%x", writer.Sum(nil)), nil
}
