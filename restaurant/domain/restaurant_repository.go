package restaurant

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RestaurantRepository interface {
	FindAllRestaurants(limit int) ([]map[string]types.AttributeValue, error)
}
