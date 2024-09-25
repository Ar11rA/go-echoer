package services

import (
	"fmt"
	"os"
	"quote-server/types"
	"quote-server/utils"
	"sync"
)

type HttpService interface {
	GetEcho(query string) (*types.EchoResponse, error)
	PostEcho(body types.EchoRequest) (*types.EchoResponse, error)
	GetQuotes(limit int32) ([]*types.QuoteResponse, error)
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

func (sc *HttpServiceImpl) GetQuotes(limit int32) ([]*types.QuoteResponse, error) {
	results := make([]*types.QuoteResponse, 0, limit)
	ch := make(chan *types.QuoteResponse, limit) // Channel to gather results
	errCh := make(chan error, limit)             // Channel to gather errors
	wg := &sync.WaitGroup{}                      // WaitGroup to wait for all goroutines to finish
	url := fmt.Sprintf("%s/random", os.Getenv("QUOTE_BASE_URL"))

	// Loop for each quote request
	for i := int32(0); i < limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			body, err := sc.HttpClient.Get(url)

			if err != nil {
				errCh <- err
				return
			}

			resp, err := utils.ParseResponse[types.QuoteResponse](body)
			if err != nil {
				errCh <- err
				return
			}

			ch <- resp
		}()
	}

	// Close channels once all requests are done
	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
	}()

	// Collect results and errors
	for {
		select {
		case res, ok := <-ch:
			if ok {
				results = append(results, res)
			}
		case err, ok := <-errCh:
			if ok {
				return nil, err
			}
		default:
			if len(results) == int(limit) && len(errCh) == 0 {
				return results, nil
			}
		}
	}
}
