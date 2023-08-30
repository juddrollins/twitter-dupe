package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type CTX struct {
	cfig config.Config
	dao  *db.Dao
}

func Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	var entry db.Entry
	json.Unmarshal([]byte(event.Body), &entry)

	var cfig = config.New()
	var ctx = &CTX{
		cfig: cfig,
		dao:  db.InitDb(&cfig.Ddb),
	}

	var testValue, err = ctx.dao.CreateRecord(entry)
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := json.Marshal(testValue)
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
