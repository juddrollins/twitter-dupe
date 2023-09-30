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

type handler struct {
	validator *validator.Validate
	cfig      config.Config
	dao       *db.Dao
}

func (h *handler) Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	var input LoginUser
	json.Unmarshal([]byte(event.Body), &input)

	// Validate User Input to match RegisterUser struct
	validationError := h.validator.Struct(input)
	if validationError != nil {
		for _, e := range validationError.(validator.ValidationErrors) {
			if e != nil {
				log.Println(e)
				return Response{StatusCode: 400}, errors.New(e.Field())
			}
		}
	}

	var user, err = h.dao.QueryRecord(input.Username)
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
	var jwt, jwt_err = util.GenerateJWT(user[0].PK)
	if jwt_err != nil {
		log.Println(err.Error())
	}

	var token_object = map[string]string{
		"token": jwt,
	}

	body, err := json.Marshal(token_object)
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

	cfig := config.New()
	var h = handler{
		validator: validator.New(),
		cfig:      cfig,
		dao:       db.InitDb(&cfig.Ddb),
	}

	lambda.Start(h.Handler)
}
