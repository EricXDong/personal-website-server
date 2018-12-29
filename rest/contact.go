package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"personal-website-server/env"
)

type ContactHandler struct {
	RestHandler
}

type ContactRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (ch ContactHandler) receiveContact(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(createHTTPBodyError(err), http.StatusBadRequest, &w)
		return
	}

	request := ContactRequest{}
	err = json.Unmarshal(data, &request)
	if err != nil {
		respondWithError(createJSONParseError(err), http.StatusBadRequest, &w)
		return
	}

	//	Generate email from request
	email := []byte("To: " + ch.env.ContactEmail + "\r\n" +
		"Subject: Contact from " + request.Email + "\r\n\r\n" +
		request.Message + "\r\n")
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		ch.env.EmailAuth,
		ch.env.EmailUsername,
		[]string{ch.env.ContactEmail},
		email,
	)
	if err != nil {
		respondWithError(&ErrorMessage{
			Error:       err.Error(),
			Description: "Error sending contact message",
		}, http.StatusInternalServerError, &w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func newContactHandler(env *env.Env) *ContactHandler {
	return &ContactHandler{
		RestHandler: RestHandler{
			env,
		},
	}
}
