package card

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type CardStorage interface {
	Insert(ctx context.Context, card domain.Card) error
	GetByUsername(ctx context.Context, bankId int, daimyoUsername string) ([]*domain.Card, error)
	GetByName(ctx context.Context, name string) (domain.Card, error)
	GetBankNames(ctx context.Context) ([]*domain.Bank, error)
	GetBankIdByName(ctx context.Context, bankName string) (int, error)
}

type CardService struct {
	transactor storage.Transactor

	storage CardStorage
}

func NewCardService(t storage.Transactor, s CardStorage) *CardService {
	return &CardService{
		transactor: t,
		storage:    s,
	}
}

func (s *CardService) Create(ctx context.Context, dto domain.CardDTO) error {
	bankName, err := s.storage.GetBankIdByName(ctx, dto.BankType)
	if err != nil {
		return err
	}

	card := domain.Card{
		Name:           dto.Name,
		LastDigits:     dto.LastDigits,
		DailyLimit:     dto.DailyLimit,
		DaimyoUsername: dto.DaimyoUsername,
		BankTypeId:     bankName,
	}

	return s.storage.Insert(ctx, card)
}

func (s *CardService) GetByUsername(ctx context.Context, bankName string, daimyoUsername string) ([]*domain.CardDTO, error) {
	bankId, err := s.storage.GetBankIdByName(ctx, bankName)
	if err != nil {
		return nil, err
	}

	cards, err := s.storage.GetByUsername(ctx, bankId, daimyoUsername)
	if err != nil {
		return nil, err
	}

	cardDTOs := make([]*domain.CardDTO, 0)

	for _, item := range cards {
		cardDTO := domain.CardDTO{
			Name:           item.Name,
			LastDigits:     item.LastDigits,
			DailyLimit:     item.DailyLimit,
			DaimyoUsername: item.DaimyoUsername,
			BankType:       bankName,
		}

		cardDTOs = append(cardDTOs, &cardDTO)
	}

	return cardDTOs, nil
}

func (s *CardService) GetBankNames(ctx context.Context) ([]*domain.BankDTO, error) {
	banks, err := s.storage.GetBankNames(ctx)
	if err != nil {
		return nil, err
	}

	bankDTOs := make([]*domain.BankDTO, 0)

	for _, item := range banks {
		bankDTO := domain.BankDTO{
			BankId: item.BankId,
			Name:   item.Name,
		}

		bankDTOs = append(bankDTOs, &bankDTO)
	}

	return bankDTOs, nil
}
