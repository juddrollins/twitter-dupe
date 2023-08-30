package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (dao Dao) GetRecord(PK string, SK string) (Entry, error) {

	result, err := dao.db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(dao.db.TableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: PK,
			},
			"SK": &types.AttributeValueMemberS{
				Value: SK,
			},
		}})
	if err != nil {
		return Entry{}, err
	}

	var entry Entry
	err = attributevalue.UnmarshalMap(result.Item, &entry)
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}
