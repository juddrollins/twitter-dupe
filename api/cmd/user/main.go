package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

type Response events.APIGatewayProxyResponse

type handler struct {
	validator *validator.Validate
	cfig      config.Config
	dao       *db.Dao
}

func (h *handler) Handler(con context.Context, event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	lambdaAuthContext, ok := event.RequestContext.Authorizer["lambda"].(map[string]any)
	if !ok {
		return Response{StatusCode: 500, Body: "no auth context"}, nil
	}

	log.Println(event.RequestContext.Authorizer["lambda"].(map[string]any))

	//Put this user id that has been validate by the authorizer as the post author
	log.Println(lambdaAuthContext["user-id"])

	log.Println("Id path param: " + event.PathParameters["id"])

	var user, err = h.dao.QueryRecord(event.PathParameters["id"])
	if err != nil {
		log.Println(err.Error())
	}

	body, err := json.Marshal(user)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "register-handler",
		},
	}

	return resp, nil
}

func main() {
	log.Println("Register Lambda Started")
	cfig := config.New()
	var h = handler{
		validator: validator.New(),
		cfig:      cfig,
		dao:       db.InitDb(&cfig.Ddb),
	}
	lambda.Start(h.Handler)
}
