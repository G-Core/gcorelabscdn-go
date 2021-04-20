package gcore

import "encoding/json"

type ErrorResponse struct {
	Errors  *json.RawMessage
	Message string
}

func (resp *ErrorResponse) Error() string {
	if resp.Message != "" {
		return resp.Message
	}

	b, err := json.MarshalIndent(resp.Errors, "", "\t")
	if err != nil {

	}

	return string(b)
}
