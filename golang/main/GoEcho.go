package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Phrase string `json:"phrase"`
}

func HandleRequest(ctx context.Context, myEvent MyEvent) (string, error) {
	fmt.Printf("Hello %s!", myEvent.Phrase)
	return fmt.Sprintf("Hello %s!", myEvent.Phrase), nil
}

func main() {
	lambda.Start(HandleRequest)
}
