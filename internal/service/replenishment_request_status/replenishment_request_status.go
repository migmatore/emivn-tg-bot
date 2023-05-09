package replenishment_request_status

import "context"

type ReplenishmentRequestStatusStorage interface {
	GetId(ctx context.Context, name string) (int, error)
}

type ReplenishmentRequestStatusService struct {
	storage ReplenishmentRequestStatusStorage
}

func NewReplenishmentRequestStatusService(storage ReplenishmentRequestStatusStorage) *ReplenishmentRequestStatusService {
	return &ReplenishmentRequestStatusService{storage: storage}
}
