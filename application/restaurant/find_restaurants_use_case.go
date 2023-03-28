package restaurant

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/marlonpatrick/big-mouth-aws-serverless-golang/domain/restaurant"
)

type FindRestaurantsUseCase struct {
	repository restaurant.RestaurantRepository
}

func NewFindRestaurantsUseCase(repository restaurant.RestaurantRepository) *FindRestaurantsUseCase {
	return &FindRestaurantsUseCase{repository}
}

func (useCase *FindRestaurantsUseCase) Execute() ([]map[string]types.AttributeValue, error) {

	items, err := useCase.repository.FindAllRestaurants(10) // TODO 10 => env variable

	if err != nil {
		log.Printf("Couldn't scan restaurants. Here's why: %v\n", err)
		return nil, err
	}

	// err = attributevalue.UnmarshalListOfMaps(response.Items, &movies)
	// if err != nil {
	// 	log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
	// }
	return items, nil
}
