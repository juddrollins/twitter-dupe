package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
)

type Response events.APIGatewayProxyResponse

type CTX struct {
	cfig config.Config
	dao  *db.Dao
}

func Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	var cfig = config.New()
	var ctx = &CTX{
		cfig: cfig,
		dao:  db.InitDb(&cfig.Ddb),
	}

	log.Println(event.PathParameters["id"])
	log.Println(event.Path)

	var user, err = ctx.dao.GetRecord("juddrollins", "")
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
	lambda.Start(Handler)
}
