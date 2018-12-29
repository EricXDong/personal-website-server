package rest

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

func respondWithError(em *ErrorMessage, status int, w *http.ResponseWriter) {
	if response, err := json.Marshal(&em); err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
	} else {
		(*w).WriteHeader(status)
		(*w).Write(response)
	}
}

func createHTTPBodyError(err error) *ErrorMessage {
	return &ErrorMessage{
		Error:       err.Error(),
		Description: "Error reading request HTTP body",
	}
}

func createJSONParseError(err error) *ErrorMessage {
	return &ErrorMessage{
		Error:       err.Error(),
		Description: "Error parsing request HTTP body to JSON",
	}
}
