package serverlessapi

type Endpoint struct {
	Method  string
	Path    string
	Target  func(Request) Response
	Timeout Timeout
}

type Timeout struct {
	Minutes int
	Seconds int
}

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
