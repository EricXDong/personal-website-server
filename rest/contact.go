package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"personal-website-server/env"
)

type ContactHandler struct {
	RestHandler
}

type ContactRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type emailParameters struct {
	From    string
	To      string
	Subject string
	Message string
}

func (ch *ContactHandler) receiveContact(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println(request.Message)

	//	Generate email from request
	buffer := new(bytes.Buffer)
	template := template.Must(template.New("email").Parse(generateEmailTemplate()))
	template.Execute(buffer, &emailParameters{
		From:    request.Email,
		To:      ch.env.ContactEmail,
		Subject: "Contact from " + request.Email,
		Message: request.Message,
	})
	template.Execute(os.Stdout, &emailParameters{
		From:    request.Email,
		To:      ch.env.ContactEmail,
		Subject: "Contact from " + request.Email,
		Message: request.Message,
	})

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", ch.env.EmailUsername, ch.env.EmailPassword, "smtp.gmail.com"),
		ch.env.EmailUsername,
		[]string{ch.env.ContactEmail},
		buffer.Bytes(),
		// []byte(request.Message),
	)
	if err != nil {
		respondWithError(&ErrorMessage{
			Error:       err,
			Description: "Error sending contact message",
		}, http.StatusInternalServerError, &w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateEmailTemplate() string {
	return `
		From: {{.From}}<br />
		To: {{.To}}<br />
		Subject: {{.Subject}}<br />
		MIME-version: 1.0<br />
		Content-Type: text/html; charset=&quot;UTF-8&quot;<br />
		<br />
		{{.Message}}
	`
}

func newContactHandler(env *env.Env) *ContactHandler {
	return &ContactHandler{
		RestHandler: RestHandler{
			env,
		},
	}
}
