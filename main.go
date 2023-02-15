// main.go
package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	A int
	B int
}

type Response struct {
	Result int
}

func GCF(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func HandleRequests(ctx context.Context, req Request) (Response, error) {
	return Response{Result: GCF(req.A, req.B)}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleRequests)
}
