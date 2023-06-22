package card

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"fmt"
)

type CardStorage interface {
	Insert(ctx context.Context, card domain.Card) error
	GetAllByUsername(ctx context.Context, bankId int, ownerUsername string) ([]*domain.Card, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.Card, error)
	GetByUsername(ctx context.Context, daimyoUsername string) (domain.Card, error)
	GetByName(ctx context.Context, name string) (domain.Card, error)
	GetById(ctx context.Context, cardId int) (domain.Card, error)
	UpdateLimit(ctx context.Context, name string, limit int) error
	GetBankNames(ctx context.Context) ([]*domain.Bank, error)
	GetBankIdByName(ctx context.Context, bankName string) (int, error)
	GetBankById(ctx context.Context, bankId int) (domain.Bank, error)
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
	bankId, err := s.storage.GetBankIdByName(ctx, dto.BankType)
	if err != nil {
		return err
	}

	card := domain.Card{
		Name:          dto.Name,
		OwnerUsername: dto.OwnerUsername,
		LastDigits:    dto.LastDigits,
		DailyLimit:    dto.DailyLimit,
		Balance:       0,
		BankTypeId:    bankId,
	}

	return s.storage.Insert(ctx, card)
}

func (s *CardService) GetAllByUsername(ctx context.Context, bankName string, ownerUsername string) ([]*domain.CardDTO, error) {
	bankId, err := s.storage.GetBankIdByName(ctx, bankName)
	if err != nil {
		return nil, err
	}

	cards, err := s.storage.GetAllByUsername(ctx, bankId, ownerUsername)
	if err != nil {
		return nil, err
	}

	cardDTOs := make([]*domain.CardDTO, 0)

	for _, item := range cards {
		cardDTO := domain.CardDTO{
			Name:          item.Name,
			OwnerUsername: item.OwnerUsername,
			LastDigits:    item.LastDigits,
			DailyLimit:    item.DailyLimit,
			Balance:       item.Balance,
			BankType:      bankName,
		}

		cardDTOs = append(cardDTOs, &cardDTO)
	}

	return cardDTOs, nil
}

func (s *CardService) GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.CardDTO, error) {
	cards, err := s.storage.GetAllByShogun(ctx, shogunUsername)
	if err != nil {
		return nil, err
	}

	cardDTOs := make([]*domain.CardDTO, 0)

	for _, item := range cards {
		cardDTO := domain.CardDTO{
			Name:          item.Name,
			OwnerUsername: item.OwnerUsername,
			LastDigits:    item.LastDigits,
			DailyLimit:    item.DailyLimit,
			Balance:       item.Balance,
		}

		cardDTOs = append(cardDTOs, &cardDTO)
	}

	return cardDTOs, nil
}

func (s *CardService) GetCardsBalancesByShogun(ctx context.Context, shogunUsername string) ([]string, error) {
	cardsBalances := make([]string, 0)

	cards, err := s.storage.GetAllByShogun(ctx, shogunUsername)
	if err != nil {
		return nil, nil
	}

	for _, item := range cards {
		cardsBalances = append(cardsBalances, fmt.Sprintf("%s - %f", item.Name, item.Balance))
	}

	return cardsBalances, nil
}

func (s *CardService) GetByUsername(ctx context.Context, ownerUsername string) (domain.CardDTO, error) {
	card, err := s.storage.GetByUsername(ctx, ownerUsername)
	if err != nil {
		return domain.CardDTO{}, err
	}

	bank, err := s.storage.GetBankById(ctx, card.BankTypeId)
	if err != nil {
		return domain.CardDTO{}, err
	}

	cardDTO := domain.CardDTO{
		Name:          card.Name,
		OwnerUsername: card.OwnerUsername,
		LastDigits:    card.LastDigits,
		DailyLimit:    card.DailyLimit,
		Balance:       card.Balance,
		BankType:      bank.Name,
	}

	return cardDTO, nil
}

func (s *CardService) GetByName(ctx context.Context, name string) (domain.CardDTO, error) {
	card, err := s.storage.GetByName(ctx, name)
	if err != nil {
		return domain.CardDTO{}, err
	}

	bank, err := s.storage.GetBankById(ctx, card.BankTypeId)
	if err != nil {
		return domain.CardDTO{}, err
	}

	cardDTO := domain.CardDTO{
		Name:          card.Name,
		OwnerUsername: card.OwnerUsername,
		LastDigits:    card.LastDigits,
		DailyLimit:    card.DailyLimit,
		Balance:       card.Balance,
		BankType:      bank.Name,
	}

	return cardDTO, nil
}

func (s *CardService) ChangeLimit(ctx context.Context, name string, limit int) error {
	return s.storage.UpdateLimit(ctx, name, limit)
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
