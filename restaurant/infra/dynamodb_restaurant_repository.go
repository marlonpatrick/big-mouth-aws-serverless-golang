package restaurant

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	dynamodb_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/common/infra/dynamodb"
)

type DynamoDBRestaurantRepository struct {
	dynamoDBClient *dynamodb.Client
}

func NewDynamoDBRestaurantRepository() (*DynamoDBRestaurantRepository, error) {

	dynamoDBClient, err := dynamodb_infra.NewDynamoDBClient()

	if err != nil {
		return nil, err
	}

	return &DynamoDBRestaurantRepository{dynamoDBClient}, nil
}

func (dynamoRepo *DynamoDBRestaurantRepository) FindAllRestaurants(limit int) ([]map[string]types.AttributeValue, error) {

	response, err := dynamoRepo.dynamoDBClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("big_mouth_restaurants"), // TODO env variable + single table data modeling
		Limit:     aws.Int32(int32(limit)),
	})

	if err != nil {
		log.Printf("Couldn't scan restaurants. Here's why: %v\n", err)
		return nil, err
	}

	// err = attributevalue.UnmarshalListOfMaps(response.Items, &movies)
	// if err != nil {
	// 	log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
	// }
	return response.Items, nil
}
