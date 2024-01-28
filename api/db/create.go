package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (dao Dao) CreateRecord(entry Entry) error {
	//entry.PK = uuid.New().String()
	entityParsed, err := attributevalue.MarshalMap(entry)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(dao.db.TableName),
	}

	_, err = dao.db.Client.PutItem(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}
