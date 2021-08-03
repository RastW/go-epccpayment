package pack

import (
	"encoding/json"
	"go-epccpayment/common/log"
	"io"
)

func Unpack(body io.Reader) (interface{}, error) {
	dec := json.NewDecoder(body)
	m := make(map[string]string, 0)
	err := dec.Decode(m)
	if err != nil {
		return nil, err
	}
	log.Info("1111111111")
	return nil, nil
}
