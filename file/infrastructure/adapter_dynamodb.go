package infrastructure

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
)

type DynamoDBRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBRepository(client *dynamodb.Client) *DynamoDBRepository {
	return &DynamoDBRepository{Client: client}
}

func (db *DynamoDBRepository) CreateExcel(data []domain.RouteStore) error {
	for _, d := range data {
		av, err := attributevalue.MarshalMap(d)
		if err != nil {
			fmt.Printf("Got error marshalling data: %s\n", err)
			return nil
		}

		output, err := db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(domain.TableRouteStore), Item: av,
		})
		_ = output
		if err != nil {
			fmt.Printf("Couldn't add item to table.: %v\n", err)
		}
	}
	return nil
}

func (db *DynamoDBRepository) GetAllExcel(routeName string) ([]domain.RouteStore, error) {
	var routeStores []domain.RouteStore
	keyExpression := expression.Key("route_name").Equal(expression.Value(routeName))
	expr, err := expression.NewBuilder().WithKeyCondition(keyExpression).Build()
	if err != nil {
		return nil, err
	}
	response, err := db.Client.Query(
		context.TODO(),
		&dynamodb.QueryInput{
			TableName:                 aws.String(domain.TableRouteStore),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		},
	)
	if err != nil {
		return nil, err
	}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &routeStores)
	return routeStores, err
}
