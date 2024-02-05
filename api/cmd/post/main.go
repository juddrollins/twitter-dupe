package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/cmd/util"
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

	var post Post
	json.Unmarshal([]byte(event.Body), &post)

	// Validate User Input to match Content struct
	validationError := h.validator.Struct(post)
	if validationError != nil {
		for _, e := range validationError.(validator.ValidationErrors) {
			if e != nil {
				log.Println(e)
				return Response{StatusCode: 400}, errors.New(e.Field())
			}
		}
	}

	userId, err := util.GetUserContext(event)
	if err != nil {
		return Response{StatusCode: 500, Body: "no auth context"}, nil
	}

	randomNumber := rand.Intn(10) + 1

	// Create a new db entry for a post
	entry := db.Entry{
		PK:        "post::" + fmt.Sprintf("%v", randomNumber),
		SK:        userId + "::" + uuid.NewString(),
		Data:      post.Content,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	err = h.dao.CreateRecord(entry)
	if err != nil {
		log.Println(err.Error())
	}

	resp := Response{
		StatusCode:      201,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "register-handler",
		},
	}

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
