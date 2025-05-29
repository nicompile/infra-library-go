package serverlessfunction

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Request struct {
	Headers                         map[string]string
	MultiValueHeaders               map[string][]string
	QueryStringParameters           map[string]string
	MultiValueQueryStringParameters map[string][]string
	PathParameters                  map[string]string
	Body                            string
}

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
}

func BuildLambdaHandler(function func(Request) Response) func(context.Context, events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return func(ctx context.Context, event events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
		req := requestFromAwsModel(event)
		resp := function(req)
		return resp.toAwsModel()
	}
}

func requestFromAwsModel(event events.APIGatewayProxyRequest) Request {
	return Request{
		Headers:                         event.Headers,
		MultiValueHeaders:               event.MultiValueHeaders,
		QueryStringParameters:           event.QueryStringParameters,
		MultiValueQueryStringParameters: event.MultiValueQueryStringParameters,
		PathParameters:                  event.PathParameters,
		Body:                            event.Body,
	}
}

func (r Response) toAwsModel() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        r.StatusCode,
		Headers:           r.Headers,
		MultiValueHeaders: r.MultiValueHeaders,
		Body:              r.Body,
	}
}
