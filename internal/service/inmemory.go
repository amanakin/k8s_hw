package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	worldAPIURL = "http://worldtimeapi.org/api/timezone/Europe/Moscow"
)

type WorldAPIResponse struct {
	// DateTime in format '2024-05-15T23:12:33.917411+03:00'
	DateTime string `json:"datetime"`
}

func NewInMemoryService() *InMemoryService {
	return &InMemoryService{}
}

type InMemoryService struct {
	requests int64
}

func (s *InMemoryService) IncRequestsCount() {
	atomic.AddInt64(&s.requests, 1)
}
func (s *InMemoryService) GetRequestsCount() int {
	return int(atomic.LoadInt64(&s.requests))
}

func (s *InMemoryService) GetTime(ctx context.Context) (time.Time, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", worldAPIURL, nil)
	if err != nil {
		return time.Time{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("bad worldapi status: %s", resp.Status)
	}

	var apiResp WorldAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.Parse(time.RFC3339Nano, apiResp.DateTime)
	if err != nil {
		return time.Time{}, err
	}
	
	return t, nil
}
