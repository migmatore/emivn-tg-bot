package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerShogunHandler() {
	h.Message(h.ShogunHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepShogunMainMenuHandler))
	h.Message(h.ShogunHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepShogunHierarchyMenuHandler))
	h.Message(h.ShogunHandler.CreateEntityMenuHandler, h.isSessionStep(domain.SessionStepShogunCreateEntityMenuHandler))

	h.Message(h.ShogunHandler.EnterDaimyoNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateDaimyoNickname))
	h.Message(h.ShogunHandler.CreateDaimyoHandler, h.isSessionStep(domain.SessionStepShogunCreateDaimyo))

	h.Message(h.ShogunHandler.EnterSamuraiNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateSamuraiNickname))
	h.Message(h.ShogunHandler.ChooseSamuraiDaimyoHandler, h.isSessionStep(domain.SessionStepShogunChooseSamuraiDaimyo))
	h.Message(h.ShogunHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepShogunCreateSamurai))

	h.Message(h.ShogunHandler.EnterCashManagerNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateCashManagerNickname))
	h.Message(h.ShogunHandler.CreateCashManagerHandler, h.isSessionStep(domain.SessionStepShogunCreateCashManager))

}
