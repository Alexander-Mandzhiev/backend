package sso_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sso"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (ss *SSOService) SignIn(c *gin.Context) {
	op := "sso.SignIn"
	var req sso.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sl.Log.Info("Handling sign-in request", slog.Int("password", int(req.GetPassword())), slog.String("op", op))

	resp, err := ss.client.SignIn(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error during sign-in", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	sl.Log.Info("Sign-in successful", slog.String("access_token", resp.GetAccessToken()), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
