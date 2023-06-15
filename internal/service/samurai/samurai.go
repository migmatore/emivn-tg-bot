package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"emivn-tg-bot/pkg/logging"
	"github.com/mr-linch/go-tg"
	"time"
)

type SamuraiStorage interface {
	Insert(ctx context.Context, samurai domain.Samurai) error
	GetByUsername(ctx context.Context, username string) (domain.Samurai, error)
	SetChatId(ctx context.Context, username string, id int64) error
	GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.Samurai, error)
}

type SamuraiTurnoverStorage interface {
	Insert(ctx context.Context, turnover domain.SamuraiTurnover) error
	CheckIfExists(ctx context.Context, date string, bankId int) (bool, error)
	GetByDateAndBank(ctx context.Context, date string, bankId int) (domain.SamuraiTurnover, error)
	Update(ctx context.Context, turnover domain.SamuraiTurnover) error
}

type CardStorage interface {
	GetBankIdByName(ctx context.Context, bankName string) (int, error)
}

type UserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type SamuraiService struct {
	transactor storage.Transactor

	storage                SamuraiStorage
	samuraiTurnoverStorage SamuraiTurnoverStorage
	cardStorage            CardStorage
	userRoleStorage        UserRoleStorage
	roleStorage            RoleStorage
}

func NewSamuraiService(
	transactor storage.Transactor,
	samuraiStorage SamuraiStorage,
	samuraiTurnoverStorage SamuraiTurnoverStorage,
	cardStorage CardStorage,
	userRole UserRoleStorage,
	role RoleStorage,
) *SamuraiService {
	return &SamuraiService{
		transactor:             transactor,
		storage:                samuraiStorage,
		samuraiTurnoverStorage: samuraiTurnoverStorage,
		cardStorage:            cardStorage,
		userRoleStorage:        userRole,
		roleStorage:            role,
	}
}

func (s *SamuraiService) Create(ctx context.Context, dto domain.SamuraiDTO) error {
	samurai := domain.Samurai{
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		DaimyoUsername: dto.DaimyoUsername,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.SamuraiRole.String())
	if err != nil {
		return err
	}

	userRole := domain.UserRole{
		Username: dto.Username,
		RoleId:   roleId,
	}

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.userRoleStorage.Insert(txCtx, userRole); err != nil {
			return err
		}

		if err := s.storage.Insert(txCtx, samurai); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *SamuraiService) SetChatId(ctx context.Context, username string, id tg.ChatID) error {
	samurai, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	// TODO: Error refactoring
	if samurai.ChatId != nil {
		logging.GetLogger(ctx).Error("Error: samurai chat id is already set")
	}

	return s.storage.SetChatId(ctx, username, int64(id))
}

func (s *SamuraiService) GetByUsername(ctx context.Context, username string) (domain.SamuraiDTO, error) {
	samurai, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return domain.SamuraiDTO{}, nil
	}

	samuraiDTO := domain.SamuraiDTO{
		Username:         samurai.Username,
		Nickname:         samurai.Nickname,
		DaimyoUsername:   samurai.DaimyoUsername,
		TurnoverPerShift: samurai.TurnoverPerShift,
		ChatId:           samurai.ChatId,
	}

	return samuraiDTO, nil
}

func (s *SamuraiService) GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.SamuraiDTO, error) {
	samurais, err := s.storage.GetAllByDaimyo(ctx, daimyoUsername)
	if err != nil {
		return nil, err
	}

	samuraiDTOs := make([]*domain.SamuraiDTO, 0)

	for _, item := range samurais {
		samuraiDTO := domain.SamuraiDTO{
			Username:         item.Username,
			Nickname:         item.Nickname,
			DaimyoUsername:   item.DaimyoUsername,
			TurnoverPerShift: item.TurnoverPerShift,
			ChatId:           item.ChatId,
		}

		samuraiDTOs = append(samuraiDTOs, &samuraiDTO)
	}

	return samuraiDTOs, nil
}

func (s *SamuraiService) CreateTurnover(ctx context.Context, dto domain.SamuraiTurnoverDTO) error {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	bankId, err := s.cardStorage.GetBankIdByName(ctx, dto.BankTypeName)
	if err != nil {
		return err
	}

	exists, err := s.samuraiTurnoverStorage.CheckIfExists(ctx, yesterday, bankId)
	if err != nil {
		return err
	}

	initialAmount := dto.FinalAmount

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if !exists {
			turnover := domain.SamuraiTurnover{
				SamuraiUsername: dto.SamuraiUsername,
				StartDate:       yesterday,
				InitialAmount:   0,
				FinalAmount:     dto.FinalAmount,
				Turnover:        dto.FinalAmount,
				BankTypeId:      bankId,
			}

			if err := s.samuraiTurnoverStorage.Insert(ctx, turnover); err != nil {
				return err
			}
		} else {
			turnover, err := s.samuraiTurnoverStorage.GetByDateAndBank(ctx, yesterday, bankId)
			if err != nil {
				return err
			}

			turnover.FinalAmount = dto.FinalAmount + turnover.InitialAmount
			turnover.Turnover = dto.FinalAmount

			initialAmount = turnover.FinalAmount

			if err := s.samuraiTurnoverStorage.Update(ctx, turnover); err != nil {
				return err
			}
		}

		turnover := domain.SamuraiTurnover{
			SamuraiUsername: dto.SamuraiUsername,
			StartDate:       time.Now().Format("2006-01-02"),
			InitialAmount:   initialAmount,
			FinalAmount:     0,
			Turnover:        initialAmount,
			BankTypeId:      bankId,
		}

		newExists, err := s.samuraiTurnoverStorage.CheckIfExists(ctx, time.Now().Format("2006-01-02"), bankId)
		if err != nil {
			return err
		}

		if !newExists {
			if err := s.samuraiTurnoverStorage.Insert(ctx, turnover); err != nil {
				return err
			}

			return nil
		}

		if err := s.samuraiTurnoverStorage.Update(ctx, turnover); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
