package service

import (
	"context"
	"time"
)

type Service interface {
	StatisticsService
	TimeService
}

type StatisticsService interface {
	IncRequestsCount()
	GetRequestsCount() int
}

type TimeService interface {
	GetTime(ctx context.Context) (time.Time, error)
}
