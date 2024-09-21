package services

import (
	"fmt"
	"os"
	"quote-server/types"
	"quote-server/utils"
)

type HttpService interface {
	GetEcho(query string) (*types.EchoResponse, error)
	PostEcho(body types.EchoRequest) (*types.EchoResponse, error)
}

type HttpServiceImpl struct {
	HttpClient utils.HttpClient
}

func NewHttpService(client utils.HttpClient) HttpService {
	return &HttpServiceImpl{HttpClient: client}
}

func (sc *HttpServiceImpl) GetEcho(query string) (*types.EchoResponse, error) {
	url := fmt.Sprintf("%s/get?query=%s", os.Getenv("HTTP_BASE_URL"), query)
	body, err := sc.HttpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching from Postman Echo: %v", err)
	}
	response, err := utils.ParseResponse[types.EchoResponse](body)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	return response, nil
}

func (sc *HttpServiceImpl) PostEcho(body types.EchoRequest) (*types.EchoResponse, error) {
	url := fmt.Sprintf("%s/post", os.Getenv("HTTP_BASE_URL"))
	responseBytes, err := sc.HttpClient.Post(url, body)
	if err != nil {
		return nil, fmt.Errorf("error posting to Postman Echo: %v", err)
	}

	response, err := utils.ParseResponse[types.EchoResponse](responseBytes)
	if err != nil {
		return nil, err
	}

	return response, nil
}
