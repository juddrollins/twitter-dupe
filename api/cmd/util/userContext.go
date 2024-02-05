package util

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func GetUserContext(event events.APIGatewayProxyRequest) (userId string, err error) {
	lambdaAuthContext, ok := event.RequestContext.Authorizer["lambda"].(map[string]any)
	if !ok {
		return "", errors.New("no auth context")
	}
	userId = lambdaAuthContext["user-id"].(string)
	return userId, nil
}
