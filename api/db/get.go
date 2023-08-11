package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (db Database) GetRecord(PK string, SK string) (Entry, error) {

	// input := &dynamodb.GetItemInput{TableName: aws.String(db.tablename),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"PK": {
	// 			S: &PK,
	// 		},
	// 	}}

	result, err := db.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.tablename),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(PK),
			},
			"SK": {
				S: aws.String(SK),
			},
		}})
	if err != nil {
		return Entry{}, err
	}

	var entry Entry
	err = dynamodbattribute.UnmarshalMap(result.Item, &entry)
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}
