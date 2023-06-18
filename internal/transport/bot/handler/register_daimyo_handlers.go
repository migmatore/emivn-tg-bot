package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerDaimyoHandler() {
	h.Message(h.DaimyoHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepDaimyoMainMenuHandler)).
		Message(h.DaimyoHandler.EnterCardName, h.isSessionStep(domain.SessionStepDaimyoEnterReplenishmentRequestCardName)).
		Message(h.DaimyoHandler.MakeReplenishmentRequest, h.isSessionStep(domain.SessionStepDaimyoMakeReplenishmentRequest)).
		Message(h.DaimyoHandler.EnterReplenishmentRequestAmount, h.isSessionStep(domain.SessionStepDaimyoEnterReplenishmentRequestAmount)).
		Message(h.DaimyoHandler.MakeReplenishmentRequest, h.isSessionStep(domain.SessionStepDaimyoMakeReplenishmentRequest))

	h.Message(h.DaimyoHandler.ReportMenuHandler, h.isSessionStep(domain.SessionStepDaimyoReportMenuHandler))
	h.Message(h.DaimyoHandler.ReportPeriodMenuHandler, h.isSessionStep(domain.SessionStepDaimyoReportPeriodMenuHandler))

	h.Message(h.DaimyoHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepDaimyoHierarchyMenuHandler))
	h.Message(h.DaimyoHandler.EnterSamuraiUsernameHandler, h.isSessionStep(domain.SessionStepDaimyoCreateSamuraiUsername))
	h.Message(h.DaimyoHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepDaimyoCreateSamuraiNickname))
}
