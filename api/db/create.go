package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (db Database) CreateRecord(entry Entry) (Entry, error) {
	//entry.PK = uuid.New().String()
	entityParsed, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		return Entry{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.tablename),
	}

	_, err = db.Client.PutItem(input)
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}
