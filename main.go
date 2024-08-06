package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func main() {
	fmt.Println("Hello, Fucking World!!")
}

func HandleRequest(req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Hello, Are you still there, being the fucking idiot??? Come on!!",
	}, nil
}
