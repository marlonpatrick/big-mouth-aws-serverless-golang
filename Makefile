build-GetIndexPageFunction:  
	GOOS=linux go build get-index-page/get_index_page_handler.go
	cp ./get_index_page_handler $(ARTIFACTS_DIR)/bootstrap