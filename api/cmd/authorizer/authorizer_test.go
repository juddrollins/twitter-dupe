package main

//Write test cases for the Authorizer function
import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juddrollins/twitter-dupe/cmd/util"
	"github.com/juddrollins/twitter-dupe/db"
)

func TestHandler(t *testing.T) {

	// Create a new db.user entry
	userEntry := db.Entry{
		PK:   "user::" + string("randomUser"),
		SK:   "",
		Data: "::",
	}

	jwtToken, _ := util.GenerateJWT(userEntry)

	//create a new request with the token
	req := events.APIGatewayV2CustomAuthorizerV2Request{
		Headers: map[string]string{
			"authorization": "Bearer " + jwtToken,
		},
	}

	// Call the Handler function
	_, err := Handler(req)

	// Check if the error is not nil
	if err != nil {
		t.Errorf("expected nil error, but got %v", err)
	}
}
