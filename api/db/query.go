package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (dao Dao) QueryRecord(PK string) ([]Entry, error) {

	result, err := dao.db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName: aws.String(dao.db.TableName),
		KeyConditions: map[string]types.Condition{
			"PK": {
				ComparisonOperator: "EQ",
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{
						Value: PK,
					},
				},
			},
		},
	})

	log.Println("Query result: ", result)

	if err != nil {
		return []Entry{}, err
	}

	var entry []Entry
	err = attributevalue.UnmarshalListOfMaps(result.Items, &entry)
	if err != nil {
		return []Entry{}, err
	}
	return entry, nil
}
