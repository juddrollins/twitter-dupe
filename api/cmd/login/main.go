package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/cmd/util"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

type Response events.APIGatewayProxyResponse

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	log.Print(event)
	log.Println(event.RequestContext.ResourceID)
	log.Println(event.Resource)

	v := validator.New()
	var cfig = config.New()
	var ctx = util.CTX{
		Cfig: cfig,
		Dao:  db.InitDb(&cfig.Ddb),
	}

	var input LoginUser
	json.Unmarshal([]byte(event.Body), &input)

	// Validate User Input to match RegisterUser struct
	validationError := v.Struct(input)
	if validationError != nil {
		for _, e := range validationError.(validator.ValidationErrors) {
			if e != nil {
				log.Println(e)
				return Response{StatusCode: 400}, errors.New(e.Field())
			}
		}
	}

	var user, err = ctx.Dao.QueryRecord(input.Username)
	if err != nil {
		log.Println(err.Error())
	}

	var user_data = strings.Split(user[0].Data, "::")
	var user_password = user_data[1]

	// Check if password matches
	if !util.CheckPasswordHash(input.Password, user_password) {
		log.Println("Password does not match")
		return Response{StatusCode: 400}, errors.New("password does not match")
	}

	// Create JWT ?? //TODO
	var success = map[string]string{
		"success": "true",
	}

	body, err := json.Marshal(success)
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
	log.Println("Login Lambda Started")
	lambda.Start(Handler)
}
