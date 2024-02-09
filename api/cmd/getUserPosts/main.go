package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

type Response events.APIGatewayProxyResponse

type User struct {
	UserId string `json:"userId" validate:"required"`
}

type handler struct {
	validator *validator.Validate
	cfig      config.Config
	dao       *db.Dao
}

type posts struct {
	MU   sync.Mutex
	Post []db.Entry
}

func (h *handler) Handler(event events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	userId := event.PathParameters["id"]

	wg := &sync.WaitGroup{}
	posts := posts{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			post, err := h.dao.QueryRecord(userId + "::" + fmt.Sprintf("%v", i))
			if err != nil {
				log.Println(err.Error())
			}
			posts.MU.Lock()
			posts.Post = append(posts.Post, post...)
		}(i)
	}
	wg.Wait()

	body, err := json.Marshal(posts.Post)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "login-handler",
			"Access-Control-Allow-Origin": "*", // Required for CORS support to work
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
