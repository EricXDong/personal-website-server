package lambdahandlers

import (
	"personal-website-server/env"
)

type LambdaContactHandler struct {
	Env *env.Env
}

// func (lch *LambdaContactHandler) Handle(email string, message string) (string, error) {
// 	client := ses.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

// 	return email + message, nil
// }
