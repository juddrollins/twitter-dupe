package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database struct {
	Client    *dynamodb.DynamoDB
	tablename string
}

type Entry struct {
	PK   string `json:"PK"`
	SK   string `json:"SK"`
	Data string `json:"Data"`
}

type EntryService interface {
	CreateRecord(m Entry) (Entry, error)
	GetRecord(PK string, SK string) (Entry, error)
}

func InitDatabase() EntryService {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Database{
		Client:    dynamodb.New(sess),
		tablename: "twitter-table-dev",
	}
}
