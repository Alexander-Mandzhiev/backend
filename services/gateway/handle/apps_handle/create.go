package apps_handle

import (
	"backend/pkg/server/respond"
	app "backend/protos/gen/go/apps"
	"context"
	"encoding/json"
	"net/http"
)

func (s *AppsService) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req app.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	resp, err := s.client.Create(context.Background(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusCreated, resp)
}
