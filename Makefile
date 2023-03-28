build-GetIndexPageFunction:  
	GOOS=linux go build functions/get-index-page/get_index_page_handler.go
	cp ./get_index_page_handler $(ARTIFACTS_DIR)/bootstrap

build-GetRestaurantsFunction:  
	GOOS=linux go build functions/get-restaurants/get_restaurants_handler.go
	cp ./get_restaurants_handler $(ARTIFACTS_DIR)/bootstrap