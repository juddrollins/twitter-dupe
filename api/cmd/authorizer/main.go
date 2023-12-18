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
func Handler(request events.APIGatewayV2CustomAuthorizerV2Request) (AuthorizationResponse, error) {
	//token := strings.Split(request.Headers["Authorization"], " ")[1]
	token := strings.Split(request.Headers["authorization"], " ")[1]

	parsedToken, err := util.ParseJWT(token)
	if err != nil {
		return AuthorizationResponse{}, fmt.Errorf(err.Error())
	}
	log.Println(parsedToken)

	// Your token validation logic goes here.
	// Check if the token is valid and if the user has the necessary permissions.

	if parsedToken.Valid() != nil {
		// If the token is valid, allow access.
		return AuthorizationResponse{
			PrincipalID: "user123", // Change this to the authenticated user's ID or username.
			PolicyDocument: PolicyDocument{
				Version: "2012-10-17",
				Statement: []Statement{
					{
						Action:   "execute-api:Invoke",
						Effect:   "Allow",
						Resource: request.RouteArn,
					},
				},
			},
		}, nil
	}

	// If the token is not valid or the user doesn't have the necessary permissions, deny access.
	return AuthorizationResponse{}, fmt.Errorf("Unauthorized")
}

func main() {
	lambda.Start(Handler)
}
