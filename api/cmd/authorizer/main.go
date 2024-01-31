package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/util"
)

type Statement struct {
	Action   string `json:"Action"`
	Effect   string `json:"Effect"`
	Resource string `json:"Resource"`
}

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

// AuthorizationResponse represents the response format expected by API Gateway Custom Authorizers.
type AuthorizationResponse struct {
	PrincipalID    string         `json:"principalId"`
	PolicyDocument PolicyDocument `json:"policyDocument"`
}

// Handler is the Lambda authorizer function handler.
func Handler(request events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	//token := strings.Split(request.Headers["Authorization"], " ")[1]
	token := strings.Split(request.Headers["authorization"], " ")[1]

	parsedToken, err := util.ParseJWT(token)
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf(err.Error())
	}

	tokenValidationError := parsedToken.Valid()
	if tokenValidationError != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf(tokenValidationError.Error())
	}

	enrichedContext := map[string]any{
		"user-id": parsedToken.User.SK,
	}

	if tokenValidationError == nil {
		log.Println("the token is valid")
		// If the token is valid, allow access.
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: parsedToken.User.SK, // Change this to the authenticated user's ID or username.
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Allow",
						Resource: []string{request.RouteArn},
					},
				},
			},
			Context: enrichedContext,
		}, nil
	}

	// If the token is not valid or the user doesn't have the necessary permissions, deny access.
	return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("Unauthorized")
}

func main() {
	lambda.Start(Handler)
}
