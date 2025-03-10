package sso_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/sso"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	op := "sso.SignIn"
	var req sso.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	accessToken, err := h.service.SignIn(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Authentication failed", sl.Err(err), slog.Int("app_id", int(req.AppId)), slog.String("op", op))
		c.JSON(http.StatusUnauthorized, universalResponse.ErrorResponse("Invalid credentials"))
		return
	}

	sl.Log.Info("User signed in", slog.Int("app_id", int(req.AppId)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(accessToken))
}
