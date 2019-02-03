package lambdahandlers

import (
	"personal-website-server/env"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaVideosRequest struct {
	Password string `json:"password"`
}

type LambdaVideosHandler struct {
	Env *env.Env
}

func (lvh *LambdaVideosHandler) Handle(password string) (events.APIGatewayProxyResponse, error) {
	if password == lvh.Env.VideosPassword {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
	}, nil
}
