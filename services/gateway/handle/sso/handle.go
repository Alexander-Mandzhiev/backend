package sso_handle

import (
	"backend/protos/gen/go/sso"
	"context"
	"github.com/gin-gonic/gin"
)

type SSOService interface {
	SignIn(ctx context.Context, req *sso.SignInRequest) (string, error)
}

type Handler struct {
	service SSOService
}

func New(service SSOService) *Handler {
	return &Handler{service: service}
}

// initSSORoutes регистрирует маршруты для аутентификации
func (h *Handler) InitSSORoutes(api *gin.RouterGroup) {
	api.POST("/login", h.signIn)
}
