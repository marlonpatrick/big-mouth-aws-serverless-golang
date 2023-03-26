package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	usecase "github.com/marlonpatrick/big-mouth-aws-serverless-golang/get-restaurants/use-case"
	dynamodb_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/infra/dynamodb"
	restaurant_infra "github.com/marlonpatrick/big-mouth-aws-serverless-golang/infra/restaurant"
)

var (
	findAllRestaurantsUseCase usecase.FindAllRestaurantsUseCase
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	items, err := findAllRestaurantsUseCase.Execute()

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "findRestaurants deu erro",
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/text; charset=UTF-8",
			},
		}, nil
	}

	marshalledItems, err := json.Marshal(items)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Marshal deu erro",
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

	findAllRestaurantsUseCase = usecase.NewFindAllRestaurantsUseCase(dynamoDBRestaurantRepository)
}

func main() {
	lambda.Start(handler)
}
