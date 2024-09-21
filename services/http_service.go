package services

import (
	"fmt"
	"quote-server/types"
	"quote-server/utils"
)

type HttpService interface {
	GetEcho(query string) (*types.EchoResponse, error)
}

type HttpServiceImpl struct {
	HttpClient utils.HttpClient
}

func NewHttpService(client utils.HttpClient) HttpService {
	return &HttpServiceImpl{HttpClient: client}
}

func (sc *HttpServiceImpl) GetEcho(query string) (*types.EchoResponse, error) {
	url := fmt.Sprintf("https://postman-echo.com/get?query=%s", query)
	response, err := sc.HttpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching from Postman Echo: %v", err)
	}
	return response, nil
}
