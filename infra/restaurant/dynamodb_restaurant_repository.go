package restaurant

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBRestaurantRepository struct {
	dynamoDBClient *dynamodb.Client
}

func NewDynamoDBRestaurantRepository(dynamoDBClient *dynamodb.Client) DynamoDBRestaurantRepository {
	return DynamoDBRestaurantRepository{dynamoDBClient}
}

func (dynamoRepo DynamoDBRestaurantRepository) FindAllRestaurants(limit int) ([]map[string]types.AttributeValue, error) {

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
