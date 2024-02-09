package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

// Test the GetUserPosts handler function
func TestHandler(t *testing.T) {
	cfig := config.New()
	var h = handler{
		validator: validator.New(),
		cfig:      cfig,
		dao:       db.InitDb(&cfig.Ddb),
	}

	//Set os ENV to local
	t.Setenv("ENV", "local")

	// Create a new request
	req := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "judd",
		},
	}

	h.Handler(req)

}
