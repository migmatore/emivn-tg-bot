package service

import "context"

type TransactorStorage interface {
	WithinTransaction(ctx context.Context, txFunc func(context context.Context) error) error
}

type TransactorService struct {
	storage TransactorStorage
}

func NewTransactorService(storage TransactorStorage) *TransactorService {
	return &TransactorService{storage: storage}
}

func (s *TransactorService) WithinTransaction(ctx context.Context, txFunc func(context context.Context) error) error {
	return s.storage.WithinTransaction(ctx, txFunc)
}
