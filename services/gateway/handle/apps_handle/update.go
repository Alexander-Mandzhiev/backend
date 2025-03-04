package apps_handle

import (
	"backend/pkg/server/respond"
	app "backend/protos/gen/go/apps"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *AppsService) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req app.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	if req.Id <= 0 {
		respond.RespondedError(w, r, http.StatusBadRequest, fmt.Errorf("ID must be provided in the request body"))
		return
	}

	resp, err := s.client.Update(context.Background(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
