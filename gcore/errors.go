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

	b, _ := json.MarshalIndent(resp.Errors, "", "\t")
	return string(b)
}
