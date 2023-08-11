package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/juddrollins/twitter-dupe/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/aws/aws-lambda-go/lambda"
)

// TODO best way to develop locally with lambda?

type Response events.APIGatewayProxyResponse

func Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	var entry db.Entry
	json.Unmarshal([]byte(event.Body), &entry)

	var dbConnection = db.InitDatabase()
	var testValue, err = dbConnection.CreateRecord(entry)
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
	lambda.Start(Handler)
	// var dbConnection = db.InitDatabase()
	// var testValue, err = dbConnection.CreateRecord(db.Entry{PK: "test", SK: "sortIt", Data: "test"})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(testValue)

	// var testValue, err = dbConnection.GetRecord("it", "sortIt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(testValue.Data)

}
