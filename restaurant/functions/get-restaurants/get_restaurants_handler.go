package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	restaurant_application "github.com/marlonpatrick/big-mouth-aws-serverless-golang/restaurant/application"
	restaurant_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/restaurant/infra"
)

var (
	findAllRestaurantsUseCase *restaurant_application.FindRestaurantsUseCase
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

	dynamoDBRestaurantRepository, err := restaurant_infra.NewDynamoDBRestaurantRepository()

	if err != nil {
		panic(err)
	}

	findAllRestaurantsUseCase = restaurant_application.NewFindRestaurantsUseCase(dynamoDBRestaurantRepository)
}

func main() {
	lambda.Start(handler)
}
