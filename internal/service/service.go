package service

import (
	"bytes"
	"context"
	"encoding/json"
	"integration_service/internal/repository"
	"io"
	"net/http"
	"time"
)

type ProxyService struct {
	repo   *repository.RequestRepository
	client *http.Client
}

func NewProxyService(repo *repository.RequestRepository) *ProxyService {
	return &ProxyService{
		repo: repo,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (service *ProxyService) Proxy(ctx context.Context, requestId string) (*http.Response, error) {
	proxyRequest, err := service.repo.GetRequestByID(ctx, requestId)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if proxyRequest.Body != nil {
		bodyBytes, err := json.Marshal(proxyRequest.Body)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(proxyRequest.Method, proxyRequest.URL, body)
	if err != nil {
		return nil, err
	}

	for key, value := range proxyRequest.Headers {
		req.Header.Set(key, value)
	}

	if proxyRequest.AuthToken != "" {
		req.Header.Set("Authorization", proxyRequest.AuthToken)
	}

	query := req.URL.Query()
	for key, value := range proxyRequest.QueryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	if proxyRequest.TimeoutSeconds > 0 {
		service.client.Timeout = time.Duration(proxyRequest.TimeoutSeconds) * time.Second
	}

	return service.client.Do(req)
}
