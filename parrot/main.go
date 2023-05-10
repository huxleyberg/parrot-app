package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	name, found := os.LookupEnv("DEFAULT_NAME")
	if !found {
		name = "world"
	}
	if tmpName, ok := request.QueryStringParameters["name"]; ok {
		name = tmpName
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %s", string(name)),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}
