package internal

import "encoding/json"

type ApiErr struct {
	ErrMessage []string `json:"errors"`
}

var unexpectedError, _ = json.Marshal(&ApiErr{ErrMessage: []string{"something went wrong"}})

func wrapApiErr(errMessages ...string) string {
	b, err := json.Marshal(&ApiErr{ErrMessage: errMessages})
	if err != nil {
		return string(unexpectedError)
	}
	return string(b)
}
