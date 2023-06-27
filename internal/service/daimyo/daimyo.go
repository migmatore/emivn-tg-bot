package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"fmt"
)

type DaimyoStorage interface {
	Insert(ctx context.Context, daimyo domain.Daimyo) error
	GetAll(ctx context.Context) ([]*domain.Daimyo, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.Daimyo, error)
	GetByUsername(ctx context.Context, username string) (domain.Daimyo, error)
	GetByNickname(ctx context.Context, nickname string) (domain.Daimyo, error)
	UpdateUsername(ctx context.Context, old string, new string) error
	//GetTurnovers(ctx context.Context, date string) ([]*domain.SamuraiReport, error)
}

type SamuraiTurnoverStorage interface {
	GetTurnover(ctx context.Context, samuraiUsername string, date string, bankId int) (float64, error)
	GetTurnoverSumWithPeriod(ctx context.Context, samuraiUsername string, startDate string, endDate string,
		bankId int) (float64, error)
	//GetTurnoversByDate(ctx context.Context, date string) ([]*domain.SamuraiTurnover, error)
}

type ControllerTurnoverStorage interface {
	GetTurnover(ctx context.Context, samuraiUsername string, date string, bankId int) (float64, error)
	//GetTurnoversByDate(ctx context.Context, date string) ([]*domain.ControllerTurnover, error)
}

type SamuraiStorage interface {
	GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.Samurai, error)
}

type UserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type DaimyoService struct {
	transactor storage.Transactor

	storage                   DaimyoStorage
	samuraiTurnoverStorage    SamuraiTurnoverStorage
	controllerTurnoverStorage ControllerTurnoverStorage
	samuraiStorage            SamuraiStorage
	userRoleStorage           UserRoleStorage
	roleStorage               RoleStorage
}

func NewDaimyoService(
	t storage.Transactor,
	storage DaimyoStorage,
	samuraiTurnover SamuraiTurnoverStorage,
	controllerTurnover ControllerTurnoverStorage,
	samuraiStorage SamuraiStorage,
	userRoleStorage UserRoleStorage,
	roleStorage RoleStorage,
) *DaimyoService {
	return &DaimyoService{
		transactor:                t,
		storage:                   storage,
		samuraiTurnoverStorage:    samuraiTurnover,
		controllerTurnoverStorage: controllerTurnover,
		samuraiStorage:            samuraiStorage,
		userRoleStorage:           userRoleStorage,
		roleStorage:               roleStorage,
	}
}

func (s *DaimyoService) Create(ctx context.Context, dto domain.DaimyoDTO) error {
	daimyo := domain.Daimyo{
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		ShogunUsername: dto.ShogunUsername,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.DaimyoRole.String())
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

		if err := s.storage.Insert(txCtx, daimyo); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *DaimyoService) GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error) {
	daimyos, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	daimyoDTOs := make([]*domain.DaimyoDTO, 0)

	for _, item := range daimyos {
		daimyoDTO := domain.DaimyoDTO{
			Username:       item.Username,
			Nickname:       item.Nickname,
			CardsBalance:   item.CardsBalance,
			ShogunUsername: item.ShogunUsername,
		}

		daimyoDTOs = append(daimyoDTOs, &daimyoDTO)
	}

	return daimyoDTOs, nil
}

func (s *DaimyoService) GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.DaimyoDTO, error) {
	daimyos, err := s.storage.GetAllByShogun(ctx, shogunUsername)
	if err != nil {
		return nil, err
	}

	daimyoDTOs := make([]*domain.DaimyoDTO, 0)

	for _, item := range daimyos {
		daimyoDTO := domain.DaimyoDTO{
			Username:       item.Username,
			Nickname:       item.Nickname,
			CardsBalance:   item.CardsBalance,
			ShogunUsername: item.ShogunUsername,
		}

		daimyoDTOs = append(daimyoDTOs, &daimyoDTO)
	}

	return daimyoDTOs, nil
}

func (s *DaimyoService) GetByUsername(ctx context.Context, username string) (domain.DaimyoDTO, error) {
	daimyo, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return domain.DaimyoDTO{}, err
	}

	daimyoDTO := domain.DaimyoDTO{
		Username:       daimyo.Username,
		Nickname:       daimyo.Nickname,
		CardsBalance:   daimyo.CardsBalance,
		ShogunUsername: daimyo.ShogunUsername,
	}

	return daimyoDTO, nil
}

func (s *DaimyoService) GetByNickname(ctx context.Context, nickname string) (domain.DaimyoDTO, error) {
	daimyo, err := s.storage.GetByNickname(ctx, nickname)
	if err != nil {
		return domain.DaimyoDTO{}, err
	}

	daimyoDTO := domain.DaimyoDTO{
		Username:       daimyo.Username,
		Nickname:       daimyo.Nickname,
		CardsBalance:   daimyo.CardsBalance,
		ShogunUsername: daimyo.ShogunUsername,
	}

	return daimyoDTO, nil
}

