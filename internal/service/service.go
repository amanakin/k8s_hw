package service

type Service interface {
	IncRequestsCount()
	GetRequestsCount() int
}
