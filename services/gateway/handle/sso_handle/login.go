package sso_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sso"
	"context"
	"encoding/json"
	"net/http"
)

func (ss *SSOService) SignIn(w http.ResponseWriter, r *http.Request) {
	var req sso.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	resp, err := ss.client.SignIn(context.Background(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
