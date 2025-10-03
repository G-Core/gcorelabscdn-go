package gcore

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	StatusCode int
	Errors     *json.RawMessage
	Message    string
}

func (resp *ErrorResponse) Error() string {
	if resp.Message != "" {
		return fmt.Sprintf("status: %d, message: %s", resp.StatusCode, resp.Message)
	}

	if resp.Errors != nil {
		b, _ := json.MarshalIndent(resp.Errors, "", "\t")
		return fmt.Sprintf("status: %d, errors: %s", resp.StatusCode, string(b))
	}

	return fmt.Sprintf("status: %d", resp.StatusCode)
}
