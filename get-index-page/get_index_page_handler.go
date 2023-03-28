package main

import (
	_ "embed"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	//go:embed resources/index.html
	IndexPageHtml string
)

func handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if len(IndexPageHtml) == 0 {
		/*
			[IMPLEMENTED] https://docs.aws.amazon.com/lambda/latest/dg/golang-exceptions.html
				For example, API Gateway treats all invocation and function errors as internal errors.
				If the Lambda API rejects the invocation request, API Gateway returns a 500 error code.

				If the function runs but returns an error, or returns a response in the wrong format, API Gateway returns a 502 error code.

				>>> To customize the error response, you must catch errors in your code and format a response in the required format. <<<

			[TO DO] https://docs.aws.amazon.com/apigateway/latest/developerguide/handle-errors-in-lambda-integration.html
				Handle standard Lambda errors in API Gateway
					You can return a error and then map that error with a Api Gateway Response

			[TO DO] Implement a kind of global error handler (GEH). The handler returns an error and GEH convert into a api gateway response.
		*/
		return events.APIGatewayProxyResponse{
			Body:       "This page is currently unavailable",
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/text; charset=UTF-8",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       IndexPageHtml,
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html; charset=UTF-8",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
