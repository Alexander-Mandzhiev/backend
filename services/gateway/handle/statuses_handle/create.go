package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"encoding/json"
	"net/http"
)

func (s *StatusesService) Create(w http.ResponseWriter, r *http.Request) {
	var req statuses.CreateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := s.client.Create(r.Context(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusCreated, res)
}
