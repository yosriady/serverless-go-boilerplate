package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type Todo struct {
	ID          string  `json:"id"`
	Description string 	`json:"description"`
	Done        bool   	`json:"done"`
	CreatedAt   string 	`json:"created_at"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{ // Success HTTP response
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}