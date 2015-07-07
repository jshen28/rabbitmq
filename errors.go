package rabbitmq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	if res.Body == nil {
		return fmt.Errorf("rabbitmq: Error %d (%s)", res.StatusCode, http.StatusText(res.StatusCode))
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("rabbitmq: Error %d (%s) when reading body: %v", res.StatusCode, http.StatusText(res.StatusCode), err)
	}
	errReply := new(Error)
	err = json.Unmarshal(slurp, errReply)
	if err == nil && errReply != nil {
		if errReply.StatusCode == 0 {
			errReply.StatusCode = res.StatusCode
		}
		errReply.Message = strings.TrimSpace(errReply.Message)
		errReply.Reason = strings.TrimSpace(errReply.Reason)
		return errReply
	}
	return fmt.Errorf("rabbitmq: Error %d (%s)", res.StatusCode, http.StatusText(res.StatusCode))
}

type Error struct {
	StatusCode int    `json:"status"`
	Message    string `json:"error"`
	Reason     string `json:"reason"`
}

func (e *Error) Error() string {
	if e.Message != "" {
		if e.Reason != "" {
			return fmt.Sprintf("rabbitmq: Error %d (%s): %s (%s)", e.StatusCode, http.StatusText(e.StatusCode), e.Message, e.Reason)
		}
		return fmt.Sprintf("rabbitmq: Error %d (%s): %s", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
	} else {
		return fmt.Sprintf("rabbitmq: Error %d (%s)", e.StatusCode, http.StatusText(e.StatusCode))
	}
}
