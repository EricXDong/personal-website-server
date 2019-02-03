package main

import (
	"encoding/json"
	"personal-website-server/env"
	lambdahandlers "personal-website-server/lambda-handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MessageType string

const (
	Videos  MessageType = "videos"
	Contact MessageType = "contact"
)

type Request struct {
	Type MessageType            `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func handleLambdaRequest(env *env.Env) {
	lambda.Start(func(request events.APIGatewayProxyRequest) (evt events.APIGatewayProxyResponse, problem error) {
		//	Handles unexpected errors
		defer func() {
			if r := recover(); r != nil {
				evt.StatusCode = 500
				err, _ := r.(error)
				evt.Body = err.Error()
			}
		}()

		var message Request
		err := json.Unmarshal([]byte(request.Body), &message)
		if err != nil {
			evt.StatusCode = 500
			evt.Body = err.Error()
		}

		//	Call appropriate handler based on type argument
		switch message.Type {
		case Videos:
			handler := lambdahandlers.LambdaVideosHandler{
				Env: env,
			}
			return handler.Handle(message.Data["password"].(string))
		default:
			evt.StatusCode = 500
			evt.Body = "Invalid event type"
		}

		return evt, problem
	})
}