func (s *DaimyoService) CreateSamuraiReport(ctx context.Context, daimyoUsername string, date string) ([]string, error) {
	reportMessages := make([]string, 0)

	//reports := make([]domain.SamuraiReport, 0)

	//controllerTurnovers, err := s.controllerTurnoverStorage.GetTurnoversByDate(ctx, date)
	//if err != nil {
	//
	//}
	//
	//for _, item := range controllerTurnovers {
	//	reports = append(reports, domain.SamuraiReport{
	//		SamuraiUsername:    item.SamuraiUsername,
	//		ControllerTurnover: item.Turnover,
	//		SamuraiTurnover:    0,
	//	})
	//}
	//s.controllerTurnoverStorage.GetTurnoversByDate(ctx, date)

	samurais, err := s.samuraiStorage.GetAllByDaimyo(ctx, daimyoUsername)
	if err != nil {
		return nil, err
	}

	for _, samurai := range samurais {
		var sberSamuraiTurnover, sberControllerTurnover, tinSamuraiTurnover, tinControllerTurnover float64
		var str string
		var err error

		tinSamuraiTurnover, err = s.samuraiTurnoverStorage.GetTurnover(ctx, samurai.Username, date, 1)
		if err != nil {
			//str += "Ошибка получения данных\n"
		}

		tinControllerTurnover, err = s.controllerTurnoverStorage.GetTurnover(ctx, samurai.Username, date, 1)
		if err != nil {
			//str += "Ошибка получения данных\n"
		}

		//str += "Сбербанк\n"
		sberSamuraiTurnover, err = s.samuraiTurnoverStorage.GetTurnover(ctx, samurai.Username, date, 2)
		if err != nil {
			//str += "Ошибка получения данных\n"
		}

		sberControllerTurnover, err = s.controllerTurnoverStorage.GetTurnover(ctx, samurai.Username, date, 2)
		if err != nil {
			//str += "Ошибка получения данных\n"
		}

		if (sberControllerTurnover-sberSamuraiTurnover == 0) && (tinControllerTurnover-tinSamuraiTurnover == 0) {
			str += fmt.Sprintf("Расхождения по %s отсутсвуют\n\n", samurai.Username)
			reportMessages = append(reportMessages, str)

			continue
		}

		str += fmt.Sprintf("%s (%s)\n\n", samurai.Nickname, date)
		str += fmt.Sprintf("Всего\n%d / %d / %d\n\n", int(tinControllerTurnover+sberControllerTurnover),
			int(tinSamuraiTurnover+sberSamuraiTurnover),
			int((tinControllerTurnover-tinSamuraiTurnover)+(sberControllerTurnover-sberSamuraiTurnover)))
		str += fmt.Sprintf("Тинькофф\n%d / %d / %d\n", int(tinControllerTurnover), int(tinSamuraiTurnover),
			int(tinControllerTurnover-tinSamuraiTurnover))
		str += fmt.Sprintf("Сбербанк\n%d / %d / %d", int(sberControllerTurnover), int(sberSamuraiTurnover),
			int(sberControllerTurnover-sberSamuraiTurnover))

		reportMessages = append(reportMessages, str)
	}

	return reportMessages, nil
}

func (s *DaimyoService) CreateSamuraiReportWithPeriod(
	ctx context.Context,
	daimyoUsername string,
	startDate string,
	endDate string,
) ([]string, error) {
	daimyo, err := s.storage.GetByUsername(ctx, daimyoUsername)
	if err != nil {
		return nil, err
	}

	reportMessages := make([]string, 0)

	samurais, err := s.samuraiStorage.GetAllByDaimyo(ctx, daimyoUsername)
	if err != nil {
		return nil, err
	}

	var turnoverSum float64

	samuraiTurnovers := make([]string, 0)

	for _, samurai := range samurais {
		var err error
		var tinTurnover, sberTurnover float64

		var samuraiTurnover string

		samuraiTurnover += fmt.Sprintf("%s\n", samurai.Nickname)

		tinTurnover, err = s.samuraiTurnoverStorage.GetTurnoverSumWithPeriod(ctx, samurai.Username, startDate, endDate, 1)
		if err != nil {

		}

		sberTurnover, err = s.samuraiTurnoverStorage.GetTurnoverSumWithPeriod(ctx, samurai.Username, startDate, endDate, 2)
		if err != nil {

		}

		samuraiTurnover += fmt.Sprintf("%d / %d", int(tinTurnover), int(sberTurnover))
		turnoverSum += tinTurnover + sberTurnover

		samuraiTurnovers = append(samuraiTurnovers, samuraiTurnover)
	}

	reportMessages = append(reportMessages, fmt.Sprintf(
		"%s (%s - %s)\nОборот: %d\n0.0015 -> %d",
		daimyo.Nickname,
		startDate,
		endDate,
		int(turnoverSum),
		int(0.0015*turnoverSum),
	))

	reportMessages = append(reportMessages, samuraiTurnovers...)

	return reportMessages, nil
}
