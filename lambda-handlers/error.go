package lambdahandlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaErrorResponse struct {
	StatusCode  int
	Error       string
	Description string
}

func CreateLambdaErrorResponse(response LambdaErrorResponse) events.APIGatewayProxyResponse {
	message, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: response.StatusCode,
		Body:       string(message),
	}
}
