package manager_account

import "go-hex-arch/internal/core/port"

type ManagerAccountHandler struct {
	svc port.ManagerAccountService
}

func New(svc port.ManagerAccountService) *ManagerAccountHandler {
	return &ManagerAccountHandler{
		svc,
	}
}
