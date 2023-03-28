package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	restaurant_usecases "github.com/marlonpatrick/big-mouth-aws-serverless-golang/application/restaurant"
	dynamodb_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/infra/dynamodb"
	restaurant_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/infra/restaurant"
)

var (
	findAllRestaurantsUseCase *restaurant_usecases.FindRestaurantsUseCase
)

func handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	items, err := findAllRestaurantsUseCase.Execute()

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "An error occurred while querying the restaurants.",
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/text; charset=UTF-8",
			},
		}, nil
	}

	marshalledItems, err := json.Marshal(items)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "An error occurred while marshal the restaurants.",
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/text; charset=UTF-8",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(marshalledItems),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
	}, nil
}

func init() {
	dynamoDBClient, err := dynamodb_infra.NewDynamoDBClient()

	if err != nil {
		panic(err)
	}

	dynamoDBRestaurantRepository := restaurant_infra.NewDynamoDBRestaurantRepository(dynamoDBClient)

	findAllRestaurantsUseCase = restaurant_usecases.NewFindRestaurantsUseCase(dynamoDBRestaurantRepository)
}

func main() {
	lambda.Start(handler)
}
