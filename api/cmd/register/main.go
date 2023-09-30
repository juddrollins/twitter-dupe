package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"

	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/cmd/util"
	"github.com/juddrollins/twitter-dupe/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type RegisterUser struct {
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

	var input RegisterUser
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

	// Hash password
	password, hashErr := util.HashPassword(input.Password)
	if hashErr != nil {
		log.Println(hashErr.Error())
	}
	uuiid_username := uuid.New().String()

	// Create entry
	entry := db.Entry{
		PK:   input.Username,
		SK:   uuiid_username,
		Data: uuiid_username + "::" + password,
	}

	var testValue, err = h.dao.CreateRecord(entry)
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

	cfig := config.New()
	var h = handler{
		validator: validator.New(),
		cfig:      cfig,
		dao:       db.InitDb(&cfig.Ddb),
	}

	lambda.Start(h.Handler)

}
