package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

type Response events.APIGatewayProxyResponse

type Post struct {
	Content string `json:"post" validate:"required,min=8"`
}

type handler struct {
	validator *validator.Validate
	cfig      config.Config
	dao       *db.Dao
}

func (h *handler) Handler(con context.Context, event events.APIGatewayProxyRequest) (Response, error) {

	var input Post
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
	var buf bytes.Buffer

	randomNumber := rand.Intn(10) + 1

	// // Create a new db entry for a post
	// entry := db.Entry{
	// 	PK:   "post::" + string(randomNumber),
	// 	SK:   "",
	// 	Data: "::",
	// }

	// var user, err = h.dao.CreateRecord()
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// body, err := json.Marshal(user)
	// if err != nil {
	// 	return Response{StatusCode: 404}, err
	// }
	// json.HTMLEscape(&buf, body)

	// resp := Response{
	// 	StatusCode:      200,
	// 	IsBase64Encoded: false,
	// 	Body:            buf.String(),
	// 	Headers: map[string]string{
	// 		"Content-Type":           "application/json",
	// 		"X-MyCompany-Func-Reply": "register-handler",
	// 	},
	// }

	return resp, nil
}

func main() {
	log.Println("Post Lambda Started")
	cfig := config.New()
	var h = handler{
		validator: validator.New(),
		cfig:      cfig,
		dao:       db.InitDb(&cfig.Ddb),
	}
	lambda.Start(h.Handler)
}
