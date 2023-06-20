package replenishment_request_status

import "context"

type ReplenishmentRequestStatusStorage interface {
	GetId(ctx context.Context, name string) (int, error)
	GetById(ctx context.Context, statusId int) (string, error)
}

type ReplenishmentRequestStatusService struct {
	storage ReplenishmentRequestStatusStorage
}

func NewReplenishmentRequestStatusService(storage ReplenishmentRequestStatusStorage) *ReplenishmentRequestStatusService {
	return &ReplenishmentRequestStatusService{storage: storage}
}
