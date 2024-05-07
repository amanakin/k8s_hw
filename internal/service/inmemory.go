package service

import "sync/atomic"

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
