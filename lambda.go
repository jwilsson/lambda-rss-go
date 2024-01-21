package main

import (
	"bytes"
	"lambda-rss/sites"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createResponse(body string, statusCode int) events.LambdaFunctionURLResponse {
	headers := map[string]string{}

	if statusCode == 200 {
		headers["Content-Type"] = "application/rss+xml"
	} else {
		headers["Content-Type"] = "text/plain"
	}

	return events.LambdaFunctionURLResponse{
		Body:            body,
		Headers:         headers,
		IsBase64Encoded: false,
		StatusCode:      statusCode,
	}
}

func handleRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	s := request.QueryStringParameters["s"]
	site := sites.GetSite(s)

	if site != nil {
		var bodyBuffer bytes.Buffer

		err := createBody(*site, &bodyBuffer)
		if err != nil {
			return createResponse("", 500), err
		}

		return createResponse(bodyBuffer.String(), 200), nil
	}

	return createResponse("", 400), nil
}

func main() {
	lambda.Start(handleRequest)
}
