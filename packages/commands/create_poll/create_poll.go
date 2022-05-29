package main

import "log"

type Request struct {
	Location string `json:"location"`
}

type R struct {
	Type int `json:"type"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       R                 `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	log.Println(in)

	data := R{
		Type: 1,
	}

	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"

	return &Response{
		Headers: headers,
		Body:    data,
	}, nil
}
