package lambdahandlers

import (
	"github.com/aws/aws-lambda-go/events"
)

type LambdaVideosHandler struct {
	LambdaHandler
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
