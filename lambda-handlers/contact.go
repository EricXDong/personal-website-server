package lambdahandlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type LambdaContactHandler struct {
	LambdaHandler
}

func (lch *LambdaContactHandler) Handle(email string, message string) (events.APIGatewayProxyResponse, error) {
	sesh, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	svc := ses.New(sesh)

	//	Wow this is terrible to look at
	emailData := &ses.SendEmailInput{
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(message),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Contact from " + email),
			},
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(lch.Env.ContactEmail)},
		},
		Source: aws.String(lch.Env.ContactEmail),
	}

	//	Error sending email
	_, err := svc.SendEmail(emailData)
	if err != nil {
		return CreateLambdaErrorResponse(LambdaErrorResponse{
			StatusCode:  500,
			Error:       err.Error(),
			Description: "Error sending contact message",
		}), nil
	}

	//	Success
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
