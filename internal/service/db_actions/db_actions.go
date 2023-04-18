package db_actions

import "context"

type DbActionsStorage interface {
	Read(ctx context.Context) ([]string, error)
	Write(ctx context.Context, text string) error
}

type DbActionsService struct {
	storage DbActionsStorage
}

func NewDbActionsService(s DbActionsStorage) *DbActionsService {
	return &DbActionsService{storage: s}
}

func (s *DbActionsService) DoAction(ctx context.Context, actionName string, text string) ([]string, error) {
	switch actionName {
	case "Read":
		return s.storage.Read(ctx)
	case "Write":
		return nil, s.storage.Write(ctx, text)
	}

	return nil, nil
}
