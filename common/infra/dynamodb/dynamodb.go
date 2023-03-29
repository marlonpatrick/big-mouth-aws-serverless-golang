package dynamodb

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
