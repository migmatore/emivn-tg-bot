package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerDaimyoHandler() {
	h.Message(h.DaimyoHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepDaimyoMainMenuHandler))

	// make replenishment request
	h.Message(h.DaimyoHandler.RepReqChooseCardHandler, h.isSessionStep(domain.SessionStepDaimyoChooseReplenishmentRequestBank))
	h.Message(h.DaimyoHandler.EnterRepReqAmountHandler, h.isSessionStep(domain.SessionStepDaimyoEnterReplenishmentRequestAmount))
	h.Message(h.DaimyoHandler.MakeRepReqHandler, h.isSessionStep(domain.SessionStepDaimyoMakeReplenishmentRequest))
	h.Message(h.DaimyoHandler.ChangeRepReqAmountHandler, h.isSessionStep(domain.SessionStepDaimyoChangeReplenishmentRequestAmount))

	// daimyo replenishment requests menu
	h.Message(h.DaimyoHandler.RepReqMenuHandler, h.isSessionStep(domain.SessionStepDaimyoRepReqMenuHandler))

	h.Message(h.DaimyoHandler.ObjRepReqSelectHandler, h.isSessionStep(domain.SessionStepDaimyoObjRepReqSelectHandler))
	h.Message(h.DaimyoHandler.ObjRepReqActionHandler, h.isSessionStep(domain.SessionStepDaimyoObjRepReqActionHandler))
	h.Message(h.DaimyoHandler.ObjRepReqAnotherAmountHandler, h.isSessionStep(domain.SessionStepDaimyoRepReqAnotherAmountHandler))

	// report
	h.Message(h.DaimyoHandler.ReportMenuHandler, h.isSessionStep(domain.SessionStepDaimyoReportMenuHandler))
	h.Message(h.DaimyoHandler.ReportPeriodMenuHandler, h.isSessionStep(domain.SessionStepDaimyoReportPeriodMenuHandler))

	// report with period
	h.Message(h.DaimyoHandler.ReportStartPeriodHandler, h.isSessionStep(domain.SessionStepDaimyoReportStartPeriod))
	h.Message(h.DaimyoHandler.ReportEndPeriodHandler, h.isSessionStep(domain.SessionStepDaimyoReportEndPeriod))

	// hierarchy
	h.Message(h.DaimyoHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepDaimyoHierarchyMenuHandler))
	h.Message(h.DaimyoHandler.EnterSamuraiNicknameHandler, h.isSessionStep(domain.SessionStepDaimyoCreateSamuraiNickname))
	h.Message(h.DaimyoHandler.SamuraiCreationHandler, h.isSessionStep(domain.SessionStepDaimyoSamuraiCreationMethod))
	h.Message(h.DaimyoHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepDaimyoCreateSamurai))
}
